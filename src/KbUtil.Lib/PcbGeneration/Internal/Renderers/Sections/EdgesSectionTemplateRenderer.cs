﻿namespace KbUtil.Lib.PcbGeneration.Internal.Renderers.Sections
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Sections;
    using System.IO;

    internal class EdgesSectionTemplateRenderer : IPcbTemplateRenderer<EdgesSectionTemplateData>
    {
        // Change this depending on the keyboard
        private static readonly string _relativeTemplatePath =
            //Path.Combine("PcbGeneration", "Internal", "Templates", "Sections", "edges_section.template.kicad_pcb");
            Path.Combine("PcbGeneration", "Internal", "Templates", "Sections", "edges_section_radix.template.kicad_pcb");

        public string Render(EdgesSectionTemplateData templateData)
        {
            return File.ReadAllText(TemplatePath);
        }

        private string TemplatePath => Path.Combine(Utilities.AssemblyDirectory, _relativeTemplatePath);
    }
}
