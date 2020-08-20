package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/msroz/windshield/pkg/osutil"
)

var (
	downloadListFilePath     string
	distDir                  string = "./dist"
	indexHTML                string = "index.html"
	userAgentPCChromeBrowser string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36"
)

func init() {
	rootCmd.AddCommand(newDownloadCmd())
}

func newDownloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download",
		Short: "Download and Save HTML Files in local",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if Verbose {
				cmd.Printf("downloadListFilePath: %s\n", downloadListFilePath)
			}
			if !osutil.Exists(downloadListFilePath) {
				return fmt.Errorf("File not found: %s", downloadListFilePath)
			}
			file, err := os.Open(downloadListFilePath)
			if err != nil {
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				u, err := url.Parse(line)
				if err != nil {
					cmd.Println(err)
					continue
				}

				if u.Scheme != "http" && u.Scheme != "https" {
					cmd.Printf("Each line text must be url. line: %s\n", line)
					continue
				}

				if Verbose {
					cmd.Printf("Downloading...%s\n", line)
				}
				byteArray, err := getRequest(line)
				if err != nil {
					cmd.Println(err)
					continue
				}

				saveDir := getSaveDir(u.Path)
				err = os.MkdirAll(saveDir, 0755)
				if err != nil {
					cmd.Println(err)
					continue
				}

				saveFilePath := getSaveFilePath(u.Path)
				err = ioutil.WriteFile(saveFilePath, byteArray, 0666)
				if err != nil {
					cmd.Println(err)
					continue
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&downloadListFilePath, "import", "", "import file path")
	cmd.Flags().StringVar(&distDir, "dist", "", "output directory (default $HOME/dist)")
	cmd.MarkFlagRequired("import")

	return cmd
}

func getRequest(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", userAgentPCChromeBrowser)
	client := new(http.Client)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code. expected: %d, actual: %d", resp.StatusCode, http.StatusOK)
	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return byteArray, nil
}

func getSaveDir(p string) string {
	var dir string
	// Pathが path/to/fileだと /path/to/file をdirとする
	if filepath.Ext(p) == "" {
		dir = distDir + p
	} else {
		// Pathが path/to/file.html だと /path/to をdirとする
		dir = distDir + filepath.Dir(p)
	}

	return dir
}

func getSaveFilePath(p string) string {
	var path string
	// Pathが path/to/fileだと /path/to/file/index.htmlとして保存する
	if filepath.Ext(p) == "" {
		path = distDir + p + "/" + indexHTML
	} else {
		// Pathが path/to/file.html だと /path/to/file.htmlとして保存する
		path = distDir + p
	}

	return path
}
