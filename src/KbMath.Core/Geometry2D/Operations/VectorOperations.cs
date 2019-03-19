namespace KbMath.Core.Geometry2D.Operations
{
    using System;
    using System.Collections.Generic;
    using System.Linq;
    
    using Extensions;
    using Models;

    public static class VectorOperations
    {
        public static Vector GetCenter(IEnumerable<Vector> vectors)
        {
            double minX = double.MaxValue;
            double minY = double.MaxValue;
            double maxX = double.MinValue;
            double maxY = double.MinValue;

            foreach (var vector in vectors)
            {
                if (vector.X > maxX) maxX = vector.X;
                if (vector.X < minX) minX = vector.X;
                if (vector.Y > maxY) maxY = vector.Y;
                if (vector.Y < minY) minY = vector.Y;
            }

            double centerX = minX + ((maxX - minX) / 2);
            double centerY = minY + ((maxY - minY) / 2);
            
            return new Vector(centerX, centerY);
        }

        public static IEnumerable<Segment> GetSegments(IEnumerable<Vector> vertices)
        {
            var segments = new List<Segment>();

            var vertexList = vertices.ToList();
            int vertexCount = vertexList.Count();

            if (vertexCount < 2)
            {
                throw new InvalidOperationException();
            }
            
            for (int curr = 0; curr < vertexCount; ++curr)
            {
                int next = (curr + 1) % vertexCount;
                segments.Add(new Segment(vertexList[curr], vertexList[next]));
            }

            return segments;
        }
        
        public static IEnumerable<Curve> GenerateCurves(List<Vector> vertices, double distance)
        {
            List<Curve> curves = new List<Curve>();

            // This will be used to get the indices of the previous and next vertices in a circular fashion.
            int GetIndex(int current, int difference, int count)
            {
                int index = (current + difference) % count;
                return index < 0 ? count + index : index;
            }

            int vertexCount = vertices.Count();
            for (int curr = 0; curr < vertexCount; curr++)
            {
                int prev = GetIndex(curr, -1, vertexCount);
                int next = GetIndex(curr, 1, vertexCount);

                var s1 = new Segment(vertices[curr], vertices[prev]);
                var s2 = new Segment(vertices[curr], vertices[next]);

                Vector start = s1.GetPoint(distance);
                Vector end = s2.GetPoint(distance);
                
                curves.Add(new Curve
                (
                    start,
                    end,
                    vertices[curr]
                ));
            }

            return curves;
        }
    }
}