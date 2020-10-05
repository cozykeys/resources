namespace KbUtil.Lib.Deserialization.Internal
{
    using KbUtil.Lib.Deserialization.Extensions;
    using KbUtil.Lib.Models.Keyboard;
    using System.Xml.Linq;

    internal class TextDeserializer : IDeserializer<Text>
    {
        public static TextDeserializer Default { get; set; } = new TextDeserializer();

        public void Deserialize(XElement textElement, Text text)
        {
            ElementDeserializer.Default.Deserialize(textElement, text);

            DeserializeContent(textElement, text);
            DeserializeTextAnchor(textElement, text);
            DeserializeFont(textElement, text);
            DeserializeFill(textElement, text);
        }

        private void DeserializeContent(XElement textElement, Text text)
        {
            if (XmlUtilities.TryGetAttribute(textElement, "Content", out XAttribute contentAttribute))
            {
                text.Content = contentAttribute.ValueAsString(text);
            }
        }

        private void DeserializeTextAnchor(XElement textElement, Text text)
        {
            if (XmlUtilities.TryGetAttribute(textElement, "TextAnchor", out XAttribute textAnchorAttribute))
            {
                text.TextAnchor = textAnchorAttribute.ValueAsString(text);
            }
        }

        private void DeserializeFont(XElement textElement, Text text)
        {
            if (XmlUtilities.TryGetAttribute(textElement, "Font", out XAttribute fontAttribute))
            {
                text.Font = fontAttribute.ValueAsString(text);
            }
        }

        private void DeserializeFill(XElement textElement, Text text)
        {
            if (XmlUtilities.TryGetAttribute(textElement, "Fill", out XAttribute fillAttribute))
            {
                text.Fill = fillAttribute.ValueAsString(text);
            }
        }
    }
}
