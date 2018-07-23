namespace KbUtil.Lib.PcbGeneration.Internal.Renderers.Sections
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Sections;
    using System.IO;

    internal class TracesSectionTemplateRenderer : IPcbTemplateRenderer<TracesSectionTemplateData>
    {
        private const string _relativeTemplatePath =
            @"PcbGeneration\Internal\Templates\Sections\traces_section.template.kicad_pcb";

        public string Render(TracesSectionTemplateData templateData)
            => File.ReadAllText(TemplatePath);

        private string TemplatePath => Path.Combine(Utilities.AssemblyDirectory, _relativeTemplatePath);
    }
}
