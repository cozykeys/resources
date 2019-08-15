namespace KbUtil.Lib.SvgGeneration.Internal
{
    using KbUtil.Lib.Models.Keyboard;
    using KbUtil.Lib.Extensions;
    using System.Collections.Generic;
    using System.Xml;

    internal class TextWriter : IElementWriter<Text>
    {
        private static readonly string DefaultTextAnchor = "middle";
        private static readonly string DefaultFont = "30px sans-serif";
        private static readonly string DefaultFill = "#000000";

        public SvgGenerationOptions GenerationOptions { get; set; }

        public void Write(XmlWriter writer, Text text)
        {
            writer.WriteStartElement("g");

            var elementWriter = new ElementWriter { GenerationOptions = GenerationOptions };

            elementWriter.WriteAttributes(writer, text);
            WriteAttributes(writer, text);

            elementWriter.WriteSubElements(writer, text);
            WriteSubElements(writer, text);

            writer.WriteEndElement();
        }

        public void WriteAttributes(XmlWriter writer, Text text)
        {
        }

        public void WriteSubElements(XmlWriter writer, Text text)
        {
            var styleDictionary = new Dictionary<string, string>();
            
            styleDictionary["font"] = !string.IsNullOrEmpty(text.Font)
                ? text.Font
                : DefaultFont;
            styleDictionary["fill"] = !string.IsNullOrEmpty(text.Fill)
                ? text.Fill
                : DefaultFill;

            string textAnchor = !string.IsNullOrEmpty(text.TextAnchor)
                ? text.TextAnchor
                : DefaultTextAnchor;
            
            if (GenerationOptions != null && GenerationOptions.EnableVisualSwitchCutouts == true)
            {
                writer.WriteStartElement("text");
                //writer.WriteAttributeString("id", "TODO");
                writer.WriteAttributeString("text-anchor", textAnchor);
                writer.WriteAttributeString("style", styleDictionary.ToCssStyleString());
                writer.WriteString(text.Content);
                writer.WriteEndElement();
            }
        }
    }
}
