namespace KbMath.Core.Geometry2D.Models
{
    using System;

    public class Curve
    {
        public Curve(Vector start, Vector end, Vector control)
        {
            Start = start;
            End = end;
            Control = control;
        }
        
        public Vector Start { get; }
        public Vector End { get; }
        public Vector Control { get; }

        public override bool Equals(object obj)
            => obj is Curve rhs
               && Start == rhs.Start
               && End == rhs.End
               && Control == rhs.Control;

        public override int GetHashCode()
        {
            int hash = 13;
            hash = (hash * 7) + Start.GetHashCode();
            hash = (hash * 7) + End.GetHashCode();
            hash = (hash * 7) + Control.GetHashCode();
            return hash;
        }

        public Curve Round(int digits)
            => new Curve(Start.Round(digits), End.Round(digits), Control.Round(digits));

        public override string ToString()
        {
            return $"Curve: Start = {Start}, End = {End}, Control = {Control}";
        }
    }
}
