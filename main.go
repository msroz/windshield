package main

import (
	"fmt"
	"os"

	"github.com/msroz/windshield/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(-1)
	}
}
