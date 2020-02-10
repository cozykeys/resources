namespace KbMath.Core.Tests.Comparers
{
    using System.Collections.Generic;
    using KbMath.Core.Geometry2D.Models;

    public class VectorComparer : IEqualityComparer<Vector>
    {
        private static readonly DoubleComparer DoubleComparer = new DoubleComparer();

        public bool Equals(Vector a, Vector b)
            => DoubleComparer.Equals(a.X, b.X) && DoubleComparer.Equals(a.Y, b.Y);

        public int GetHashCode(Vector obj)
        {
            throw new System.NotImplementedException();
        }
    }
}