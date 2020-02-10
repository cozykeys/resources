namespace KbUtil.Lib.PcbGeneration
{
    using KbUtil.Lib.Models.Keyboard;
    using KbUtil.Lib.Models.Pcb;
    using KbUtil.Lib.PcbGeneration.Internal.Models.Sections;
    using KbUtil.Lib.PcbGeneration.Internal.Renderers.Sections;
    using System;
    using System.Collections.Generic;
    using System.IO;

    public class PcbGenerator
    {
        public static void GeneratePcb(string keyboardName, List<Switch> switches, string outputPath, PcbGenerationOptions options = null)
        {
            var pcbData = new PcbData(switches);

            var pcb = new List<string>
            {
                "(kicad_pcb (version 3) (host pcbnew \"(2014-02-26 BZR 4721)-product\")",
                RenderHeaderSection(keyboardName),
                RenderNetsSection(keyboardName, pcbData),
                RenderControllerSection(keyboardName),
                RenderEdgesSection(keyboardName),
                RenderSwitchesSection(keyboardName, pcbData),
                RenderRgbUnderglowSection(keyboardName),
                RenderTracesSection(keyboardName),
                ")"
            };

            var pcbString = string.Join($"{Environment.NewLine}{Environment.NewLine}", pcb);

            File.WriteAllText(outputPath, pcbString);
        }


        private static string RenderHeaderSection(string keyboardName)
        {
            var headerTemplateData = new HeaderSectionTemplateData();
            var headerTemplateRenderer = new HeaderSectionTemplateRenderer
            {
                KeyboardName = keyboardName
            };
            return headerTemplateRenderer.Render(headerTemplateData);
        }

        private static string RenderNetsSection(string keyboardName, PcbData pcbData)
        {
            var templateRenderer = new NetsSectionTemplateRenderer
            {
                KeyboardName = keyboardName
            };
            return templateRenderer.Render(new NetsSectionTemplateData
            {
                NetDictionary = pcbData.NetDictionary
            });
        }

        private static string RenderControllerSection(string keyboardName)
        {
            var templateData = new ControllerSectionTemplateData();
            var templateRenderer = new ControllerSectionTemplateRenderer
            {
                KeyboardName = keyboardName
            };
            return templateRenderer.Render(templateData);
        }

        private static string RenderEdgesSection(string keyboardName)
        {
            var templateData = new EdgesSectionTemplateData();
            var templateRenderer = new EdgesSectionTemplateRenderer
            {
                KeyboardName = keyboardName
            };
            return templateRenderer.Render(templateData);
        }

        private static string RenderSwitchesSection(string keyboardName, PcbData pcbData)
        {
            var templateRenderer = new SwitchesSectionTemplateRenderer
            {
                KeyboardName = keyboardName
            };
            return templateRenderer.Render(new SwitchesSectionTemplateData
            {
                PcbData = pcbData
            });
        }

        private static string RenderRgbUnderglowSection(string keyboardName)
        {
            var templateRenderer = new RgbUnderglowSectionTemplateRenderer
            {
                KeyboardName = keyboardName
            };
            return templateRenderer.Render(new RgbUnderglowSectionTemplateData());
        }

        private static string RenderTracesSection(string keyboardName)
        {
            var templateRenderer = new TracesSectionTemplateRenderer
            {
                KeyboardName = keyboardName
            };
            return templateRenderer.Render(new TracesSectionTemplateData());
        }
    }
}
