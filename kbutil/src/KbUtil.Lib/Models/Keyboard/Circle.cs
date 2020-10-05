namespace KbUtil.Lib.Models.Keyboard
{
    using KbUtil.Lib.Models.Attributes;

    [GroupChild]
    public class Circle : Element
    {
        public float Size { get; set; }
        public string Fill { get; set; }
        public string Stroke { get; set; }
        public string StrokeWidth { get; set; }

        public override float Height => Size;
        public override float Width => Size;
    }
}
