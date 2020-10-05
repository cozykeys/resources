namespace KbMath.Core.Geometry2D.Extensions
{
    using System;

    public static class FloatExtensions
    {
        public static float ToRadians(this float degrees)
            => (float) ((Math.PI / 180) * degrees);
    }
}