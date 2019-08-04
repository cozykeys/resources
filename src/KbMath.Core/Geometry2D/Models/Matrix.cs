using System;

namespace KbMath.Core.Geometry2D.Models
{
    public class Matrix2x2
    {
        public Matrix2x2(double[,] data)
        {
            if (data.Length != 4 || data.Rank != 2)
            {
                throw new Exception("Invalid data to instantiate a 2 x 2 matrix");
            }
            
            Data = data;
        }
        
        public double[,] Data { get; }

        public override bool Equals(object obj)
        {
            if (obj is Matrix2x2 rhs)
            {
                int elementsPerDimension = Data.Length / Data.Rank;
                for (int i = 0; i < Data.Rank; ++i)
                {
                    for (int j = 0; j < elementsPerDimension; ++j)
                    {
                        if (Math.Abs(Data[i, j] - rhs.Data[i, j]) > 0.01)
                        {
                            return false;
                        }
                    }
                }

                return true;
            }

            return false;
        }

        public static Vector operator*(Matrix2x2 lhs, Vector rhs)
            => new Vector(
                    (lhs.Data[0,0] * rhs.X) + (lhs.Data[1,0] * rhs.Y),
                    (lhs.Data[0,1] * rhs.X) + (lhs.Data[1,1] * rhs.Y)
                );
        
        public override int GetHashCode()
        {
            int hash = 13;
            hash = (hash * 7) + Data.GetHashCode();
            return hash;
        }

        public override string ToString()
            => $"Matrix:\n|   {Data[0,0]}   {Data[1,0]}   |\n|   {Data[0,1]}   {Data[1,1]}   |";
    }
}
