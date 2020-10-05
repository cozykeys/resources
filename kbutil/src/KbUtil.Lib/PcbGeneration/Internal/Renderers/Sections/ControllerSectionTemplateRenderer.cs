namespace KbUtil.Lib.PcbGeneration.Internal.Renderers.Sections
{
    using KbUtil.Lib.PcbGeneration.Internal.Extensions;
    using KbUtil.Lib.PcbGeneration.Internal.Models.Sections;
    using System.IO;

    internal class ControllerSectionTemplateRenderer : IPcbTemplateRenderer<ControllerSectionTemplateData>
    {
        public string KeyboardName { get; set; }

        public string Render(ControllerSectionTemplateData templateData)
        {
            string templatePath = Path.Combine(
                Utilities.AssemblyDirectory,
                "PcbGeneration",
                "Internal",
                "Templates",
                "Sections",
                $"controller_section_{KeyboardName}.template.kicad_pcb");

            return File
                .ReadAllText(templatePath)
                .StripComments();
        }
    }
}
