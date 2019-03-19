namespace KbMath.Core.Common.Extensions
{
    using System;

    public static class FloatExtensions
    {
        public static double ToDouble(this float f)
            => f;

        public static double Mod(this float f, float b)
            => (f - b * Math.Floor(f / b));
        
        public static double Pow(this float f, float e)
            => Math.Pow(f, e);
        
        public static double Sqrt(this float f)
            => Math.Sqrt(f);
    }
}
