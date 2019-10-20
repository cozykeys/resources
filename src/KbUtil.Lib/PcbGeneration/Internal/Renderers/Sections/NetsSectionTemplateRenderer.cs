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
        public string KeyboardName { get; set; }

        private List<string> _universalNets = new List<string> {
            "N-5V-0",
            "N-GND-1",
            "N-RGB-D0",
            "N-RGB-D1",
            "N-RGB-D2",
            "N-RGB-D3",
            "N-RGB-D4",
            "N-RGB-D5",
            "N-RGB-D6",
            "N-RGB-D7",
            "N-RGB-D8",
            "N-RGB-D9",
            "N-RGB-D10",
            "N-RGB-D11",
            "N-LED-0"
        };

        private static readonly string _relativeTemplatePath =
            Path.Combine("PcbGeneration", "Internal", "Templates", "Sections", "nets_section.template.kicad_pcb");

        public string Render(NetsSectionTemplateData templateData)
            => File.ReadAllText(TemplatePath)
                .Replace("${nets}", RenderNets(templateData))
                .Replace("${net_class}", RenderNetClass(templateData));

        private string RenderNets(NetsSectionTemplateData sectionTemplateData)
        {
            var orderedNets = sectionTemplateData.NetDictionary.OrderBy(net => net.Value);

            var renderer = new NetTemplateRenderer
            {
                KeyboardName = KeyboardName
            };

            var nets = new List<string>();
            foreach (var net in orderedNets)
            {
                nets.Add(renderer.Render(new NetTemplateData
                {
                    Id = net.Value.ToString(),
                    Name = net.Key
                }));
            }

            int i = nets.Count();
            if (i != 82) throw new Exception($"i = {i}");
            foreach (var n in _universalNets)
            {
                nets.Add(renderer.Render(new NetTemplateData { Id = $"{i++}", Name = n }));
            }

            return string.Join(Environment.NewLine, nets);
        }

        public string RenderNetClass(NetsSectionTemplateData templateData)
        {
            var orderedNetNames = templateData.NetDictionary
                .OrderBy(net => net.Value)
                .Select(net => net.Key)
                .ToList();

            orderedNetNames.AddRange(_universalNets);

            var renderer = new NetClassTemplateRenderer
            {
                KeyboardName = KeyboardName
            };
            return renderer.Render(new NetClassTemplateData
            {
                NetNames = orderedNetNames
            });
        }

        private string TemplatePath => Path.Combine(Utilities.AssemblyDirectory, _relativeTemplatePath);
    }
}
