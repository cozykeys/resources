package svg

import (
	"errors"
	"fmt"
	"kb/pkg/models"
	"log"
	"os"
	"path"
	"strings"
)

const (
	DefaultSVGWidth  = 500
	DefaultSVGHeight = 500

	// TODO: Make this configurable in options? Or move to globals?
	DefaultSVGMarginHorizontal = 10
	DefaultSVGMarginVertical   = 10
)

type Options struct {
	EnableKeycapOverlays      bool
	EnableLegends             bool
	EnableVisualSwitchCutouts bool
	IndentString              string
	SquashLayers              bool
}

func temp() error {
	stack := &models.Stack{}
	width := stack.GetWidth()
	height := stack.GetWidth()
	log.Print(fmt.Sprintf("%f x %f", width, height))
	return nil
}

func Generate(kb *models.Keyboard, outDir string, opts *Options) error {
	xmlIndentString := "  "
	if opts.IndentString != "" {
		xmlIndentString = opts.IndentString
	}

	xmlSettings := &xmlWriterSettings{
		Indent:              true,
		IndentChars:         xmlIndentString,
		NewLineOnAttributes: true,
	}

	// Create output directory
	err := os.MkdirAll(outDir, 0755)
	if err != nil {
		return err
	}

	if opts.SquashLayers {
		return generateLayersSquashed(kb, outDir, opts, xmlSettings)
	} else {
		return generateLayers(kb, outDir, opts, xmlSettings)
	}
}

func generateLayersSquashed(kb *models.Keyboard, outDir string, opts *Options, xmlSettings *xmlWriterSettings) error {
	// TODO: Implement this
	return errors.New("not yet implemented: generateLayersSquashed")
}

func generateLayers(kb *models.Keyboard, outDir string, opts *Options, xmlSettings *xmlWriterSettings) error {
	for _, layer := range kb.Layers {
		outFile := path.Join(outDir,
			fmt.Sprintf("%s_%s.svg", strings.ToLower(kb.Name), strings.ToLower(layer.Name)))

		w := CreateXMLWriter(xmlSettings)

		writeSVGOpenTag(w, int(layer.GetWidth()), int(layer.Height))

		lw := &layerWriter{options: opts}
		lw.write(w, &layer)

		writeSVGCloseTag(w)

		w.writeToFile(outFile)
	}

	return nil
}

func writeSVGOpenTag(w *xmlWriter, width int, height int) {
	if width == 0 {
		width = DefaultSVGWidth
	}
	if height == 0 {
		height = DefaultSVGHeight
	}

	width += DefaultSVGMarginHorizontal
	height += DefaultSVGMarginVertical

	w.writeStartElement("svg")
	w.writeAttributeString("xmlns", "http://www.w3.org/2000/svg")
	w.writeAttributeString("width", fmt.Sprintf("%dmm", width))
	w.writeAttributeString("height", fmt.Sprintf("%dmm", height))
	w.writeAttributeString("viewBox", fmt.Sprintf("0 0 %d %d", width, height))
}

func writeSVGCloseTag(w *xmlWriter) {
	w.writeEndElement()
}

/*
func addChildComponent(parent *etree.Element, cmp models.Component) {
	switch cmp := cmp.(type) {
	case *models.Key:
		addChildKey(parent, cmp)
	case *models.Circle:
		addChildCircle(parent, cmp)
	default:
		panic(fmt.Sprintf("Unknown type %T", cmp))
	}
}

func addChildKey(parent *etree.Element, key *models.Key) {
	g := parent.CreateElement("g")
	g.CreateAttr("id", key.Name)
	g.CreateAttr("transform", fmt.Sprintf("translate(%.3f,%.3f)", key.XOffset, key.YOffset))

	e := g.CreateElement("path")
	e.CreateAttr("style", "fill:none;fill-opacity:1;stroke:#000000;stroke-width:0.5")
	e.CreateAttr("d", switchCutoutPathData)

	t1 := g.CreateElement("path")
	t1.CreateAttr("style", "fill:#5f5f5f;fill-opacity:1;stroke:#000000;stroke-width:0.5")
	t1.CreateAttr("d", keycapPathDataOuter)

	t2 := g.CreateElement("path")
	t2.CreateAttr("style", "fill:#7f7f7f;fill-opacity:1;stroke:#000000;stroke-width:0.5")
	t2.CreateAttr("d", keycapPathDataInner)
}

func addChildCircle(parent *etree.Element, circle *models.Circle) {
	g := parent.CreateElement("g")
	g.CreateAttr("id", circle.Name)
	g.CreateAttr("transform", fmt.Sprintf("translate(%.3f,%.3f)", circle.XOffset, circle.YOffset))

	e := g.CreateElement("circle")
	e.CreateAttr("style", "fill:none;stroke:#000000;stroke-width:0.5")
	e.CreateAttr("r", "1.0")
}
*/
