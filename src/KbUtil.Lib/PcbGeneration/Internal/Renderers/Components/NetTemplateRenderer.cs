namespace KbUtil.Lib.PcbGeneration.Internal
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Components;
    using System.IO;

    internal class NetTemplateRenderer : IPcbTemplateRenderer<NetTemplateData>
    {
        private const string _relativeTemplatePath =
            @"PcbGeneration\Internal\Templates\Components\net.template.kicad_pcb";

        public string Render(NetTemplateData templateData)
            => File.ReadAllText(TemplatePath)
                .Replace("${id}", templateData.Id)
                .Replace("${name}", templateData.Name);

        private string TemplatePath => Path.Combine(Utilities.AssemblyDirectory, _relativeTemplatePath);
    }
}
