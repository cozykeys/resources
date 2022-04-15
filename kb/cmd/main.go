package main

import (
	"kb/cmd/svg"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "kb"}
	rootCmd.AddCommand(svg.Command)

	// TODO: The following commands exist in the C# kbutil CLI and
	// still need to be ported over.
	/*
		draw-svg-holes
		draw-svg-path
		draw-switches
		expand-vertices
		expand-vertices2
		generate-curves
		gen-key-bearings
		gen-svg
		scratch
	*/

	rootCmd.Execute()
}
