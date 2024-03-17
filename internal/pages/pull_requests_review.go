package pages

import (
	"bytes"
	"shuttle-extensions-template/internal/services"
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
	Skip key.Binding
}

func (r reviewKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			r.Skip,
		},
	}
}

func (r reviewKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		r.Skip,
	}
}

func newReviewKeyMap() reviewKeyMap {
	return reviewKeyMap{
		Skip: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "skip the current pr"),
		),
	}
}

type PullRequestReview struct {
	keyMap   reviewKeyMap
	help     help.Model
	viewPort viewport.Model

	githubPrService *services.GitHubPullRequestService

	ready         bool
	width, height int
	currentPr     *services.GitHubPullRequest
}

func NewPullRequestReview() *PullRequestReview {
	return &PullRequestReview{
		keyMap: newReviewKeyMap(),
		help:   help.New(),

		githubPrService: services.NewGitHubPullRequestService(),

		currentPr: nil,
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
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		p.SetSize(msg.Width-h, msg.Height-v)

	}

	if !p.ready {
		bytes := bytes.NewBufferString("")
		err := quick.Highlight(bytes, p.currentPr.Diff, "diff", "terminal16m", "dracula")
		if err != nil {
			panic(err)
		}

		diffStrings := strings.Split(bytes.String(), "\n")
		renderedDiffStrings := make([]string, 0, len(diffStrings))
		for _, diffString := range diffStrings {
			renderedDiffStrings = append(renderedDiffStrings, "\033[0m"+diffString+"\033[0m")
		}

		p.viewPort = viewport.New(p.width, p.height)
		p.viewPort.SetContent(strings.Join(renderedDiffStrings, "\n"))
		p.ready = true
	}

	var (
		cmd  tea.Cmd
		cmds = make([]tea.Cmd, 0)
	)

	p.viewPort, cmd = p.viewPort.Update(msg)
	cmds = append(cmds, cmd)

	return p, tea.Batch(cmds...)
}

var (
	contentBox = lipgloss.NewStyle()

	borderBox = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).PaddingLeft(1)

	titleBox = contentBox.Copy().Bold(true)
)

func (p *PullRequestReview) View() string {
	var body string
	help := p.help.View(p.keyMap)
	helpHeight := lipgloss.Height(help)

	if p.currentPr != nil {
		pr := p.currentPr
		title := pr.Title

		style := glamour.DefaultStyles["dracula"]
		style.Document.Margin = func() *uint {
			var zero uint = 0
			return &zero
		}()
		renderer, err := glamour.NewTermRenderer(glamour.WithStyles(*style), glamour.WithWordWrap(p.width/2-2))
		if err != nil {
			panic(err)
		}

		description, err := renderer.Render(pr.Description)
		if err != nil {
			panic(err)
		}

		comments := strings.Join(pr.Comments, "\n\n")
		statusChecks := strings.Join(pr.StatusChecks, "\n\n")
		diff := p.viewPort.View()

		title = titleBox.Width(p.width-1).Render(title) + "\n"
		titleHeight := lipgloss.Height(title)

		remainingHeight := p.height - (titleHeight + helpHeight)

		left := lipgloss.PlaceHorizontal(
			p.width/2, lipgloss.Left,
			borderBox.
				Copy().
				Width(p.width/2).
				Height(remainingHeight-2).
				Render(description),
		)
		rightTop := lipgloss.PlaceVertical(
			remainingHeight/2, lipgloss.Top,
			contentBox.Copy().Width(p.width/2-2).Render(
				lipgloss.JoinVertical(
					lipgloss.Top,
					borderBox.
						Copy().
						Width(p.width/2-4).
						Render(comments),
					borderBox.
						Copy().
						Width(p.width/2-4).
						Render(statusChecks)),
			),
		)
		rightBottom := lipgloss.PlaceVertical(
			remainingHeight/2, lipgloss.Top,
			borderBox.
				Copy().
				Width(p.width/2-4).
				Height(remainingHeight/2-1).
				MaxHeight(remainingHeight/2-1).
				Render(diff),
		)
		right := lipgloss.PlaceHorizontal(
			p.width/2-2, lipgloss.Left,
			lipgloss.JoinVertical(lipgloss.Top,
				rightTop, rightBottom),
		)

		content := lipgloss.JoinHorizontal(lipgloss.Left, left, right)
		//contentHeight := lipgloss.Height(content)

		body = lipgloss.JoinVertical(lipgloss.Top, title, contentBox.Copy().Height(remainingHeight).Render(content))
	} else {
		body = "loading..."
	}

	return docStyle.Render(
		lipgloss.JoinVertical(
			0,
			body,
			help,
		),
	)
}

func (p *PullRequestReview) SetSize(width, height int) {
	p.width = width
	p.height = height
}

var _ tea.Model = &PullRequestReview{}
