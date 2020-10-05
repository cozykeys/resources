namespace KbUtil.Lib.Models.Path
{
    using Keyboard;

    public class AbsoluteVerticalLineTo : Element, IPathComponent
    {
        public float Y { get; set; }

        public string Data => throw new System.NotImplementedException();
    }
}
