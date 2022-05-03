package expand

import (
	"log"
	"os"

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
		Use:   "expand <input-path> <output-path> <distance>",
		Short: "Expand a set of vertices",
		Args:  cobra.ExactArgs(3),
		Run:   expandFunc,
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

func expandFunc(cmd *cobra.Command, args []string) {
	log.Print("Executing command: expand")
	arguments.inputFile = args[0]
	arguments.outputFile = args[1]
	arguments.distance = kb.MustParseFloat64(args[2])
	logInput(arguments, flags)

	err := func() error {
		inputVertices, err := kb.GetInputVertices(arguments.inputFile)
		if err != nil {
			return err
		}

		log.Print("Successfully parsed input vertices:")
		for _, p := range inputVertices {
			log.Printf("- %s", string(kb.MustMarshalJSON(p, false)))
		}

		inputSegments, err := kb.ConvertPointsToSegments(inputVertices)
		if err != nil {
			return err
		}

		log.Print("Successfully convert input vertices to segments:")
		for _, s := range inputSegments {
			log.Printf("- %s", string(kb.MustMarshalJSON(s, false)))
		}

		expandedLines, err := kb.ExpandSegments(inputSegments, arguments.distance)
		if err != nil {
			return err
		}

		log.Print("Successfully expanded segments:")
		for _, l := range expandedLines {
			log.Printf("- %s", string(kb.MustMarshalJSON(l, false)))
		}

		expandedVertices, err := kb.GetIntersectionPoints(expandedLines)
		if err != nil {
			return err
		}

		log.Print("Successfully expanded vertices:")
		for _, p := range expandedVertices {
			log.Printf("- %s", string(kb.MustMarshalJSON(p, false)))
		}

		if flags.svg != "" {
			err := kb.WriteExpandSVG(flags.svg, inputVertices, expandedVertices)
			if err != nil {
				return err
			}
		}

		err = os.WriteFile(arguments.outputFile,
			kb.MustMarshalJSON(expandedVertices, false), 0644)
		if err != nil {
			return err
		}

		return nil
	}()

	if err != nil {
		log.Fatalf("fatal error: %v", err)
	}
}
