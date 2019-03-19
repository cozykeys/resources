namespace KbMath.Core.Geometry2D.Extensions
{
    using KbMath.Core.Common.Extensions;
    using KbMath.Core.Geometry2D.Models;

    public static class VectorExtensions
    {
        public static double Magnitude(this Vector vector)
            => (vector.X.Pow(2) + vector.Y.Pow(2)).Sqrt();

        public static Vector Add(this Vector lhs, Vector rhs)
            => new Vector(lhs.X + rhs.X, lhs.Y + rhs.Y);
    }
}
