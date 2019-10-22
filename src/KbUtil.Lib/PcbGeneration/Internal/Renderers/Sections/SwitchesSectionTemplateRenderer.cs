namespace KbUtil.Lib.PcbGeneration.Internal.Renderers.Sections
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Sections;
    using System.IO;
    using System;
    using System.Collections.Generic;
    using KbUtil.Lib.PcbGeneration.Internal.Models.Components;
    using KbMath.Core.Geometry2D.Extensions;

    internal class SwitchesSectionTemplateRenderer : IPcbTemplateRenderer<SwitchesSectionTemplateData>
    {
        public string KeyboardName { get; set; }

        private static readonly string _relativeTemplatePath =
            Path.Combine("PcbGeneration", "Internal", "Templates", "Sections", "switches_section.template.kicad_pcb");

        public string Render(SwitchesSectionTemplateData templateData)
            => File.ReadAllText(TemplatePath)
                .Replace("${switches}", RenderSwitches(templateData));

        public string RenderSwitches(SwitchesSectionTemplateData templateData)
        {
            var switches = new List<string>();

            // Swap this if the board needs flippable MX switch footprints
            var mxRenderer = new MxTemplateRenderer
            {
                KeyboardName = KeyboardName
            };
            //var mxRenderer = new MxFlipTemplateRenderer
            //{
            //    KeyboardName = KeyboardName
            //};
            var diodeRenderer = new DiodeTemplateRenderer
            {
                KeyboardName = KeyboardName
            };

            for (int i = 0; i < templateData.PcbData.RowCount; ++i)
            {
                for (int j = 0; j < templateData.PcbData.ColumnCount; ++j)
                {
                    if (templateData.PcbData.Switches[i][j] == null)
                    {
                        continue;
                    }

                    var switchLabel = $"SW{i}:{j}";
                    var switchX = templateData.PcbData.Switches[i][j].X;
                    var switchY = templateData.PcbData.Switches[i][j].Y;
                    var switchRotation = 0 - templateData.PcbData.Switches[i][j].Rotation;

                    var diodeLabel = $"D{i}:{j}";
                    float diodeX;
                    float diodeY;
                    float diodeRotation;
                    if (templateData.PcbData.Switches[i][j].DiodePosition == "left")
                    {
                        // TODO: Flip this on and fix trace generation. 
                        double d = 8.0;
                        double theta = ((double)10 + 90).ToRadians();
                        double dx = d * Math.Sin(theta);
                        double dy = d * Math.Cos(theta);

                        diodeY = (float)(switchY - dy);
                        diodeX = (float)(switchX - dx);

                        //diodeY = switchY;
                        //diodeX = switchX - 9 - templateData.PcbData.Switches[i][j].DiodeAdjust;
                        diodeRotation = switchRotation + 90;
                    }
                    else if (templateData.PcbData.Switches[i][j].DiodePosition == "right")
                    {
                        // TODO: Flip this on and fix trace generation. 
                        double d = 8.0;
                        double theta = ((double)10 + 90).ToRadians();
                        double dx = d * Math.Sin(theta);
                        double dy = d * Math.Cos(theta);

                        diodeY = (float)(switchY - dy);
                        diodeX = (float)(switchX + dx);

                        //diodeY = switchY;
                        //diodeX = switchX + 9 + templateData.PcbData.Switches[i][j].DiodeAdjust;
                        diodeRotation = switchRotation + 90;
                    }
                    else if (templateData.PcbData.Switches[i][j].DiodePosition == "top")
                    {
                        diodeRotation = switchRotation;
                        diodeY = switchY - 9 - templateData.PcbData.Switches[i][j].DiodeAdjust;
                        diodeX = switchX;
                    }
                    else if (templateData.PcbData.Switches[i][j].DiodePosition == "bottom")
                    {
                        diodeRotation = switchRotation;
                        diodeY = switchY + 9 + templateData.PcbData.Switches[i][j].DiodeAdjust;
                        diodeX = switchX;
                    }
                    else
                    {
                        throw new Exception("Invalid diode position");
                    }

                    var diodePadRotation = switchRotation;

                    var diodeNetName = $"N-diode-{i}-{j}";
                    var diodeNetId = templateData.PcbData.NetDictionary[diodeNetName];
                    var columnNetName = $"N-col-{j}";
                    var columnNetId = templateData.PcbData.NetDictionary[columnNetName];
                    var rowNetName = $"N-row-{i}";
                    var rowNetId = templateData.PcbData.NetDictionary[rowNetName];

                    switches.Add(mxRenderer.Render(new MxTemplateData
                    {
                        Label = switchLabel,
                        X = switchX,
                        Y = switchY,
                        Rotation = switchRotation,
                        DiodeNetId = diodeNetId,
                        DiodeNetName = diodeNetName,
                        ColumnNetId = columnNetId,
                        ColumnNetName = columnNetName
                    }));

                    switches.Add(diodeRenderer.Render(new DiodeTemplateData
                    {
                        Label = diodeLabel,
                        X = diodeX,
                        Y = diodeY,
                        Rotation = diodeRotation,
                        PadRotation = diodePadRotation,
                        DiodeNetId = diodeNetId,
                        DiodeNetName = diodeNetName,
                        RowNetId = rowNetId,
                        RowNetName = rowNetName
                    }));
                }
            }

            return string.Join(Environment.NewLine, switches);
        }

        private string TemplatePath => Path.Combine(Utilities.AssemblyDirectory, _relativeTemplatePath);
    }
}
