package cmd

import (
	"github.com/spf13/cobra"
)

var (
	Verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "windshield",
	Short: "CLI tools",
	Long:  "CLI tools",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()

	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}
