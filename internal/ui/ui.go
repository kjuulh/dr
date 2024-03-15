package ui

import (
	"context"
	"shuttle-extensions-template/internal/app"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	prTitles []string = []string{
		"something",
		"something 123",
		"something 234",
	}
)

func ReviewApp(ctx context.Context) error {
	p := tea.NewProgram(app.NewApp(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}
