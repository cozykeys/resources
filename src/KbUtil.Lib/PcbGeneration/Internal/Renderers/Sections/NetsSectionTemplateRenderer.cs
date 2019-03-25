namespace KbUtil.Lib.PcbGeneration.Internal.Renderers.Sections
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Components;
    using KbUtil.Lib.PcbGeneration.Internal.Models.Sections;
    using System;
    using System.Collections.Generic;
    using System.IO;
    using System.Linq;

    internal class NetsSectionTemplateRenderer : IPcbTemplateRenderer<NetsSectionTemplateData>
    {
        private static readonly string _relativeTemplatePath =
            Path.Combine("PcbGeneration", "Internal", "Templates", "Sections", "nets_section.template.kicad_pcb");

        public string Render(NetsSectionTemplateData templateData)
            => File.ReadAllText(TemplatePath)
                .Replace("${nets}", RenderNets(templateData))
                .Replace("${net_class}", RenderNetClass(templateData));

        private string RenderNets(NetsSectionTemplateData sectionTemplateData)
        {
            var orderedNets = sectionTemplateData.NetDictionary.OrderBy(net => net.Value);

            var renderer = new NetTemplateRenderer();

            var nets = new List<string>();
            foreach (var net in orderedNets)
            {
                nets.Add(renderer.Render(new NetTemplateData
                {
                    Id = net.Value.ToString(),
                    Name = net.Key
                }));
            }

            return string.Join(Environment.NewLine, nets);
        }

        public string RenderNetClass(NetsSectionTemplateData templateData)
        {
            var orderedNetNames = templateData.NetDictionary
                .OrderBy(net => net.Value)
                .Select(net => net.Key)
                .ToList();

            var renderer = new NetClassTemplateRenderer();
            return renderer.Render(new NetClassTemplateData
            {
                NetNames = orderedNetNames
            });
        }

        private string TemplatePath => Path.Combine(Utilities.AssemblyDirectory, _relativeTemplatePath);
    }
}
