namespace KbUtil.Lib.Geometry2D.Operations
{
    using System.Collections.Generic;
    using System.Linq;
    using Extensions;
    using Models;

    public static class SegmentOperations
    {
        public static IEnumerable<Line> GetLines(IEnumerable<Segment> segments)
            => segments.Select(segment => segment.ToLine());
    }
}