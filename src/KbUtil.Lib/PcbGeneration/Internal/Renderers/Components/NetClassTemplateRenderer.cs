namespace KbUtil.Lib.PcbGeneration.Internal
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Components;
    using System;
    using System.IO;
    using System.Linq;

    internal class NetClassTemplateRenderer : IPcbTemplateRenderer<NetClassTemplateData>
    {
        public string KeyboardName { get; set; }

        private static readonly string _relativeTemplatePath =
            Path.Combine("PcbGeneration", "Internal", "Templates", "Components", "net_class.template.kicad_pcb");

        public string Render(NetClassTemplateData templateData)
            => File.ReadAllText(TemplatePath)
                .Replace("${add_nets}", RenderAddNets(templateData));

        private string RenderAddNets(NetClassTemplateData templateData)
        {
            var renderer = new AddNetTemplateRenderer
            {
                KeyboardName = KeyboardName
            };
            var addNets = templateData.NetNames.Select(netName => renderer.Render(
                new AddNetTemplateData
                {
                    Name = netName
                }));

            return string.Join(Environment.NewLine, addNets);
        }

        private string TemplatePath => Path.Combine(Utilities.AssemblyDirectory, _relativeTemplatePath);
    }
}
