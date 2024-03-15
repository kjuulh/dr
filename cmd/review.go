package cmd

import (
	"log"
	"shuttle-extensions-template/internal/ui"

	"github.com/spf13/cobra"
)

func ReviewCmd() *cobra.Command {
	var (
		squad string
	)

	cmd := &cobra.Command{
		Use: "review",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := ui.ReviewApp(cmd.Context()); err != nil {
				log.Fatal(err)
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&squad, "squad", "", "which squad to filter for, @lunarway/squad-aura")

	return cmd
}
