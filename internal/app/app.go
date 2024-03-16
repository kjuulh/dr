package app

import (
	"fmt"
	"shuttle-extensions-template/internal/pages"

	tea "github.com/charmbracelet/bubbletea"
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
	}

	cmds := make([]tea.Cmd, 0)

	if a.pages[a.currentPage] != nil {
		newPage, newCmd := a.pages[a.currentPage].Update(msg)
		a.pages[a.currentPage] = newPage
		cmds = append(cmds, newCmd)
	}

	return a, tea.Batch(cmds...)
}

func (a *App) View() string {
	return a.pages[a.currentPage].View()
}

var _ tea.Model = &App{}

type Page interface {
	tea.Model
}
