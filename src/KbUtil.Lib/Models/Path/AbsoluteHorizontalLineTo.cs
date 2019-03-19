namespace KbUtil.Lib.Models.Path
{
    using Keyboard;

    public class AbsoluteHorizontalLineTo : Element, IPathComponent
    {
        public float X { get; set; }

        public string Data => throw new System.NotImplementedException();
    }
}
