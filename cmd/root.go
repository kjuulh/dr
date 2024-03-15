package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "shuttle-extensions-template",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("Hello, shuttle-extensions-template")

			return nil
		},
	}

	cmd.AddCommand(ReviewCmd())

	return cmd
}
