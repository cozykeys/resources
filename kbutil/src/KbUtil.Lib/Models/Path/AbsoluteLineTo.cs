namespace KbUtil.Lib.Models.Path
{
    using Keyboard;
    using Geometry;

    public class AbsoluteLineTo : Element, IPathComponent
    {
        public Vec2 EndPoint { get; set; }

        public string Data => $"L {EndPoint.X} {EndPoint.Y}";
    }
}
