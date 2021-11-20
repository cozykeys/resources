namespace KbMath.Core.Geometry2D.Extensions
{
    using System;

    public static class DoubleExtensions
    {
        public static double ToRadians(this double degrees)
            => (Math.PI / 180) * degrees;

        public static bool Equals(this double a, double b, double precision)
            => Math.Abs(a - b) <= precision;
    }
}
