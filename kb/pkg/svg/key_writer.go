package svg

import (
	"kb/pkg/models"
	"log"
	"strings"

	"github.com/beevik/etree"
)

func writeKey(parent *etree.Element, key *models.Key) error {
	writer := &keyWriter{
		parent: parent,
		key:    key,
	}

	return writer.write()
}

type keyWriter struct {
	parent *etree.Element
	key    *models.Key
}

func (w *keyWriter) write() error {
	/*
	   elementWriter.WriteAttributes(writer, key);
	   WriteAttributes(writer, key);

	   elementWriter.WriteSubElements(writer, key);
	   WriteSubElements(writer, key);

	   writer.WriteEndElement();
	*/
	/*
		TODO
		if !w.key.Visible {
			return nil
		}
	*/

	g := w.parent.CreateElement("g")

	// Attributes
	g.CreateAttr("id", w.key.Name)
	writeTransform(g, w.key)

	// Elements
	if w.key.Debug {
		log.Printf("Writing debug overlay...")
		writeDebugOverlay(g, w.key)
	}

	w.writeSwitchCutoutPath(g)
	/*
	   WriteKeycapOverlay(writer, key);
	   WriteKeyLegends(writer, key);
	*/

	return nil
}

/*
   private void WriteSwitchCutoutPath(XmlWriter writer, Key key)
   {
       var pathWriter = new PathWriter { GenerationOptions = GenerationOptions };
       pathWriter.Write(writer, _switchPath);
   }
*/
func (w *keyWriter) writeSwitchCutoutPath(parent *etree.Element) {
	e := parent.CreateElement("path")
	e.CreateAttr("style", "fill:none;fill-opacity:1;stroke:#000000;stroke-width:0.5")
	e.CreateAttr("d", switchCutoutPathData)
}

/*
   private void WriteKeycapOverlay(XmlWriter writer, Key key)
   {
       if (GenerationOptions == null || GenerationOptions.EnableKeycapOverlays == false)
       {
           return;
       }

       // Give these short names so the resulting path data is readable
       float w = key.Width;
       float h = key.Height;

       // Next we write it with a style that is more visually pleasing
       writer.WriteStartElement("path");
       writer.WriteAttributeString("id", $"{key.Name}KeycapOverlay");
       writer.WriteAttributeString("d", $"M -{w / 2},-{h / 2} h {w} v {h} h -{w} v -{h} h {w}");

       var styleDictionary = new Dictionary<string, string>
       {
           { "fill", !string.IsNullOrWhiteSpace(key.Fill) ? key.Fill : "#ffffff" },
           { "stroke", !string.IsNullOrWhiteSpace(key.Stroke) ? key.Stroke : "#000000" },
           { "stroke-width", "0.5" },
       };

       writer.WriteAttributeString("style", styleDictionary.ToCssStyleString());
       writer.WriteEndElement();
   }
*/
/*
   private void WriteKeyLegends(XmlWriter writer, Key key)
   {
       if (GenerationOptions == null || GenerationOptions.EnableKeycapOverlays == false)
       {
           return;
       }

       if (key.Legends == null || !key.Legends.Any())
       {
           return;
       }

       int legendIndex = 0;
       var legendWriter = new LegendWriter { GenerationOptions = GenerationOptions };
       foreach (Legend legend in key.Legends)
       {
           legend.Name = $"{key.Name}Legend{legendIndex}";
           legendWriter.Write(writer, legend);
           legendIndex++;
       }
   }
*/

var switchCutoutPathData string = strings.Join([]string{
	"M    0 -7",
	"L    7 -7",
	"L    7 -6",
	"L  7.8 -6",
	"L  7.8  6",
	"L    7  6",
	"L    7  7",
	"L   -7  7",
	"L   -7  6",
	"L -7.8  6",
	"L -7.8 -6",
	"L   -7 -6",
	"L   -7 -7",
	"L    0 -7",
}, " ")

var keycapPathDataOuter string = strings.Join([]string{
	"M     0 -9.05",
	"L  8.05 -9.05",
	"Q  9.05 -9.05",
	"   9.05 -8.05",
	"L  9.05  8.05",
	"Q  9.05  9.05",
	"   8.05  9.05",
	"L -8.05  9.05",
	"Q -9.05  9.05",
	"  -9.05  8.05",
	"L -9.05 -8.05",
	"Q -9.05 -9.05",
	"  -8.05 -9.05",
	"L     0 -9.05",
}, " ")

var keycapPathDataInner string = strings.Join([]string{
	"M     0 -8.05",
	"L  4.05 -8.05",
	"Q  6.05 -8.05",
	"   6.05 -6.05",
	"L  6.05  4.55",
	"Q  6.05  6.55",
	"   4.05  6.55",
	"L -4.05  6.55",
	"Q -6.05  6.55",
	"  -6.05  4.55",
	"L -6.05 -6.05",
	"Q -6.05 -8.05",
	"  -4.05 -8.05",
	"L     0 -8.05",
}, " ")
