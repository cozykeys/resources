namespace KbMath.Console.Commands
{
    using System;
    using System.Collections.Generic;
    using KbMath.Core.Geometry2D.Extensions;
    using KbMath.Core.Geometry2D.Models;
    using Microsoft.Extensions.CommandLineUtils;
    using Services;

    internal class ScratchCommand
    {
        public ScratchCommand(ISvgService svgService)
        {
            CommandLineApplication command = ApplicationContext.CommandLineApplication
                .Command("scratch", config =>
                {
                    config.Description = "Command for general purpose scratchpad use";
                    config.OnExecute(() => Execute());
                });
        }

        public void Temp2(double theta, double m)
        { 
            double dt = 0.17453292519943295 + 4.71238898038469;

            double t0 = dt - theta;
            double t1 = dt + theta;

            var vectorLists = new List<List<Vector>> {
                new List<Vector> // Top Right
                {
                    new Vector(-5,0),                // Placeholder
                    Project(205.126, 73.272, t0, m), // r01c07
                    Project(205.126, 73.272, t1, m), // r01c07
                    Project(223.018,  65.04, t0, m), // r01c08
                    Project(223.018,  65.04, t1, m), // r01c08
                    Project(240.737, 55.823, t0, m), // r01c09
                    Project(240.737, 55.823, t1, m), // r01c09
                    Project(260.366, 57.439, t0, m), // r01c10
                    Project(260.366, 57.439, t1, m), // r01c10
                    Project(279.647, 57.085, t0, m), // r01c11
                    Project(279.647, 57.085, t1, m), // r01c11
                    Project(298.408, 53.777, t0, m), // r01c12
                    Project(298.408, 53.777, t1, m), // r01c12
                },
                new List<Vector> // Bottom Right
                {
                    new Vector(-5,0),                // Placeholder
                    Project(208.434, 92.032, t0, m), // r02c07
                    Project(208.434, 92.032, t1, m), // r02c07
                    Project(226.326, 83.8,   t0, m), // r02c08
                    Project(226.326, 83.8,   t1, m), // r02c08
                    Project(244.045, 74.584, t0, m), // r02c09
                    Project(244.045, 74.584, t1, m), // r02c09
                    Project(263.674, 76.2,   t0, m), // r02c10
                    Project(263.674, 76.2,   t1, m), // r02c10
                    Project(282.955, 75.846, t0, m), // r02c11
                    Project(282.955, 75.846, t1, m), // r02c11
                    Project(301.716, 72.538, t0, m), // r02c12
                    Project(301.716, 72.538, t1, m), // r02c12
                },
                new List<Vector> // Top Left
                {
                    new Vector(5,0),                           // Placeholder
                    Project(144.874, 73.272, Math.PI - t0, m), // r01c00
                    Project(144.874, 73.272, Math.PI - t1, m), // r01c00
                    Project(126.982,  65.04, Math.PI - t0, m), // r01c01
                    Project(126.982,  65.04, Math.PI - t1, m), // r01c01
                    Project(109.263, 55.823, Math.PI - t0, m), // r01c02
                    Project(109.263, 55.823, Math.PI - t1, m), // r01c02
                    Project(89.634,  57.439, Math.PI - t0, m), // r01c03
                    Project(89.634,  57.439, Math.PI - t1, m), // r01c03
                    Project(70.353,  57.085, Math.PI - t0, m), // r01c04
                    Project(70.353,  57.085, Math.PI - t1, m), // r01c04
                    Project(51.592,  53.777, Math.PI - t0, m), // r01c05
                    Project(51.592,  53.777, Math.PI - t1, m), // r01c05
                },
                new List<Vector> // Bottom Left
                {
                    new Vector(5,0),                           // Placeholder
                    Project(141.566, 92.032, Math.PI - t0, m), // r02c05
                    Project(141.566, 92.032, Math.PI - t1, m), // r02c05
                    Project(123.674, 83.8,   Math.PI - t0, m), // r02c04
                    Project(123.674, 83.8,   Math.PI - t1, m), // r02c04
                    Project(105.955, 74.584, Math.PI - t0, m), // r02c03
                    Project(105.955, 74.584, Math.PI - t1, m), // r02c03
                    Project(86.326,  76.2,   Math.PI - t0, m), // r02c02
                    Project(86.326,  76.2,   Math.PI - t1, m), // r02c02
                    Project(67.045,  75.846, Math.PI - t0, m), // r02c01
                    Project(67.045,  75.846, Math.PI - t1, m), // r02c01
                    Project(48.284,  72.538, Math.PI - t0, m), // r02c00
                    Project(48.284,  72.538, Math.PI - t1, m), // r02c00
                },
            };

            foreach (var vectors in vectorLists)
            {
                vectors[0] = new Vector(vectors[1].X + vectors[0].X, vectors[1].Y + vectors[0].Y);
                for (int i = 0; i < vectors.Count - 1; ++i)
                {
                    Vector v0 = vectors[i];
                    Vector v1 = vectors[i+1];

                    string start = $"(start {v0.X} {v0.Y})";
                    string end = $"(end {v1.X} {v1.Y})";
                    Console.WriteLine($"(segment {start} {end} (width 0.2032) (layer Back) (net 5))");
                }
            }
        }

        public void Temp()
        {
            // Triangle 1
            {
                var dx = 7;
                var dy = 8.89;
                var theta = Math.Atan(dx/dy);
                var m = dx / Math.Sin(theta);
                Temp2(theta, m);
            }
        
            // Triangle 1
            {
                var dx = 7;
                var dy = 9.525;
                var theta = Math.Atan(dx/dy);
                var m = dx / Math.Sin(theta);
                Temp2(theta, m);
            }

            // Triangle 1
            {
                var dx = 7;
                var dy = 10.16;
                var theta = Math.Atan(dx/dy);
                var m = dx / Math.Sin(theta);
                Temp2(theta, m);
            }
        }

        private void Foo()
        {
            double diameter = 18.1;
            double radius = diameter / 2;

            //double m1 = Math.Sqrt(2 * (radius * radius));
            double m1 = 19.05 / 2.0;
            double m2 = m1 - 0.635;
            double m3 = m1 + 0.635;
            
            var magnitudes = new List<double> { m1, m2, m3 };

            //double topRightTheta = (45 - 10).ToRadians();
            //double topLeftTheta = (90 + 45 - 10).ToRadians();
            //double botLeftTheta = (180 + 45 - 10).ToRadians();
            //double botRightTheta = (270 + 45 - 10).ToRadians();
            //double topTheta = (90 - 10).ToRadians();
            //double botTheta = (270 - 10).ToRadians();

            double topRightTheta = (45 + 10).ToRadians();
            double topLeftTheta = (90 + 45 + 10).ToRadians();
            double botLeftTheta = (180 + 45 + 10).ToRadians();
            double botRightTheta = (270 + 45 + 10).ToRadians();
            double topTheta = (90 + 10).ToRadians();
            double botTheta = (270 + 10).ToRadians();

            
            foreach (var m in magnitudes)
            {
                var vectors = new List<Vector>
                {
                    Project(205.126, 73.272, botTheta, m), // r01c07
                    Project(223.018,  65.04, botTheta, m), // r01c08
                    Project(240.737, 55.823, botTheta, m), // r01c09
                    Project(260.366, 57.439, botTheta, m), // r01c10
                    Project(279.647, 57.085, botTheta, m), // r01c11
                    Project(298.408, 53.777, botTheta, m), // r01c12
                };

                for (int i = 0; i < vectors.Count - 1; ++i)
                {
                    Vector v0 = vectors[i];
                    Vector v1 = vectors[i+1];

                    string start = $"(start {v0.X} {v0.Y})";
                    string end = $"(end {v1.X} {v1.Y})";
                    Console.WriteLine($"(segment {start} {end} (width 0.2032) (layer Back) (net 5))");
                }
            }
        }

        public int Execute()
        {
            Temp();
            //Foo();

            return 0;
        }

        private Vector Project(double x, double y, double theta, double magnitude)
        { 
            var v = new Vector(x, -y).Project(theta, magnitude);
            return new Vector(v.X, -v.Y);
        }
    }
}
