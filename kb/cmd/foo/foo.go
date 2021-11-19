package foo

import (
	"log"

	"kb/pkg/kb"

	"github.com/spf13/cobra"
)

var (
	cfg     Config
	Command = &cobra.Command{
		Use:   "foo [directory]",
		Short: "Test command",
		Args:  cobra.MinimumNArgs(1),
		Run:   executeFoo,
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

func executeFoo(cmd *cobra.Command, args []string) {
	log.Print("Executing command: foo")
	cfg.path = args[0]
	cfg.Log()

	kb.Bar(cfg.path, !cfg.noDryRun)
}
