package svg

import (
	"bytes"
	"fmt"
	"kb/pkg/models"
	"os"
	"path"
	"strings"

	"github.com/beevik/etree"
)

type svgOptions struct {
	KeycapOverlaysEnabled bool
}

// TODO: This shouldn't be global
var options *svgOptions = &svgOptions{
	KeycapOverlaysEnabled: true,
}

func GenerateSVG(kb *models.Keyboard, outputDirectory string) error {
	generator := &generator{
		keyboard:        kb,
		outputDirectory: outputDirectory,
	}
	return generator.generate()
}

type generator struct {
	keyboard        *models.Keyboard
	outputDirectory string
}

/*
   public static void GenerateSvg(Keyboard keyboard, string outputDirectory, SvgGenerationOptions options = null)
   {
       var settings = new XmlWriterSettings
       {
           Indent = true,
           IndentChars = options?.IndentString ?? "  ",
           NewLineOnAttributes = true
       };

       Directory.CreateDirectory(outputDirectory);

       if (options != null && options.SquashLayers)
       {
           GenerateLayersSquashed(keyboard, outputDirectory, options, settings);
       }
       else
       {
           GenerateLayers(keyboard, outputDirectory, options, settings);
       }
   }
*/
func (g *generator) generate() error {
	// Create output directory
	err := os.MkdirAll(g.outputDirectory, 0755)
	if err != nil {
		return err
	}

	return g.generateLayers()
}

/*
   private static void GenerateLayers(Keyboard keyboard, string outputDirectory, SvgGenerationOptions options, XmlWriterSettings settings)
   {
       foreach (var layer in keyboard.Layers)
       {
           string path = System.IO.Path.Combine(outputDirectory, $"{keyboard.Name}_{layer.Name}.svg");

           using (FileStream stream = File.Open(path, FileMode.Create))
           using (XmlWriter writer = XmlWriter.Create(stream, settings))
           {
               WriteSvgOpenTag(writer, (int)layer.Width, (int)layer.Height);

               var layerWriter = new LayerWriter { GenerationOptions = options };
               layerWriter.Write(writer, layer);

               WriteSvgCloseTag(writer);
           }
       }
   }
*/
func (g *generator) generateLayers() error {
	for _, layer := range g.keyboard.Layers {
		filename := fmt.Sprintf("%s_%s.svg",
			strings.ToLower(g.keyboard.Name), strings.ToLower(layer.Name))
		outputFile := path.Join(g.outputDirectory, filename)

		// Create XML Document
		doc := etree.NewDocument()
		doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

		// Set up the root SVG element
		w := int32(g.keyboard.Width + 10)  // TODO
		h := int32(g.keyboard.Height + 10) // TODO
		root := doc.CreateElement("svg")
		root.CreateAttr("width", fmt.Sprintf("%dmm", w))
		root.CreateAttr("height", fmt.Sprintf("%dmm", h))
		root.CreateAttr("viewBox", fmt.Sprintf("0 0 %d %d", w, h))
		root.CreateAttr("xmlns", "http://www.w3.org/2000/svg")

		err := writeLayer(root, &layer)
		if err != nil {
			//return err
		}

		/*
			for _, cmp := range kb.Components {
				addChildComponent(root, cmp)
			}
		*/

		doc.Indent(4)
		b := &bytes.Buffer{}
		doc.WriteTo(b)

		os.WriteFile(outputFile, b.Bytes(), 0644)
	}

	return nil
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
*/

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
