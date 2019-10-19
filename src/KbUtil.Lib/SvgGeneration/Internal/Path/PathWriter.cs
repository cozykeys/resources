// ReSharper disable IdentifierTypo
namespace KbUtil.Lib.SvgGeneration.Internal.Path
{
    using System.Xml;
    using KbUtil.Lib.Models.Path;
    using System.Collections.Generic;
    using KbUtil.Lib.Extensions;

    internal class PathWriter : IElementWriter<Path>
    {
        private static Dictionary<string, string> _pathStyleVisual = new Dictionary<string, string>
        {
            { "fill", "none" },
            { "stroke", "#0000ff" },
            { "stroke-width", "0.01" },
        };

        private static readonly Dictionary<string, string> PathStylePonoko = new Dictionary<string, string>
        {
            { "fill", "none" },
            { "stroke", "#0000ff" },
            { "stroke-width", "0.5" },
        };

        private static readonly string DefaultFill = "none";
        private static readonly string DefaultFillOpacity = "1";
        private static readonly string DefaultStroke = "#0000ff";
        private static readonly string DefaultStrokeWidth = "0.5";

        public SvgGenerationOptions GenerationOptions { get; set; }

        public void Write(XmlWriter writer, Path path)
        {
            writer.WriteStartElement("g");

            var elementWriter = new ElementWriter { GenerationOptions = GenerationOptions };

            elementWriter.WriteAttributes(writer, path);
            WriteAttributes(writer, path);

            elementWriter.WriteSubElements(writer, path);
            WriteSubElements(writer, path);

            writer.WriteEndElement();
        }

        public void WriteAttributes(XmlWriter writer, Path path)
        {
        }

        public void WriteSubElements(XmlWriter writer, Path path)
        {
            var styleDictionary = new Dictionary<string, string>();
            
            styleDictionary["fill"] = !string.IsNullOrEmpty(path.Fill)
                ? path.Fill
                : DefaultFill;
            styleDictionary["fill-opacity"] = !string.IsNullOrEmpty(path.FillOpacity)
                ? path.FillOpacity
                : DefaultFillOpacity;
            styleDictionary["stroke"] = !string.IsNullOrEmpty(path.Stroke)
                ? path.Stroke
                : DefaultStroke;
            styleDictionary["stroke-width"] = !string.IsNullOrEmpty(path.StrokeWidth)
                ? path.StrokeWidth
                : DefaultStrokeWidth;

            //WritePath(writer, path, _pathStyleVisual);
            WritePath(writer, path, styleDictionary);
        }

        private void WritePath(XmlWriter writer, Path path, Dictionary<string, string> styleDictionary)
        {
            writer.WriteStartElement("path");
            writer.WriteAttributeString("style", styleDictionary.ToCssStyleString());
            writer.WriteAttributeString("d", path.Data);
            writer.WriteEndElement();
        }
    }
}
