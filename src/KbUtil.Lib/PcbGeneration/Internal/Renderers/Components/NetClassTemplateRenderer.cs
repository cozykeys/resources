namespace KbUtil.Lib.PcbGeneration.Internal
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Components;
    using System;
    using System.IO;
    using System.Linq;

    internal class NetClassTemplateRenderer : IPcbTemplateRenderer<NetClassTemplateData>
    {
        private const string _relativeTemplatePath =
            @"PcbGeneration\Internal\Templates\Components\net_class.template.kicad_pcb";

        public string Render(NetClassTemplateData templateData)
            => File.ReadAllText(TemplatePath)
                .Replace("${add_nets}", RenderAddNets(templateData));

        private string RenderAddNets(NetClassTemplateData templateData)
        {
            var renderer = new AddNetTemplateRenderer();
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
