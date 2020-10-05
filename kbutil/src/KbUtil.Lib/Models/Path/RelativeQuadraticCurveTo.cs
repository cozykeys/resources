namespace KbUtil.Lib.Models.Path
{
    using Keyboard;
    using Geometry;

    public class RelativeQuadraticCurveTo : Element, IPathComponent
    {
        public Vec2 EndPoint { get; set; }
        public Vec2 ControlPoint { get; set; }

        public string Data => throw new System.NotImplementedException();
    }
}
