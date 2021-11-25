package svg

import (
	"encoding/json"
	"log"
	"os"

	"kb/pkg/kb"

	"github.com/spf13/cobra"
)

var (
	cfg     Config
	Command = &cobra.Command{
		Use:   "svg <input-file.json> <output-file.svg>",
		Short: "Generate SVG file from keyboard data",
		Args:  cobra.ExactArgs(2),
		Run:   svg,
	}
)

type Config struct {
	path     string
	noDryRun bool
}

func (cfg Config) Log() {
	log.Printf("Config:")
	log.Printf("- path: %q", cfg.path)
	log.Printf("- dryRun: %t", !cfg.noDryRun)
}

func init() {
	Command.SetArgs([]string{"directory"})
	Command.Flags().BoolVar(&cfg.noDryRun, "no-dry-run", false,
		"Disables the default \"dry-run\" behavior")
}

func svg(cmd *cobra.Command, args []string) {
	log.Print("Executing command: svg")
	cfg.path = args[0]
	cfg.Log()

	inputFilePath := args[0]
	outputFilePath := args[1]

	bytes, err := os.ReadFile(inputFilePath)
	if err != nil {
		log.Fatalf("failed to read input file: %v\n", err)
	}

	keyboard := &kb.Keyboard{}
	if err := json.Unmarshal(bytes, keyboard); err != nil {
		log.Fatalf("failed to unmarshal input file: %v\n", err)
	}

	svg, err := keyboard.ToSvg([]string{})
	if err != nil {
		log.Fatalf("failed to generate svg: %v\n", err)
	}

	err = os.WriteFile(outputFilePath, []byte(svg), os.FileMode(int(0664)))
	if err != nil {
		log.Fatalf("failed to write to output file: %v\n", err)
	}
}
