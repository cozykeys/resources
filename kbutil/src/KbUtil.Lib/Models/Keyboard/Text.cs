namespace KbUtil.Lib.Models.Keyboard
{
    using KbUtil.Lib.Models.Attributes;

    [GroupChild]
    public class Text : Element
    {
        public string Content { get; set; }
        public string TextAnchor { get; set; }
        public string Font { get; set; }
        public string Fill { get; set; }
    }
}
