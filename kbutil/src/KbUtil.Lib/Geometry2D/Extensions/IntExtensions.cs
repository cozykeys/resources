namespace KbMath.Core.Geometry2D.Extensions
{
    using System;

    public static class IntExtensions
    {
        public static double ToRadians(this int degrees)
            => ((Math.PI / 180) * degrees);
    }
}