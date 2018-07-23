namespace KbUtil.Lib.PcbGeneration.Internal.Models.Components
{
    internal class DiodeTemplateData
    {
        public float X { get; set; }
        public float Y { get; set; }
        public float Rotation { get; set; }
        public string Label { get; set; }
        public int DiodeNetId { get; set; }
        public string DiodeNetName { get; set; }
        public int RowNetId { get; set; }
        public string RowNetName { get; set; }
    }
}
