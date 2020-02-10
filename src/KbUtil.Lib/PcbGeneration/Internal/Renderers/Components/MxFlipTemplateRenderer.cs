namespace KbUtil.Lib.PcbGeneration.Internal
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Components;
    using System.IO;

    internal class MxFlipTemplateRenderer : IPcbTemplateRenderer<MxTemplateData>
    {
        public string KeyboardName { get; set; }

        private static readonly string _relativeTemplatePath =
            Path.Combine("PcbGeneration", "Internal", "Templates", "Components", "mx_flip.template.kicad_pcb");

        public string Render(MxTemplateData templateData)
            => File.ReadAllText(TemplatePath)
                .Replace("${x}", templateData.X.ToString())
                .Replace("${y}", templateData.Y.ToString())
                .Replace("${rotation}", templateData.Rotation.ToString())
                .Replace("${label}", templateData.Label)
                .Replace("${diode_net_id}", templateData.DiodeNetId.ToString())
                .Replace("${diode_net_name}", templateData.DiodeNetName)
                .Replace("${column_net_id}", templateData.ColumnNetId.ToString())
                .Replace("${column_net_name}", templateData.ColumnNetName);

        private string TemplatePath => Path.Combine(Utilities.AssemblyDirectory, _relativeTemplatePath);
    }
}
