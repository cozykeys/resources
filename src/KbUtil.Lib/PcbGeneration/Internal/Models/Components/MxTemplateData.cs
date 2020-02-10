namespace KbUtil.Lib.PcbGeneration.Internal.Models.Components
{
    internal class MxTemplateData
    {
        public float X { get; set; }
        public float Y { get; set; }
        public float Rotation { get; set; }
        public string Label { get; set; }
        public string ResistorLabel { get; set; }
        public string DiodeLabel { get; set; }
        public int DiodeNetId { get; set; }
        public string DiodeNetName { get; set; }
        public int LedNetId { get; set; }
        public string LedNetName { get; set; }
        public int MosfetNetId { get; set; }
        public string MosfetNetName { get; set; }
        public int VccNetId { get; set; }
        public string VccNetName { get; set; }
        public int ColumnNetId { get; set; }
        public string ColumnNetName { get; set; }
        public int RowNetId { get; set; }
        public string RowNetName { get; set; }
    }
}
