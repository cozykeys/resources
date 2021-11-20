namespace KbUtil.Lib.Tests.Geometry2D.Segment
{
    using System.Collections.Generic;
    using KbUtil.Lib.Geometry2D.Extensions;
    using KbUtil.Lib.Geometry2D.Models;
    using Xunit;

    public class ThetaTests
    {
        private const int DecimalPrecision = 4;
        
        public static IEnumerable<object[]> TestData = new List<object[]>
        {
            new object[] { new Segment(new Vector(0, 0), new Vector(1, 0)), 0.0 },
            new object[] { new Segment(new Vector(0, 0), new Vector(2, 1)), 0.4636476090008061 },
            new object[] { new Segment(new Vector(0, 0), new Vector(1, 2)), 1.1071487177940904 },
            new object[] { new Segment(new Vector(0, 0), new Vector(0, 1)), 1.5707963267948966 },
            new object[] { new Segment(new Vector(0, 0), new Vector(-1, 2)), 2.034443936 },
            new object[] { new Segment(new Vector(0, 0), new Vector(-2, 1)), 2.677945045 },
            new object[] { new Segment(new Vector(0, 0), new Vector(-1, 0)), 3.141592653589793 },
            new object[] { new Segment(new Vector(0, 0), new Vector(-2, -1)), 3.605240263 },
            new object[] { new Segment(new Vector(0, 0), new Vector(-1, -2)), 4.248741371 },
            new object[] { new Segment(new Vector(0, 0), new Vector(0, -1)), 4.71238898038469 },
            new object[] { new Segment(new Vector(0, 0), new Vector(1, -2)), 5.176036589 },
            new object[] { new Segment(new Vector(0, 0), new Vector(2, -1)), 5.819537698 }
        };

        [Theory]
        [MemberData(nameof(TestData))]
        public void TestTheta(Segment segment, double expectedTheta)
            => Assert.Equal(expectedTheta, segment.Theta(), DecimalPrecision);
    }
}
