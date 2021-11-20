namespace KbUtil.Lib.Tests.Geometry2D.Line
{
    using System.Collections.Generic;
    using Core.Geometry2D.Extensions;
    using Core.Geometry2D.Models;
    using Xunit;
    
    public class XInterceptTests
    {
        public static IEnumerable<object[]> TestData = new List<object[]>
        {
            new object[] { new Line(1.0, 0.0), 0.0 },
            new object[] { new Line(1.0, -1.0), 1.0 },
            new object[] { new Line(-1.0, 1.0), 1.0 }
        };

        [Theory]
        [MemberData(nameof(TestData))]
        public void TestXIntercept(Line line, double expectedXIntercept)
            => Assert.Equal(expectedXIntercept, line.XIntercept());
    }
}
