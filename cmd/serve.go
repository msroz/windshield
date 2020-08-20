package cmd

import (
	"net/http"

	"github.com/spf13/cobra"
)

var (
	port      string
	targetDir string = "./dist"
)

func init() {
	rootCmd.AddCommand(newServeCmd())
}

func newServeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve static files in local",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if Verbose {
				cmd.Printf("dir:%s port: %s\n", targetDir, port)
			}

			http.Handle("/", http.FileServer(http.Dir(targetDir)))
			cmd.Printf("Serving %s on HTTP port: %s\n", targetDir, port)
			err := http.ListenAndServe(":"+port, nil)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&targetDir, "dir", "", "output directory (default $HOME/dist)")
	cmd.Flags().StringVarP(&port, "port", "p", "8001", "port (default 8001)")
	cmd.MarkFlagRequired("dir")

	return cmd
}
