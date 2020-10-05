namespace KbUtil.Lib.Models.Path
{
    using Keyboard;
    using Geometry;

    public class AbsoluteMoveTo : Element, IPathComponent
    {
        public Vec2 EndPoint { get; set; }

        public string Data => $"M {EndPoint.X} {EndPoint.Y}";
    }
}
