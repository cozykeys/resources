namespace KbUtil.Lib.PcbGeneration.Internal
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Components;
    using System.IO;

    internal class NetTemplateRenderer : IPcbTemplateRenderer<NetTemplateData>
    {
        public string KeyboardName { get; set; }

        private static readonly string _relativeTemplatePath =
            Path.Combine("PcbGeneration", "Internal", "Templates", "Components", "net.template.kicad_pcb");

        public string Render(NetTemplateData templateData)
            => File.ReadAllText(TemplatePath)
                .Replace("${id}", templateData.Id)
                .Replace("${name}", templateData.Name);

        private string TemplatePath => Path.Combine(Utilities.AssemblyDirectory, _relativeTemplatePath);
    }
}
