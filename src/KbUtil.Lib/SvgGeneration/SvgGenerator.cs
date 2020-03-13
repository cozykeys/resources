namespace KbUtil.Lib.SvgGeneration
{
    using System.Collections.Generic;
    using System.IO;
    using System.Linq;
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

            if (options != null && options.SquashLayers)
            {
                GenerateLayersSquashed(keyboard, outputDirectory, options, settings);
            }
            else
            {
                GenerateLayers(keyboard, outputDirectory, options, settings);
            }
        }

        private static void GenerateLayersSquashed(Keyboard keyboard, string outputDirectory, SvgGenerationOptions options, XmlWriterSettings settings)
        {
            string path = System.IO.Path.Combine(outputDirectory, $"{keyboard.Name}.svg");

            using (FileStream stream = File.Open(path, FileMode.Create))
            using (XmlWriter writer = XmlWriter.Create(stream, settings))
            {
                int maxWidth = -1;
                int maxHeight = -1;

                foreach (var layer in keyboard.Layers)
                {
                    if ((int)layer.Width > maxWidth)
                        maxWidth = (int)layer.Width;
                    if ((int)layer.Height > maxHeight)
                        maxHeight = (int)layer.Height;
                }

                WriteSvgOpenTag(writer, maxWidth, maxHeight);

                var orderedLayers = keyboard.Layers.OrderBy(l => l.ZIndex).ToList();
                foreach (var layer in orderedLayers)
                {
                    var layerWriter = new LayerWriter { GenerationOptions = options };
                    layerWriter.Write(writer, layer);
                }

                WriteSvgCloseTag(writer);
            }
        }

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
