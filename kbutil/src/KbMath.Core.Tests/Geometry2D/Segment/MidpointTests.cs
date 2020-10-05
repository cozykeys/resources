namespace KbMath.Core.Tests.Geometry2D.Segment
{
    using System.Collections.Generic;
    using Core.Geometry2D.Extensions;
    using Core.Geometry2D.Models;
    using Xunit;

    public class MidpointTests
    {
        public static IEnumerable<object[]> TestData = new List<object[]>
        {
            new object[] { new Segment(new Vector(-1.0, 0.0), new Vector(1.0, 0.0)), new Vector(0.0, 0.0) },
            new object[] { new Segment(new Vector(0.0, -1.0), new Vector(0.0, 1.0)), new Vector(0.0, 0.0) }
        };

        [Theory]
        [MemberData(nameof(TestData))]
        public void TestMidpoint(Segment segment, Vector expectedMidpoint)
            => Assert.Equal(expectedMidpoint, segment.Midpoint());
    }
}
