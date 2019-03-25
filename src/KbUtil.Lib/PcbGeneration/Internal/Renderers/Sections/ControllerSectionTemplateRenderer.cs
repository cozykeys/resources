namespace KbUtil.Lib.PcbGeneration.Internal.Renderers.Sections
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Sections;
    using System.IO;

    internal class ControllerSectionTemplateRenderer : IPcbTemplateRenderer<ControllerSectionTemplateData>
    {
        private static readonly string _relativeTemplatePath =
            Path.Combine("PcbGeneration", "Internal", "Templates", "Sections", "controller_section.template.kicad_pcb");

        public string Render(ControllerSectionTemplateData templateData)
        {
            return File.ReadAllText(TemplatePath);
        }

        private string TemplatePath => Path.Combine(Utilities.AssemblyDirectory, _relativeTemplatePath);
    }
}
