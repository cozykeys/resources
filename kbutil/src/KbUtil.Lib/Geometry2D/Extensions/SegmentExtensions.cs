using System;
using KbUtil.Lib.Common.Extensions;

namespace KbUtil.Lib.Geometry2D.Extensions
{
    using KbUtil.Lib.Geometry2D.Models;

    public static class SegmentExtensions
    {
        public static double Dx(this Segment segment)
            => segment.End.X - segment.Start.X;
        
        public static double Dy(this Segment segment)
            => segment.End.Y - segment.Start.Y;
        
        public static Vector Midpoint(this Segment segment)
            => new Vector(
                segment.Start.X + (segment.Dx() / 2),
                segment.Start.Y + (segment.Dy() / 2));

        public static double Theta(this Segment segment)
        {
            double dx = segment.Dx();
            double dy = segment.Dy();

            // Early out if the segment is parallel to an axis.
            if (dy.Equals(0) && dx > 0) return (0.0 * Math.PI);
            if (dx.Equals(0) && dy > 0) return (0.5 * Math.PI);
            if (dy.Equals(0) && dx < 0) return (1.0 * Math.PI);
            if (dx.Equals(0) && dy < 0) return (1.5 * Math.PI);

            double theta = Math.Atan(dy / dx);

            // Because slope doesn't take direction into account, we have to manually adjust theta for segments whose
            // direction points into quadrants 2 and 3.
            if ((dx < 0 && dy > 0) || (dx < 0 && dy < 0))
                theta += Math.PI;

            return theta.Mod(Math.PI * 2);
        }

        public static double Slope(this Segment segment)
        {
            var dx = segment.Dx();
            if (dx.Equals(0))
            {
                return double.NaN;
            }
            
            return segment.Dy() / dx;
        }

        public static Line ToLine(this Segment segment)
        {
            var slope = segment.Slope();

            var yIntercept = double.IsNaN(slope)
                ? double.NaN
                : segment.Start.Y - (slope * segment.Start.X);
            
            return new Line(slope, yIntercept);
        }

        public static double Magnitude(this Segment segment)
            => new Vector(segment.Dx(), segment.Dy()).Magnitude();

        public static Vector GetPoint(this Segment segment, double distance)
        {
            if (distance > segment.Magnitude())
            {
                throw new InvalidOperationException();
            }

            double theta = segment.Theta();

            double dx = distance * Math.Cos(theta);
            double dy = distance * Math.Sin(theta);
            
            return new Vector(segment.Start.X + dx, segment.Start.Y + dy);
        }
    }
}
