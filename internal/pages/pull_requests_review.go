package pages

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
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
}

func NewPullRequestReview() *PullRequestReview {
	return &PullRequestReview{
		keyMap: newReviewKeyMap(),
		help:   help.New(),
	}
}

func (p *PullRequestReview) Init() tea.Cmd {
	return nil
}

func (p *PullRequestReview) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return p, nil
}

func (p *PullRequestReview) View() string {
	return docStyle.Render(
		p.help.View(p.keyMap),
	)
}

var _ tea.Model = &PullRequestReview{}
