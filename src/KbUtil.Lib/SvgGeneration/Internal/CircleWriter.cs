namespace KbUtil.Lib.SvgGeneration.Internal
{
    using KbUtil.Lib.Models.Keyboard;
    using KbUtil.Lib.Extensions;
    using System.Collections.Generic;
    using System.Xml;

    internal class CircleWriter : IElementWriter<Circle>
    {
        private static readonly string DefaultFill = "none";
        private static readonly string DefaultStroke = "#0000ff";
        private static readonly string DefaultStrokeWidth = "0.01";

        public SvgGenerationOptions GenerationOptions { get; set; }

        public void Write(XmlWriter writer, Circle circle)
        {
            if (!circle.Visible)
                return;

            writer.WriteStartElement("g");

            var elementWriter = new ElementWriter { GenerationOptions = GenerationOptions };

            elementWriter.WriteAttributes(writer, circle);
            WriteAttributes(writer, circle);

            elementWriter.WriteSubElements(writer, circle);
            WriteSubElements(writer, circle);

            writer.WriteEndElement();
        }

        public void WriteAttributes(XmlWriter writer, Circle circle)
        {
        }

        public void WriteSubElements(XmlWriter writer, Circle circle)
        {
            var styleDictionary = new Dictionary<string, string>();
            
            styleDictionary["fill"] = !string.IsNullOrEmpty(circle.Fill)
                ? circle.Fill
                : DefaultFill;
            styleDictionary["stroke"] = !string.IsNullOrEmpty(circle.Stroke)
                ? circle.Stroke
                : DefaultStroke;
            styleDictionary["stroke-width"] = !string.IsNullOrEmpty(circle.StrokeWidth)
                ? circle.StrokeWidth
                : DefaultStrokeWidth;

            // First we write it with the style that Ponoko expects
            writer.WriteStartElement("circle");
            //writer.WriteAttributeString("id", "TODO");
            writer.WriteAttributeString("r", $"{circle.Size/2}");
            writer.WriteAttributeString("style", "fill:none;stroke:#0000ff;stroke-width:0.01");
            writer.WriteEndElement();

            // Next we write it with a style that is more visually pleasing
            if (GenerationOptions != null && GenerationOptions.EnableVisualSwitchCutouts == true)
            {
                writer.WriteStartElement("circle");
                //writer.WriteAttributeString("id", "TODO");
                writer.WriteAttributeString("r", $"{circle.Size/2}");
                writer.WriteAttributeString("style", styleDictionary.ToCssStyleString());
                writer.WriteEndElement();
            }
        }
    }
}
