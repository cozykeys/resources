namespace KbUtil.Lib.Tests.Comparers
{
    using System.Collections.Generic;
    using KbUtil.Lib.Geometry2D.Extensions;

    public class DoubleComparer : IEqualityComparer<double>
    {
        private const double Precision = 0.001;
        
        public bool Equals(double x, double y)
            => x.Equals(y, Precision);

        public int GetHashCode(double obj)
        {
            throw new System.NotImplementedException();
        }
    }
}