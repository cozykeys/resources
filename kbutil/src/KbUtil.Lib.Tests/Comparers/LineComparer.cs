namespace KbUtil.Lib.Tests.Comparers
{
    using System.Collections.Generic;
    using KbUtil.Lib.Geometry2D.Models;

    public class LineComparer : IEqualityComparer<Line>
    {
        private static readonly DoubleComparer DoubleComparer = new DoubleComparer();

        public bool Equals(Line a, Line b)
            => DoubleComparer.Equals(a.Slope, b.Slope) && DoubleComparer.Equals(a.YIntercept, b.YIntercept);

        public int GetHashCode(Line obj)
        {
            throw new System.NotImplementedException();
        }
    }
}