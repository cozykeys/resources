namespace KbUtil.Lib.Models.Path
{
    using Keyboard;
    using Geometry;

    public class RelativeLineTo : Element, IPathComponent
    {
        public Vec2 EndPoint { get; set; }

        public string Data => throw new System.NotImplementedException();
    }
}
