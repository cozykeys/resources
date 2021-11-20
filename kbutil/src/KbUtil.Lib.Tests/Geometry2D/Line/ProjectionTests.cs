namespace KbUtil.Lib.Tests.Geometry2D.Line
{
    using System;
    using System.Collections.Generic;
    using KbUtil.Lib.Geometry2D.Extensions;
    using KbUtil.Lib.Geometry2D.Models;
    using KbUtil.Lib.Tests.Comparers;
    using Xunit;
    
    public class ProjectionTests
    {
        private static readonly VectorComparer VectorComparer = new VectorComparer();
        
        public static IEnumerable<object[]> TestData = new List<object[]>
        {
            new object[] { new Vector(0.0, 0.0), 0.0, 1.0, new Vector(1.0, 0.0) },
            new object[] { new Vector(0.0, 0.0), Math.PI, 1.0, new Vector(-1.0, 0.0), },
            new object[] { new Vector(0.0, 0.0), Math.PI/4.0, Math.Sqrt(2), new Vector(1.0, 1.0) }
        };

        [Theory]
        [MemberData(nameof(TestData))]
        public void TestProjection(Vector vector, double theta, double magnitude, Vector expectedProjection)
            => Assert.Equal(expectedProjection, vector.Project(theta, magnitude), VectorComparer);
    }
}
