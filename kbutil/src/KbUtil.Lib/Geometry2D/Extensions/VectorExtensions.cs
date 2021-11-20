namespace KbUtil.Lib.Geometry2D.Extensions
{
    using System;
    using KbUtil.Lib.Common.Extensions;
    using KbUtil.Lib.Geometry2D.Models;

    public static class VectorExtensions
    {
        public static double Magnitude(this Vector vector)
            => (vector.X.Pow(2) + vector.Y.Pow(2)).Sqrt();

        public static Vector Add(this Vector lhs, Vector rhs)
            => new Vector(lhs.X + rhs.X, lhs.Y + rhs.Y);

        public static Vector Project(this Vector vector, double theta, double magnitude)
            => new Vector(vector.X + magnitude * Math.Cos(theta), vector.Y + magnitude * Math.Sin(theta));
    }
}
