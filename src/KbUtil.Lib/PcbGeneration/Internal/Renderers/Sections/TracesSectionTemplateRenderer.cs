namespace KbUtil.Lib.PcbGeneration.Internal.Renderers.Sections
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Sections;
    using System.IO;

    internal class TracesSectionTemplateRenderer : IPcbTemplateRenderer<TracesSectionTemplateData>
    {
        public string KeyboardName { get; set; }

        public string Render(TracesSectionTemplateData templateData)
        {
            string templatePath = Path.Combine(
                Utilities.AssemblyDirectory,
                "PcbGeneration",
                "Internal",
                "Templates",
                "Sections",
                $"traces_section_{KeyboardName}.template.kicad_pcb");

            return File.ReadAllText(templatePath);
        }
    }
}
