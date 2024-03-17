package pages

import (
	"bytes"
	"shuttle-extensions-template/internal/services"
	"shuttle-extensions-template/internal/utility"
	"strings"

	"github.com/alecthomas/chroma/quick"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const PullRequestReviewPage = "pull_request_review"

type reviewKeyMap struct {
	Skip    key.Binding
	TabNext key.Binding
	Help    key.Binding
}

func (r reviewKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			r.Skip,
			r.TabNext,
		},
		{
			r.Help,
		},
	}
}

func (r reviewKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		r.Skip,
		r.Help,
	}
}

func newReviewKeyMap() reviewKeyMap {
	return reviewKeyMap{
		Skip: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "skip the current pr"),
		),
		TabNext: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "switch to next interactive panel"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
	}
}

type PullRequestReview struct {
	keyMap      reviewKeyMap
	help        help.Model
	diff        viewport.Model
	description viewport.Model

	githubPrService *services.GitHubPullRequestService

	ready         bool
	width, height int
	currentPr     *services.GitHubPullRequest
	focus         int
}

func NewPullRequestReview() *PullRequestReview {
	return &PullRequestReview{
		keyMap: newReviewKeyMap(),
		help:   help.New(),

		githubPrService: services.NewGitHubPullRequestService(),

		currentPr: nil,
		focus:     0,
	}
}

func (p *PullRequestReview) Init() tea.Cmd {
	if p.currentPr == nil {
		pr, ok := p.githubPrService.GetNext()
		if ok {
			p.currentPr = pr
		}
	}

	return nil
}

func (p *PullRequestReview) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, p.keyMap.Skip):
			pr, ok := p.githubPrService.GetNext()
			if ok {
				p.currentPr = pr
			}

			return p, nil
		case key.Matches(msg, p.keyMap.TabNext):
			p.focus += 1
			p.focus %= 2
		case key.Matches(msg, p.keyMap.Help):
			p.help.ShowAll = !p.help.ShowAll

			// TODO: Don't refresh all the state only set height/width
			p.ready = false
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		p.SetSize(msg.Width-h, msg.Height-v)

		p.ready = false

	}

	if !p.ready {
		bytes := bytes.NewBufferString("")
		err := quick.Highlight(bytes, p.currentPr.Diff, "diff", "terminal16m", "dracula")
		if err != nil {
			panic(err)
		}
		height := p.getContentHeight()

		p.diff = p.createViewPort(bytes.String(), height/2)

		style := glamour.DefaultStyles["dracula"]
		style.Document.Margin = func() *uint {
			var zero uint = 0
			return &zero
		}()
		renderer, err := glamour.NewTermRenderer(glamour.WithStyles(*style), glamour.WithWordWrap(p.width/2-6))
		if err != nil {
			panic(err)
		}

		description, err := renderer.Render(p.currentPr.Description)
		if err != nil {
			panic(err)
		}
		p.description = p.createViewPort(description, height)

		p.ready = true
	}

	var (
		cmd  tea.Cmd
		cmds = make([]tea.Cmd, 0)
	)

	if p.focus == 0 {
		p.description, cmd = p.description.Update(msg)
		cmds = append(cmds, cmd)
	}
	if p.focus == 1 {
		p.diff, cmd = p.diff.Update(msg)
		cmds = append(cmds, cmd)
	}

	return p, tea.Batch(cmds...)
}

func (p *PullRequestReview) createViewPort(input string, height int) viewport.Model {
	diffStrings := strings.Split(input, "\n")
	renderedDiffStrings := make([]string, 0, len(diffStrings))
	for _, diffString := range diffStrings {
		// if len(diffString) > p.width/2-8 {
		// 	renderedDiffStrings = append(renderedDiffStrings, "\033[0m"+diffString[0:p.width/2-8]+"\033[0m")
		// } else {
		renderedDiffStrings = append(renderedDiffStrings, "\033[0m"+diffString+"\033[0m")
		// }

	}

	viewPort := viewport.New(p.width/2-6, height)
	viewPort.SetContent(strings.Join(renderedDiffStrings, "\n"))

	return viewPort
}

var (
	contentBox = lipgloss.NewStyle()

	titleBox = contentBox.Copy().Bold(true)
)

func borderBox(focus bool) lipgloss.Style {
	borderBox := lipgloss.NewStyle().Border(lipgloss.NormalBorder()).PaddingLeft(1)

	if focus {
		borderBox = borderBox.BorderForeground(lipgloss.Color("#FFFFFF"))
	} else {
		borderBox = borderBox.BorderForeground(lipgloss.Color("#AAAAAA"))

	}

	return borderBox
}

func (p *PullRequestReview) renderTitle() (string, int) {
	title := titleBox.Width(p.width-1).Render(p.currentPr.Title) + "\n"
	titleHeight := lipgloss.Height(title)

	return title, titleHeight
}

func (p *PullRequestReview) renderHelp() (string, int) {
	help := p.help.View(p.keyMap)
	helpHeight := lipgloss.Height(help)

	return help, helpHeight
}

func (p *PullRequestReview) getContentHeight() int {
	_, helpHeight := p.renderHelp()
	_, titleHeight := p.renderTitle()

	return p.height - (helpHeight + titleHeight) - 2
}

func (p *PullRequestReview) View() string {
	var body string
	help, _ := p.renderHelp()

	if p.currentPr != nil {
		pr := p.currentPr
		title, _ := p.renderTitle()

		comments := strings.Join(pr.Comments, "\n\n")
		statusChecks := strings.Join(pr.StatusChecks, "\n\n")
		diff := p.diff.View()

		remainingHeight := p.getContentHeight()

		left := lipgloss.PlaceHorizontal(
			p.width/2, lipgloss.Left,
			borderBox(p.focus == 0).
				Copy().
				Width(p.width/2).
				Height(remainingHeight-2).
				Render(p.description.View()),
		)
		rightTop := lipgloss.PlaceVertical(
			remainingHeight/2-1, lipgloss.Top,
			contentBox.Copy().Width(p.width/2-2).Render(
				lipgloss.JoinVertical(
					lipgloss.Top,
					borderBox(false).
						Copy().
						Width(p.width/2-4).
						Render(comments),
					borderBox(false).
						Copy().
						Width(p.width/2-4).
						Render(statusChecks)),
			),
		)

		rightBottom := lipgloss.PlaceVertical(
			remainingHeight-lipgloss.Height(rightTop), lipgloss.Top,
			borderBox(p.focus == 1).
				Copy().
				Width(p.width/2-4).
				Height(remainingHeight-lipgloss.Height(rightTop)).
				//MaxHeight(remainingHeight/2-1).
				Render(diff),
		)
		right := lipgloss.PlaceHorizontal(
			p.width/2, lipgloss.Left,
			lipgloss.JoinVertical(lipgloss.Top,
				rightTop, rightBottom),
		)

		content := lipgloss.JoinHorizontal(lipgloss.Left, left, right)
		//contentHeight := lipgloss.Height(content)

		body = lipgloss.JoinVertical(
			lipgloss.Top,
			title,
			contentBox.Copy().Height(remainingHeight).Render(content),
		)
	} else {
		body = "loading..."
	}

	notificationBox := lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Width(30).Render("something\nsomething\nsomething")

	content := docStyle.Render(
		lipgloss.JoinVertical(
			0,
			body,
			help,
		),
	)
	content = utility.PlaceOverlay(
		p.width-30, p.height,
		notificationBox, content,
		false,
	)

	return content
}

func (p *PullRequestReview) SetSize(width, height int) {
	p.width = width
	p.height = height
}

var _ tea.Model = &PullRequestReview{}
