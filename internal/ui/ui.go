package ui

import (
	"context"
	"shuttle-extensions-template/internal/app"
	"shuttle-extensions-template/internal/pages"

	tea "github.com/charmbracelet/bubbletea"
)

func ReviewApp(ctx context.Context) error {
	p := tea.NewProgram(app.NewApp(app.WithPage(pages.PullRequestTablePage)), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}
