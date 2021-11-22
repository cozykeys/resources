using System.Drawing;
using KbUtil.Lib.Common.Extensions;

namespace KbUtil.Lib.Geometry2D.Extensions
{
    using System;
    using Models;

    public static class LineExtensions
    {
        /// <summary>
        /// Get the intersection point of two lines.
        /// </summary>
        /// <param name="lhs">The first line.</param>
        /// <param name="rhs">The second line.</param>
        /// <returns>The point at which the lines intersect.</returns>
        public static Vector Intersection(this Line lhs, Line rhs)
        {
            // If the slopes are equal, there is either no intersection point or infinite intersection points.
            if (lhs.Slope.Equals(rhs.Slope))
            {
                return null;
            }
            
            double x = (rhs.YIntercept - lhs.YIntercept) / (lhs.Slope - rhs.Slope);
            double y = lhs.Slope * x + lhs.YIntercept;

            return new Vector(x, y);
        }

        /// <summary>
        /// Get a line parallel to the given line that is a specified distance away.
        /// </summary>
        /// <param name="line">The original line.</param>
        /// <param name="distance">The distance from the original line.</param>
        /// <returns>The line parallel to the original line.</returns>
        public static Line Parallel(this Line line, double distance)
        {
            if (distance.Equals(0))
            {
                return line;
            }
            
            if (line.Slope.Equals(0))
            {
                return new Line(line.Slope, line.YIntercept + distance);
            }

            if (double.IsNaN(line.Slope))
            {
                throw new InvalidOperationException();
            }
            
            // Calculate theta for direction perpendicular to the line. Note that lines have no concept of direction so
            // there are two valid perpendicular directions; we are simply choosing one of them.
            double perpendicularTheta = (Math.Atan(line.Slope) + (Math.PI / 2));

            // Get the point that is the given distance away in the perpendicular direction.
            double x = distance * Math.Cos(perpendicularTheta);
            double y = distance * Math.Sin(perpendicularTheta);

            var pointOnParallelLine = new Vector(x, line.YIntercept + y);
            
            // Using the point, calculate the Y Intercept of the new line. (y = mx + b => b = y - mx)
            var yIntercept = pointOnParallelLine.Y - line.Slope * pointOnParallelLine.X;
            
            return new Line(line.Slope, yIntercept);
        }

        /// <summary>
        /// Get the value of X for the point where the line crosses the X Axis.
        /// </summary>
        /// <param name="line">The line.</param>
        /// <returns>The value of X for the point where the line crosses the X Axis.</returns>
        public static double XIntercept(this Line line)
            => -line.YIntercept / line.Slope;
    }
}
