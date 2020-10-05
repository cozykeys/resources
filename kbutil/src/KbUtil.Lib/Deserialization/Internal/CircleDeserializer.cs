namespace KbUtil.Lib.Deserialization.Internal
{
    using KbUtil.Lib.Deserialization.Extensions;
    using KbUtil.Lib.Models.Keyboard;
    using System.Xml.Linq;

    internal class CircleDeserializer : IDeserializer<Circle>
    {
        public static CircleDeserializer Default { get; set; } = new CircleDeserializer();

        public void Deserialize(XElement circleElement, Circle circle)
        {
            ElementDeserializer.Default.Deserialize(circleElement, circle);

            DeserializeSize(circleElement, circle);
            DeserializeFill(circleElement, circle);
            DeserializeStroke(circleElement, circle);
            DeserializeStrokeWidth(circleElement, circle);
        }

        private void DeserializeSize(XElement circleElement, Circle circle)
        {
            if (XmlUtilities.TryGetAttribute(circleElement, "Size", out XAttribute sizeAttribute))
            {
                circle.Size = sizeAttribute.ValueAsFloat(circle);
            }
        }

        private void DeserializeFill(XElement circleElement, Circle circle)
        {
            if (XmlUtilities.TryGetAttribute(circleElement, "Fill", out XAttribute sizeAttribute))
            {
                circle.Fill = sizeAttribute.ValueAsString(circle);
            }
        }

        private void DeserializeStroke(XElement circleElement, Circle circle)
        {
            if (XmlUtilities.TryGetAttribute(circleElement, "Stroke", out XAttribute sizeAttribute))
            {
                circle.Stroke = sizeAttribute.ValueAsString(circle);
            }
        }

        private void DeserializeStrokeWidth(XElement circleElement, Circle circle)
        {
            if (XmlUtilities.TryGetAttribute(circleElement, "StrokeWidth", out XAttribute sizeAttribute))
            {
                circle.StrokeWidth = sizeAttribute.ValueAsString(circle);
            }
        }
    }
}
