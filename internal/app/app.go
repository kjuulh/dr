package app

import (
	"shuttle-extensions-template/internal/pages"

	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	pages       map[string]Page
	currentPage string
}

func NewApp() *App {
	return &App{
		pages: map[string]Page{
			"pull_request_table": pages.NewPullRequestTable(),
		},
		currentPage: "pull_request_table",
	}
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
