namespace KbUtil.Lib.SvgGeneration
{
    using System.IO;
    using System.Xml;
    using KbUtil.Lib.Models.Keyboard;
    using KbUtil.Lib.SvgGeneration.Internal;
    using System;

    public class SvgGenerator
    {
        public static void GenerateSvg(Keyboard keyboard, string outputDirectory, SvgGenerationOptions options = null)
        {
            var settings = new XmlWriterSettings
            {
                Indent = true,
                IndentChars = options?.IndentString ?? "  ",
                NewLineOnAttributes = true
            };

            Directory.CreateDirectory(outputDirectory);

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

        private static void WriteSvgOpenTag(XmlWriter writer, int width, int height)
        {
            width = width > 0 ? width : 500;
            height = height > 0 ? height : 500;

            writer.WriteStartElement("svg", "http://www.w3.org/2000/svg");
            writer.WriteAttributeString("width", $"{width}mm");
            writer.WriteAttributeString("height", $"{height}mm");
            writer.WriteAttributeString("viewBox", $"0 0 {width} {height}");
        }

        private static void WriteSvgCloseTag(XmlWriter writer)
        {
            writer.WriteEndElement();
        }
    }
}
