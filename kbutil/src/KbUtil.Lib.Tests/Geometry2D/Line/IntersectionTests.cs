namespace KbUtil.Lib.Tests.Geometry2D.Line
{
    using System.Collections.Generic;
    using Core.Geometry2D.Extensions;
    using Core.Geometry2D.Models;
    using Xunit;
    
    public class IntersectionTests
    {
        public static IEnumerable<object[]> TestData = new List<object[]>
        {
            new object[] { new Line(1.0, 0.0), new Line(-1.0, 0.0), new Vector(0.0, 0.0) },
            new object[] { new Line(0.5, 2.0), new Line(2.0, 1.0), new Vector(2.0 / 3.0, 7.0 / 3.0) },
            new object[] { new Line(1.0, 0.0), new Line(1.0, 0.0), null }
        };

        [Theory]
        [MemberData(nameof(TestData))]
        public void TestIntersection(Line lhs, Line rhs, Vector expectedIntersection)
            => Assert.Equal(expectedIntersection, lhs.Intersection(rhs));
    }
}
