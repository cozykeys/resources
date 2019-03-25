namespace KbUtil.Lib.PcbGeneration.Internal
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Components;
    using System.IO;

    internal class AddNetTemplateRenderer : IPcbTemplateRenderer<AddNetTemplateData>
    {
        private static readonly string _relativeTemplatePath =
            Path.Combine("PcbGeneration", "Internal", "Templates", "Components", "add_net.template.kicad_pcb");

        public string Render(AddNetTemplateData templateData)
            => File.ReadAllText(TemplatePath)
                .Replace("${name}", templateData.Name);

        private string TemplatePath => Path.Combine(Utilities.AssemblyDirectory, _relativeTemplatePath);
    }
}
