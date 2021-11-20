﻿namespace KbMath.Core.Tests.Geometry2D.Line
{
    using System;
    using System.Collections.Generic;
    using Core.Geometry2D.Extensions;
    using Core.Geometry2D.Models;
    using Xunit;
    using KbMath.Core.Tests.Comparers;
    
    public class ParallelTests
    {
        private static readonly LineComparer LineComparer = new LineComparer();
        
        public static IEnumerable<object[]> TestData = new List<object[]>
        {
            new object[] { new Line(1.0, 0.0), 1.0, new Line(1.0, Math.Sqrt(2)) },
            new object[] { new Line(-1.0, 0.0), 1.0, new Line(-1.0, Math.Sqrt(2)) },
            new object[] { new Line(0.0, 0.0), 1.0, new Line(0.0, 1.0) }
        };

        [Theory]
        [MemberData(nameof(TestData))]
        public void TestParallel(Line line, double distance, Line expectedParallelLine)
            => Assert.Equal(expectedParallelLine, line.Parallel(distance), LineComparer);
    }
}
