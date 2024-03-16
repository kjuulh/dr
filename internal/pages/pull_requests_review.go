package pages

import (
	"shuttle-extensions-template/internal/services"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
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
	keyMap reviewKeyMap
	help   help.Model

	githubPrService *services.GitHubPullRequestService

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

	return p, nil
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
		description := pr.Description
		comments := strings.Join(pr.Comments, "\n\n")
		statusChecks := strings.Join(pr.StatusChecks, "\n\n")
		diff := pr.Diff

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
