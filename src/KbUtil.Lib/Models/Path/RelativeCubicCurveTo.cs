namespace KbUtil.Lib.Models.Path
{
    using Keyboard;
    using Geometry;

    public class RelativeCubicCurveTo : Element, IPathComponent
    {
        public Vec2 EndPoint { get; set; }
        public Vec2 ControlPointA { get; set; }
        public Vec2 ControlPointB { get; set; }

        public string Data => throw new System.NotImplementedException();
    }
}
