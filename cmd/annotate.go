package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/msroz/windshield/pkg/osutil"
)

var (
	annotation string
)

func init() {
	rootCmd.AddCommand(newAnnotateCmd())
}

func newAnnotateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "annotate",
		Short: "Annotate in files",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if Verbose {
				cmd.Printf("dir:%s\n", targetDir)
				cmd.Printf("annotation:%s\n", annotation)
			}
			list := osutil.Dirwalk(targetDir)
			for _, filePath := range list {
				if Verbose {
					cmd.Printf("filePath:%s\n", filePath)
				}
				file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
				if err != nil {
					cmd.Println(err)
					continue
				}
				defer file.Close()
				fmt.Fprintln(file, annotation)
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&targetDir, "dir", "d", "output directory (default $HOME/dir)")
	cmd.Flags().StringVar(&annotation, "annotation", "a", "annotation")
	cmd.MarkFlagRequired("dir")
	cmd.MarkFlagRequired("annotation")

	return cmd
}
