using System.Diagnostics;

namespace KbMath.Console.Services
{
    using System.Collections.Generic;
    using Core.Geometry2D.Models;
    using Models;

    public interface ISvgService
    {
        Svg CreateSvg();
        void Append(Svg svg, IEnumerable<Vector> vectors, Dictionary<string, string> styleOverrides);
        void Append(Svg svg, IEnumerable<Segment> segments, Dictionary<string, string> styleOverrides);
        void Append(Svg svg, IEnumerable<Curve> curves, Dictionary<string, string> styleOverrides);
        void DrawPath(Svg svg, IEnumerable<Curve> curves, Dictionary<string, string> styleOverrides);
        void DrawSwitches(Svg svg, IEnumerable<Switch> switches, Dictionary<string, string> styleOverrides);
        void WriteToFile(Svg svg, string path);
    }
}