package osutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func Dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, Dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}

		// 隠しファイルはskip
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}
