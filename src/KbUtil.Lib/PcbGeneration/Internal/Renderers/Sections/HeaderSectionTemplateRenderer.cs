namespace KbUtil.Lib.PcbGeneration.Internal.Renderers.Sections
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Sections;
    using System.IO;

    internal class HeaderSectionTemplateRenderer : IPcbTemplateRenderer<HeaderSectionTemplateData>
    {
        private static readonly string _relativeTemplatePath =
            Path.Combine("PcbGeneration", "Internal", "Templates", "Sections", "header_section.template.kicad_pcb");

        public string Render(HeaderSectionTemplateData templateData)
            => File.ReadAllText(TemplatePath);

        private string TemplatePath => Path.Combine(Utilities.AssemblyDirectory, _relativeTemplatePath);
    }
}
