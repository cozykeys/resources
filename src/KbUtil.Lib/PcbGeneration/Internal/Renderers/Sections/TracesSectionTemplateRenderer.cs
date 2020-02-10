namespace KbUtil.Lib.PcbGeneration.Internal.Renderers.Sections
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Sections;
    using System.IO;
    using System.Text.RegularExpressions;

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

            string raw = File.ReadAllText(templatePath);

            // TODO: Move this into a Utility class
            System.Console.WriteLine("Stripping comments");
            string processed = Regex.Replace(raw, @"/\*.*\*/", "");

            return processed;
        }
    }
}
