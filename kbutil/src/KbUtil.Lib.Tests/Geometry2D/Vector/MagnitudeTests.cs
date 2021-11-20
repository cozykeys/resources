namespace KbUtil.Lib.Tests.Geometry2D.Vector
{
    using System.Collections.Generic;
    using Core.Geometry2D.Extensions;
    using Core.Geometry2D.Models;
    using Xunit;

    public class MagnitudeTests
    {
        public static IEnumerable<object[]> TestData = new List<object[]>
        {
            new object[] { new Vector(1.0, 0.0), 1.0 },
            new object[] { new Vector(0.0, 1.0), 1.0 },
            new object[] { new Vector(1.0, 1.0), 1.4142135623730951 }
        };

        [Theory]
        [MemberData(nameof(TestData))]
        public void TestMagnitude(Vector vector, double expectedMagnitude)
            => Assert.Equal(expectedMagnitude, vector.Magnitude());
    }
}
