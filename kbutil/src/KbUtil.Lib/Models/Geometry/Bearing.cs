namespace KbUtil.Lib.Models.Geometry
{
    using KbMath.Core.Geometry2D.Models;

    public class Bearing
    {
        public Vector Position { get; set; }
        public float Rotation { get; set; }
        public int Row { get; set; }
        public int Column { get; set; }
    }
}
