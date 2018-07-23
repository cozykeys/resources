namespace KbUtil.Lib.PcbGeneration.Internal
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Components;
    using System.IO;

    internal class AddNetTemplateRenderer : IPcbTemplateRenderer<AddNetTemplateData>
    {
        private const string _relativeTemplatePath =
            @"PcbGeneration\Internal\Templates\Components\add_net.template.kicad_pcb";

        public string Render(AddNetTemplateData templateData)
            => File.ReadAllText(TemplatePath)
                .Replace("${name}", templateData.Name);

        private string TemplatePath => Path.Combine(Utilities.AssemblyDirectory, _relativeTemplatePath);
    }
}
