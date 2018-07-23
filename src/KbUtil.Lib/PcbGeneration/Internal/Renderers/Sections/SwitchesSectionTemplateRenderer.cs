namespace KbUtil.Lib.PcbGeneration.Internal.Renderers.Sections
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Sections;
    using System.IO;
    using System;
    using System.Collections.Generic;
    using KbUtil.Lib.PcbGeneration.Internal.Models.Components;

    internal class SwitchesSectionTemplateRenderer : IPcbTemplateRenderer<SwitchesSectionTemplateData>
    {
        private const string _relativeTemplatePath =
            @"PcbGeneration\Internal\Templates\Sections\switches_section.template.kicad_pcb";

        public string Render(SwitchesSectionTemplateData templateData)
            => File.ReadAllText(TemplatePath)
                .Replace("${switches}", RenderSwitches(templateData));

        public string RenderSwitches(SwitchesSectionTemplateData templateData)
        {
            var switches = new List<string>();

            var mxFlipRenderer = new MxFlipTemplateRenderer();
            var diodeRenderer = new DiodeTemplateRenderer();

            for (int i = 0; i < templateData.PcbData.RowCount; ++i)
            {
                for (int j = 0; j < templateData.PcbData.ColumnCount; ++j)
                {
                    if (templateData.PcbData.Switches[i][j] == null)
                    {
                        continue;
                    }

                    switches.Add(mxFlipRenderer.Render(new MxFlipTemplateData
                    {
                        X = templateData.PcbData.Switches[i][j].X,
                        Y = templateData.PcbData.Switches[i][j].Y,
                        Rotation = templateData.PcbData.Switches[i][j].Rotation,
                        Label = $"SW{i}:{j}",
                        DiodeNetId = templateData.PcbData.NetDictionary[$"N-diode-{i}-{j}"],
                        DiodeNetName = $"N-diode-{i}-{j}",
                        ColumnNetId = templateData.PcbData.NetDictionary[$"N-col-{j}"],
                        ColumnNetName = $"N-col-{j}"
                    }));
                    switches.Add(diodeRenderer.Render(new DiodeTemplateData
                    {
                        X = templateData.PcbData.Switches[i][j].X + 9,
                        Y = templateData.PcbData.Switches[i][j].Y,
                        Rotation = templateData.PcbData.Switches[i][j].Rotation + 90,
                        Label = $"D{i}:{j}",
                        DiodeNetId = templateData.PcbData.NetDictionary[$"N-diode-{i}-{j}"],
                        DiodeNetName = $"N-diode-{i}-{j}",
                        RowNetId = templateData.PcbData.NetDictionary[$"N-row-{i}"],
                        RowNetName = $"N-row-{i}"
                    }));
                }
            }

            return string.Join(Environment.NewLine, switches);
        }

        private string TemplatePath => Path.Combine(Utilities.AssemblyDirectory, _relativeTemplatePath);
    }
}
