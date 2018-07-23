namespace KbUtil.Lib.PcbGeneration
{
    using KbUtil.Lib.Models.Keyboard;
    using KbUtil.Lib.Models.Pcb;
    using KbUtil.Lib.PcbGeneration.Internal.Models.Sections;
    using KbUtil.Lib.PcbGeneration.Internal.Renderers.Sections;
    using System;
    using System.Collections.Generic;

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
                //RenderEdgesSection(keyboard),
                RenderSwitchesSection(pcbData),
                ")"
            };

            var pcbString = string.Join($"{Environment.NewLine}{Environment.NewLine}", pcb);
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

        private static string RenderEdgesSection(Keyboard keyboard)
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

        private static IEnumerable<Key> GetKeys(Group group)
        {
            var keys = new List<Key>();
            foreach (var child in group.Children)
            {
                if (child is Key)
                {
                    keys.Add((Key)child);
                }
                else if (child is Group)
                {
                    keys.AddRange(GetKeys((Group)child));
                }
            }
            return keys;
        }

        public static string GetKeyPosition(Key key)
        {
            double xRelative = key.XOffset;
            double yRelative = key.YOffset;

            Element iter = key.Parent;
            while (!(iter is Layer))
            {
                if (iter.Rotation >= 0.00001)
                {
                    var xRotated = xRelative * Math.Cos(iter.Rotation) - yRelative * Math.Sin(iter.Rotation);
                    var yRotated = xRelative * Math.Sin(iter.Rotation) + yRelative * Math.Cos(iter.Rotation);

                    xRelative = iter.XOffset + xRotated;
                    yRelative = iter.YOffset + yRotated;
                }
                else
                {
                    xRelative = iter.XOffset + xRelative;
                    yRelative = iter.YOffset + yRelative;
                }

                iter = iter.Parent;
            }

            double xAbsolute = xRelative;
            double yAbsolute = yRelative;

            return $"({xAbsolute}, {yAbsolute})";
        }

    }
}
