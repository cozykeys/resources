namespace KbMath.Core.Geometry2D.Models
{
    public class Ray
    {
        public Ray(Vector start, Vector direction)
        {
            Start = start;
            Direction = direction;
        }
        
        public Vector Start { get; }
        public Vector Direction { get; }

        public override bool Equals(object obj)
            => obj is Ray rhs && Start.Equals(rhs.Start) && Direction.Equals(rhs.Direction);

        public override int GetHashCode()
        {
            int hash = 13;
            hash = (hash * 7) + Start.GetHashCode();
            hash = (hash * 7) + Direction.GetHashCode();
            return hash;
        }
        
        public override string ToString()
            => $"Ray: Start = {Start}, Direction = {Direction}";
    }
}
