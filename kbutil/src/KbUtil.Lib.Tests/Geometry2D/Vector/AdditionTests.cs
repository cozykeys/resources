namespace KbUtil.Lib.Tests.Geometry2D.Vector
{
    using System.Collections.Generic;
    using Core.Geometry2D.Extensions;
    using Core.Geometry2D.Models;
    using Xunit;
    
    public class AdditionTests
    {
        public static IEnumerable<object[]> TestData = new List<object[]>
        {
            new object[] { new Vector(1.0, 0.0), new Vector(0.0, 1.0), new Vector(1.0, 1.0) }
        };

        [Theory]
        [MemberData(nameof(TestData))]
        public void TestAddition(Vector lhs, Vector rhs, Vector expectedSum)
            => Assert.Equal(expectedSum, lhs.Add(rhs));
    }
}
