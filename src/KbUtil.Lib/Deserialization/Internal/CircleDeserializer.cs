namespace KbUtil.Lib.Deserialization.Internal
{
    using KbUtil.Lib.Deserialization.Extensions;
    using KbUtil.Lib.Models.Keyboard;
    using System.Xml.Linq;

    internal class CircleDeserializer : IDeserializer<Circle>
    {
        public static CircleDeserializer Default { get; set; } = new CircleDeserializer();

        public void Deserialize(XElement holeElement, Circle circle)
        {
            ElementDeserializer.Default.Deserialize(holeElement, circle);

            DeserializeSize(holeElement, circle);
        }

        private void DeserializeSize(XElement holeElement, Circle circle)
        {
            if (XmlUtilities.TryGetAttribute(holeElement, "Size", out XAttribute sizeAttribute))
            {
                circle.Size = sizeAttribute.ValueAsFloat(circle);
            }
        }

        private void DeserializeFill(XElement holeElement, Circle circle)
        {
            if (XmlUtilities.TryGetAttribute(holeElement, "Fill", out XAttribute sizeAttribute))
            {
                circle.Fill = sizeAttribute.ValueAsString(circle);
            }
        }

        private void DeserializeStroke(XElement holeElement, Circle circle)
        {
            if (XmlUtilities.TryGetAttribute(holeElement, "Stroke", out XAttribute sizeAttribute))
            {
                circle.Stroke = sizeAttribute.ValueAsString(circle);
            }
        }

        private void DeserializeStrokeWidth(XElement holeElement, Circle circle)
        {
            if (XmlUtilities.TryGetAttribute(holeElement, "StrokeWidth", out XAttribute sizeAttribute))
            {
                circle.StrokeWidth = sizeAttribute.ValueAsString(circle);
            }
        }
    }
}
