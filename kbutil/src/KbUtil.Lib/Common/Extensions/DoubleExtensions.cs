namespace KbUtil.Lib.Common.Extensions
{
    using System;

    public static class DoubleExtensions
    {
        public static double Mod(this double d, double b)
            => (d - b * Math.Floor(d / b));
        
        public static double Pow(this double d, double e)
            => Math.Pow(d, e);
        
        public static double Sqrt(this double d)
            => Math.Sqrt(d);
    }
}
