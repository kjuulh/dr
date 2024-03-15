package pages

import (
	"log"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const PullRequestTablePage = "pull_request_page"

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)
)

type tableKeyMap struct {
	Begin key.Binding
	Help  key.Binding
	Quit  key.Binding
}

func newTableKeyMap() tableKeyMap {
	return tableKeyMap{
		Begin: key.NewBinding(
			key.WithKeys("b"),
			key.WithHelp("b", "begin reviewing pull requests"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "esc", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}
}

func (t tableKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		t.Begin, t.Help, t.Quit,
	}
}

func (t tableKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			t.Begin,
		},
		{
			t.Help, t.Quit,
		},
	}
}

type item struct {
	title string
}

func (i item) Title() string {
	return i.title
}

func (i item) Description() string {
	return i.title
}

func (i item) FilterValue() string {
	return i.title
}

var _ list.DefaultItem = &item{}

type PullRequestTable struct {
	list   list.Model
	keyMap tableKeyMap
	help   help.Model
}

func (p *PullRequestTable) Init() tea.Cmd {
	return nil
}

func (p *PullRequestTable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, p.keyMap.Begin):
			log.Println("should move to begin pr list")
			return p, nil
		case key.Matches(msg, p.keyMap.Help):
			p.help.ShowAll = !p.help.ShowAll
		case key.Matches(msg, p.keyMap.Quit):
			return p, tea.Quit

		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		p.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	p.list, cmd = p.list.Update(msg)
	return p, cmd
}

func (p *PullRequestTable) View() string {

	help := p.help.View(p.keyMap)
	helpHeight := lipgloss.Height(help)

	listHeight := p.list.Height()
	p.list.SetHeight(listHeight - helpHeight)
	list := p.list.View()
	p.list.SetHeight(listHeight)

	return docStyle.Render(
		lipgloss.JoinVertical(0, list, help),
	)
}

func NewPullRequestTable() *PullRequestTable {
	list := list.New([]list.Item{
		item{
			title: "something",
		},
		item{
			title: "something 123",
		},
		item{
			title: "something 456",
		},
	}, list.NewDefaultDelegate(), 0, 0)
	list.Title = "Pending Pull Requests"

	list.SetFilteringEnabled(false)
	list.SetShowStatusBar(false)
	list.SetShowHelp(false)

	return &PullRequestTable{
		list:   list,
		keyMap: newTableKeyMap(),
		help:   help.New(),
	}
}
