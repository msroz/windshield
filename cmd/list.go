package cmd

import (
	"github.com/spf13/cobra"

	"github.com/msroz/windshield/pkg/osutil"
)

func init() {
	rootCmd.AddCommand(newListCmd())
}

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Print files",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if Verbose {
				cmd.Printf("dir:%s\n", targetDir)
			}
			list := osutil.Dirwalk(targetDir)
			for _, file := range list {
				cmd.Println(file)
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&targetDir, "dir", "", "output directory (default $HOME/dir)")
	cmd.MarkFlagRequired("dir")

	return cmd
}
