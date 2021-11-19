package main

import (
	"kb/cmd/foo"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "kb"}
	rootCmd.AddCommand(foo.Command)
	rootCmd.Execute()
}
