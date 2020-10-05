namespace KbUtil.Lib.PcbGeneration.Internal.Renderers.Sections
{
    using KbUtil.Lib.PcbGeneration.Internal.Extensions;
    using KbUtil.Lib.PcbGeneration.Internal.Models.Sections;
    using System.IO;

    internal class RgbUnderglowSectionTemplateRenderer : IPcbTemplateRenderer<RgbUnderglowSectionTemplateData>
    {
        public string KeyboardName { get; set; }

        public string Render(RgbUnderglowSectionTemplateData templateData)
        {
            string templatePath = Path.Combine(
                Utilities.AssemblyDirectory,
                "PcbGeneration",
                "Internal",
                "Templates",
                "Sections",
                $"rgb_underglow_section_{KeyboardName}.template.kicad_pcb");

            return File
                .ReadAllText(templatePath)
                .StripComments();
        }
    }
}
