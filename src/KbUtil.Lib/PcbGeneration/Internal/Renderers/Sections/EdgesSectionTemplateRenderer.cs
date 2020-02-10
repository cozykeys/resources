namespace KbUtil.Lib.PcbGeneration.Internal.Renderers.Sections
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Sections;
    using System.IO;

    internal class EdgesSectionTemplateRenderer : IPcbTemplateRenderer<EdgesSectionTemplateData>
    {
        public string KeyboardName { get; set; }

        public string Render(EdgesSectionTemplateData templateData)
        {
            string templatePath = Path.Combine(
                Utilities.AssemblyDirectory,
                "PcbGeneration",
                "Internal",
                "Templates",
                "Sections",
                $"edges_section_{KeyboardName}.template.kicad_pcb");

            return File.ReadAllText(templatePath);
        }
    }
}
