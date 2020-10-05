namespace KbUtil.Lib.PcbGeneration.Internal.Renderers.Sections
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Sections;
    using System.IO;

    internal class HeaderSectionTemplateRenderer : IPcbTemplateRenderer<HeaderSectionTemplateData>
    {
        public string KeyboardName { get; set; }

        public string Render(HeaderSectionTemplateData templateData)
        {
            string templatePath = Path.Combine(
                Utilities.AssemblyDirectory,
                "PcbGeneration",
                "Internal",
                "Templates",
                "Sections",
                "header_section.template.kicad_pcb");

            return File.ReadAllText(templatePath);
        }
    }
}
