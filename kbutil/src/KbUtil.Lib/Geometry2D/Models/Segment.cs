namespace KbUtil.Lib.Geometry2D.Models
{
    public class Segment
    {
        public Segment(Vector start, Vector end)
        {
            Start = start;
            End = end;
        }

        public Vector Start { get; }
        public Vector End { get; }

        public override bool Equals(object obj)
            => obj is Segment rhs && Start.Equals(rhs.Start) && End.Equals(rhs.End);

        public override int GetHashCode()
        {
            int hash = 13;
            hash = (hash * 7) + Start.GetHashCode();
            hash = (hash * 7) + End.GetHashCode();
            return hash;
        }

        public override string ToString()
            => $"Segment: Start = {Start}, End = {End}";
    }
}
