namespace KbUtil.Lib.Geometry2D.Models
{
    public class Line
    {
        public Line(double slope, double yIntercept)
        {
            Slope = slope;
            YIntercept = yIntercept;
        }
        
        public double Slope { get; }
        public double YIntercept { get; }

        public override bool Equals(object obj)
            => obj is Line rhs && Slope.Equals(rhs.Slope) && YIntercept.Equals(rhs.YIntercept);

        public override int GetHashCode()
        {
            int hash = 13;
            hash = (hash * 7) + Slope.GetHashCode();
            hash = (hash * 7) + YIntercept.GetHashCode();
            return hash;
        }

        public override string ToString()
            => $"Line: Slope = {Slope}, YIntercept = {YIntercept}";
    }
}
