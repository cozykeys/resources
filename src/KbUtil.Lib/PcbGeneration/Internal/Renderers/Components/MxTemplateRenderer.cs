namespace KbUtil.Lib.PcbGeneration.Internal
{
    using KbUtil.Lib.PcbGeneration.Internal.Models.Components;
    using System.IO;

    internal class MxTemplateRenderer : IPcbTemplateRenderer<MxTemplateData>
    {
        public string KeyboardName { get; set; }

        private static readonly string _relativeTemplatePath =
            Path.Combine("PcbGeneration", "Internal", "Templates", "Components", "mx.template.kicad_pcb");
            //Path.Combine("PcbGeneration", "Internal", "Templates", "Components", "mx_raven.template.kicad_pcb");

        public string Render(MxTemplateData templateData)
            => File.ReadAllText(TemplatePath)
                .Replace("${x}", templateData.X.ToString())
                .Replace("${y}", templateData.Y.ToString())
                .Replace("${rotation}", templateData.Rotation.ToString())
                .Replace("${vertical_rotation}", (templateData.Rotation + 90).ToString())
                .Replace("${label}", templateData.Label)
                .Replace("${resistor_label}", templateData.ResistorLabel)
                .Replace("${diode_label}", templateData.DiodeLabel)
                .Replace("${diode_net_id}", templateData.DiodeNetId.ToString())
                .Replace("${diode_net_name}", templateData.DiodeNetName)
                .Replace("${vcc_net_id}", templateData.VccNetId.ToString())
                .Replace("${vcc_net_name}", templateData.VccNetName)
                .Replace("${led_net_id}", templateData.LedNetId.ToString())
                .Replace("${led_net_name}", templateData.LedNetName)
                .Replace("${mosfet_net_id}", templateData.MosfetNetId.ToString())
                .Replace("${mosfet_net_name}", templateData.MosfetNetName)
                .Replace("${row_net_id}", templateData.RowNetId.ToString())
                .Replace("${row_net_name}", templateData.RowNetName)
                .Replace("${column_net_id}", templateData.ColumnNetId.ToString())
                .Replace("${column_net_name}", templateData.ColumnNetName);

        private string TemplatePath => Path.Combine(Utilities.AssemblyDirectory, _relativeTemplatePath);
    }
}
