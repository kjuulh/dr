package app

import (
	"fmt"
	"shuttle-extensions-template/internal/pages"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)
)

type AppOptions func(*App)

func WithPage(page string) AppOptions {
	return func(a *App) {
		a.currentPage = page
	}
}

type App struct {
	pages       map[string]Page
	currentPage string

	width, height int
}

func NewApp(opts ...AppOptions) *App {
	app := &App{
		pages: map[string]Page{
			pages.PullRequestTablePage:  pages.NewPullRequestTable(),
			pages.PullRequestReviewPage: pages.NewPullRequestReview(),
		},
		currentPage: pages.PullRequestTablePage,
	}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

func (a *App) Init() tea.Cmd {
	return nil
}

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			return a, tea.Quit
		}
	case pages.ChangePage:
		page := msg.Page()
		_, ok := a.pages[page]
		if !ok {
			panic(fmt.Errorf("error: page was not found: %s", page))
		}
		a.currentPage = page

		cmds = append(cmds, a.pages[page].Init())
		a.pages[page].SetSize(a.width, a.height)
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		a.SetSize(msg.Width-h, msg.Height-v)
	}

	if a.pages[a.currentPage] != nil {
		newPage, newCmd := a.pages[a.currentPage].Update(msg)
		a.pages[a.currentPage] = newPage.(Page)
		cmds = append(cmds, newCmd)
	}

	return a, tea.Batch(cmds...)
}

func (a *App) View() string {
	return a.pages[a.currentPage].View()
}

func (a *App) SetSize(width, height int) {
	a.width = width
	a.height = height
}

var _ tea.Model = &App{}

type Page interface {
	tea.Model
	SetSize(width, height int)
}
