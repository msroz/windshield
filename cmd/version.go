package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string = "0.0.1"

func init() {
	rootCmd.AddCommand(newVersionCmd())
}

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("windshield version: %s", Version)
		},
	}

	return cmd
}
