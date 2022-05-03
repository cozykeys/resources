package curves

import (
	"log"

	"kb/pkg/kb"

	"github.com/spf13/cobra"
)

var (
	flags     = new(curvesFlags)
	arguments = new(curvesArgs)
)

type curvesFlags struct {
	svg string
}

type curvesArgs struct {
	inputFile  string
	outputFile string
	distance   float64
}

func Register(parent *cobra.Command) {
	command := &cobra.Command{
		Use:   "curves <input-path> <output-path> <distance> [OPTIONS]",
		Short: "Generate curves from a set of vertices",
		Args:  cobra.ExactArgs(3),
		Run:   curvesFunc,
	}

	command.PersistentFlags().StringVarP(&flags.svg,
		"svg", "s", "", "render an SVG with the results")

	parent.AddCommand(command)
}

func logInput(args *curvesArgs, flags *curvesFlags) {
	log.Printf("Args:")
	log.Printf("- inputFile: %q", args.inputFile)
	log.Printf("- outputFile: %q", args.outputFile)
	log.Printf("- distance: %f", args.distance)
	log.Printf("Flags:")
	log.Printf("- svg: %q", flags.svg)
}

func curvesFunc(cmd *cobra.Command, args []string) {
	log.Print("Executing command: curves")
	arguments.inputFile = args[0]
	arguments.outputFile = args[1]
	arguments.distance = kb.MustParseFloat64(args[2])
	logInput(arguments, flags)

	// TODO
	err := func() error {
		inputVertices, err := kb.GetInputVertices(arguments.inputFile)
		if err != nil {
			return err
		}

		log.Print("Successfully parsed input vertices:")
		for _, p := range inputVertices {
			log.Printf("- %s", string(kb.MustMarshalJSON(p, false)))
		}

		curves, err := kb.GenerateCurves(inputVertices, arguments.distance)
		if err != nil {
			return err
		}

		log.Print("Successfully generated curves:")
		for _, c := range curves {
			log.Printf("- %s", string(kb.MustMarshalJSON(c, false)))
		}

		if flags.svg != "" {
			err := kb.WriteCurvesSVG(flags.svg, inputVertices, curves)
			if err != nil {
				return err
			}
		}

		return nil
	}()

	if err != nil {
		log.Fatalf("fatal error: %v", err)
	}
}
