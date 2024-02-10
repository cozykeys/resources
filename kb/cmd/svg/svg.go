package svg

import (
	"kb/pkg/svg"
	"kb/pkg/unmarshal"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfg     Config
	Command = &cobra.Command{
		Use:   "svg <input-file.xml> <output-directory>",
		Short: "Generate SVG file from keyboard data",
		Args:  cobra.ExactArgs(2),
		Run:   svgFunc,
	}
)

type Config struct {
	inputFile string
	outputDir string
}

func (cfg Config) Log() {
	log.Printf("Config:")
	log.Printf("- inputFile: %q", cfg.inputFile)
	log.Printf("- outputDir: %q", cfg.outputDir)
}

func init() {
	Command.SetArgs([]string{"directory"})
}

func svgFunc(cmd *cobra.Command, args []string) {
	log.Print("Executing command: svg")
	cfg.inputFile = args[0]
	cfg.outputDir = args[1]
	cfg.Log()

	bytes, err := os.ReadFile(cfg.inputFile)
	if err != nil {
		log.Fatalf("failed to read input file: %v\n", err)
	}

	keyboard, err := unmarshal.Unmarshal(bytes)
	if err != nil {
		log.Fatalf("failed to unmarshal input file: %v\n", err)
	}

	options := &svg.Options{
		EnableKeycapOverlays: true,
	}

	err = svg.Generate(keyboard, cfg.outputDir, options)
	if err != nil {
		log.Fatalf("failed to generate svg: %v\n", err)
	}
}
