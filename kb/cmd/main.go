package main

import (
	"kb/cmd/svg"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "kb"}
	rootCmd.AddCommand(svg.Command)
	rootCmd.Execute()
}
