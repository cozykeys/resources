namespace KbUtil.Lib.SvgGeneration.Internal
{
    using System.Collections.Generic;
    using System.Xml;
    using KbUtil.Lib.Extensions;
    using KbUtil.Lib.Models.Keyboard;

    internal class LegendWriter : IElementWriter<Legend>
    {
        public SvgGenerationOptions GenerationOptions { get; set; }

        public void Write(XmlWriter writer, Legend legend)
        {
            writer.WriteStartElement("text");

            var elementWriter = new ElementWriter { GenerationOptions = GenerationOptions };

            elementWriter.WriteAttributes(writer, legend);
            WriteAttributes(writer, legend);

            elementWriter.WriteSubElements(writer, legend);
            WriteSubElements(writer, legend);

            writer.WriteEndElement();
        }

        public void WriteAttributes(XmlWriter writer, Legend legend)
        {
            writer.WriteAttributeString("text-anchor", "middle");

            float fontSize = legend.FontSize is default(float) ? 4 : legend.FontSize;

            var styleDictionary = new Dictionary<string, string>
            {
                { "fill", !string.IsNullOrWhiteSpace(legend.Color) ? legend.Color : "#000000" },
                { "dominant-baseline", "central" },
                { "text-anchor", "middle" },
                { "font-size", $"{fontSize}px" },
                { "font-family", "sans-serif" },
                { "font-weight", "normal" },
                { "font-style", "normal" },
            };

            writer.WriteAttributeString("style", styleDictionary.ToCssStyleString());
        }

        public void WriteSubElements(XmlWriter writer, Legend legend)
        {
            writer.WriteString(legend.Text);
        }
    }
}
