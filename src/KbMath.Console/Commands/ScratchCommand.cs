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

            var vectors = new List<Vector>
            {
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
