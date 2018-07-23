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

        public static void GeneratePcb(List<Switch> switches, string outputDirectory, PcbGenerationOptions options = null)
        {
            var pcbData = new PcbData(switches);

            var pcb = new List<string>
            {
                "(kicad_pcb (version 3) (host pcbnew \"(2014-02-26 BZR 4721)-product\")",
                RenderHeaderSection(),
                RenderNetsSection(pcbData),
                RenderControllerSection(),
                RenderEdgesSection(),
                RenderSwitchesSection(pcbData),
                RenderTracesSection(),
                ")"
            };

            var pcbString = string.Join($"{Environment.NewLine}{Environment.NewLine}", pcb);

            Directory.CreateDirectory(outputDirectory);
            File.WriteAllText(Path.Combine(outputDirectory, "Ergo87.kicad_pcb"), pcbString);
        }


        private static string RenderHeaderSection()
        {
            var headerTemplateData = new HeaderSectionTemplateData();
            var headerTemplateRenderer = new HeaderSectionTemplateRenderer();
            return headerTemplateRenderer.Render(headerTemplateData);
        }

        private static string RenderNetsSection(PcbData pcbData)
        {
            var templateRenderer = new NetsSectionTemplateRenderer();
            return templateRenderer.Render(new NetsSectionTemplateData
            {
                NetDictionary = pcbData.NetDictionary
            });
        }

        private static string RenderControllerSection()
        {
            var templateData = new ControllerSectionTemplateData();
            var templateRenderer = new ControllerSectionTemplateRenderer();
            return templateRenderer.Render(templateData);
        }

        private static string RenderEdgesSection()
        {
            var templateData = new EdgesSectionTemplateData();
            var templateRenderer = new EdgesSectionTemplateRenderer();
            return templateRenderer.Render(templateData);
        }

        private static string RenderSwitchesSection(PcbData pcbData)
        {
            var templateRenderer = new SwitchesSectionTemplateRenderer();
            return templateRenderer.Render(new SwitchesSectionTemplateData
            {
                PcbData = pcbData
            });
        }

        private static string RenderTracesSection()
        {
            var templateRenderer = new TracesSectionTemplateRenderer();
            return templateRenderer.Render(new TracesSectionTemplateData());
        }
    }
}
