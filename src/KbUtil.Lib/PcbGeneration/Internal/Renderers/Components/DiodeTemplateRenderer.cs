namespace KbUtil.Lib.PcbGeneration.Internal
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Components;
    using System.IO;

    internal class DiodeTemplateRenderer : IPcbTemplateRenderer<DiodeTemplateData>
    {
        private const string _relativeTemplatePath =
            @"PcbGeneration\Internal\Templates\Components\diode.template.kicad_pcb";

        public string Render(DiodeTemplateData templateData)
            => File.ReadAllText(TemplatePath)
                .Replace("${label}", templateData.Label)
                .Replace("${x}", templateData.X.ToString())
                .Replace("${y}", templateData.Y.ToString())
                .Replace("${rotation}", templateData.Rotation.ToString())
                .Replace("${pad_rotation}", templateData.PadRotation.ToString())
                .Replace("${diode_net_id}", templateData.DiodeNetId.ToString())
                .Replace("${diode_net_name}", templateData.DiodeNetName)
                .Replace("${row_net_id}", templateData.RowNetId.ToString())
                .Replace("${row_net_name}", templateData.RowNetName);

        private string TemplatePath => Path.Combine(Utilities.AssemblyDirectory, _relativeTemplatePath);
    }
}
