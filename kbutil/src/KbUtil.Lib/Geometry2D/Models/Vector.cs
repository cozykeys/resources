namespace KbMath.Core.Geometry2D.Models
{
    using System;

    public class Vector
    {
        public Vector(double x, double y)
        {
            X = x;
            Y = y;
        }
        
        public double X { get; }
        public double Y { get; }

        public override bool Equals(object obj)
            => obj is Vector rhs && X.Equals(rhs.X) && Y.Equals(rhs.Y);
        
        public override int GetHashCode()
        {
            int hash = 13;
            hash = (hash * 7) + X.GetHashCode();
            hash = (hash * 7) + Y.GetHashCode();
            return hash;
        }

        public Vector Round(int digits)
            => new Vector(Math.Round(X, digits), Math.Round(Y, digits));
        
        public static Vector operator+(Vector lhs, Vector rhs)
            => new Vector(lhs.X + rhs.X, lhs.Y + rhs.Y);

        public override string ToString()
            => $"Vector: X = {X}, Y = {Y}";
    }
}
