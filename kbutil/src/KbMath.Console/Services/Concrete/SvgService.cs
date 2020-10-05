using System.Diagnostics;

namespace KbMath.Console.Services.Concrete
{
    using System;
    using System.Collections.Generic;
    using System.IO;
    using System.Linq;
    using System.Text;
    
    using Models;
    using Core.Geometry2D.Models;

    public class SvgService : ISvgService
    {
        public Svg CreateSvg()
        {
            return new Svg();
        }

        public void Append(Svg svg, IEnumerable<Vector> vectors, Dictionary<string, string> styleOverrides)
        {
            var styles = new Dictionary<string, string>
            {
                { "r", styleOverrides.ContainsKey("r") ? styleOverrides["r"] : "0.2" },
                { "stroke", styleOverrides.ContainsKey("stroke") ? styleOverrides["stroke"] : "black" },
                { "stroke-width", styleOverrides.ContainsKey("stroke-width") ? styleOverrides["stroke-width"] : "0.1" },
                { "fill", styleOverrides.ContainsKey("fill") ? styleOverrides["fill"] : "black" }
            };
            
            string style = string.Join(" ", styles.Select(kvp => $"{kvp.Key}=\"{kvp.Value}\""));
            
            foreach (Vector vector in vectors)
            {
                svg.Elements.Add(new SvgElement
                {
                    Content = $"<circle transform=\"translate({vector.X},{vector.Y})\" {style} />"
                });
            }
        }

        public void Append(Svg svg, IEnumerable<Segment> segments, Dictionary<string, string> styleOverrides)
        {
            var styles = new Dictionary<string, string>
            {
                { "stroke", styleOverrides.ContainsKey("stroke") ? styleOverrides["stroke"] : "black" },
                { "stroke-width", styleOverrides.ContainsKey("stroke-width") ? styleOverrides["stroke-width"] : "0.1" },
                { "fill", styleOverrides.ContainsKey("fill") ? styleOverrides["fill"] : "black" }
            };
            
            string style = string.Join(" ", styles.Select(kvp => $"{kvp.Key}=\"{kvp.Value}\""));
            
            foreach (Segment segment in segments)
            {
                string pathData = $"M {segment.Start.X},{segment.Start.Y} L {segment.End.X},{segment.End.Y}";
                svg.Elements.Add(new SvgElement
                {
                    Content = $"<path {style} d=\"{pathData}\" />"
                });
            }
        }
        
        public void Append(Svg svg, IEnumerable<Curve> curves, Dictionary<string, string> styleOverrides)
        {
            var styles = new Dictionary<string, string>
            {
                { "stroke", styleOverrides.ContainsKey("stroke") ? styleOverrides["stroke"] : "black" },
                { "stroke-width", styleOverrides.ContainsKey("stroke-width") ? styleOverrides["stroke-width"] : "0.1" },
                { "fill", styleOverrides.ContainsKey("fill") ? styleOverrides["fill"] : "black" }
            };
            
            string style = string.Join(" ", styles.Select(kvp => $"{kvp.Key}=\"{kvp.Value}\""));
            
            foreach (Curve curve in curves)
            {
                string pathData = $"M {curve.Start.X},{curve.Start.Y} Q {curve.Control.X},{curve.Control.Y}, {curve.End.X},{curve.End.Y}";
                svg.Elements.Add(new SvgElement
                {
                    Content = $"<path {style} d=\"{pathData}\" />"
                });
            }
        }

        public void DrawPath(Svg svg, IEnumerable<Curve> curves, Dictionary<string, string> styleOverrides)
        {
            var styles = new Dictionary<string, string>
            {
                { "stroke", styleOverrides.ContainsKey("stroke") ? styleOverrides["stroke"] : "black" },
                { "stroke-width", styleOverrides.ContainsKey("stroke-width") ? styleOverrides["stroke-width"] : "0.1" },
                { "fill", styleOverrides.ContainsKey("fill") ? styleOverrides["fill"] : "black" }
            };
            
            string style = string.Join(" ", styles.Select(kvp => $"{kvp.Key}=\"{kvp.Value}\""));
            
            var pathDataComponents = new List<string>();

            var curveList = curves.ToList();
            var curveCount = curveList.Count();
            for (int i = 0; i < curveCount; ++i)
            {
                pathDataComponents.Add(i == 0
                    ? $"M {curveList[i].Start.X},{curveList[i].Start.Y}"
                    : $"L {curveList[i].Start.X},{curveList[i].Start.Y}");

                pathDataComponents.Add($"Q {curveList[i].Control.X},{curveList[i].Control.Y} {curveList[i].End.X},{curveList[i].End.Y}");

                if (i == (curveCount - 1))
                {
                    pathDataComponents.Add($"L {curveList[0].Start.X},{curveList[0].Start.Y}");
                }
            }

            string pathData = string.Join(" ", pathDataComponents);
            
            svg.Elements.Add(new SvgElement { Content = $"<path {style} d=\"{pathData}\" />" });
        }

        public void DrawSwitches(Svg svg, IEnumerable<Switch> switches, Dictionary<string, string> styleOverrides)
        {
            var styles = new Dictionary<string, string>
            {
                { "stroke", styleOverrides.ContainsKey("stroke") ? styleOverrides["stroke"] : "black" },
                { "stroke-width", styleOverrides.ContainsKey("stroke-width") ? styleOverrides["stroke-width"] : "0.1" },
                { "fill", styleOverrides.ContainsKey("fill") ? styleOverrides["fill"] : "none" }
            };
            
            string style = string.Join(";", styles.Select(kvp => $"{kvp.Key}:{kvp.Value}"));
            
            foreach (var @switch in switches)
            {
                string id = $"Switch_{@switch.Row}_{@switch.Column}";
                const string pathData = "M -7,-7 H 7 v 1 H 7.8 V 6 H 7 V 7 H -7 V 6 H -7.8 V -6 H -7 V -7 H 7";

                string transform = $"translate({@switch.X},{@switch.Y})";

                if (!@switch.Rotation.Equals(0))
                {
                    transform = $"{transform} rotate({@switch.Rotation})";
                }
                
                var pathAttributes = new Dictionary<string, string>
                {
                    {"id", $"{id}"},
                    {"transform", $"{transform}"},
                    {"d", $"{pathData}"},
                    {"style", $"{style}"}
                };

                string attributes = string.Join(" ", pathAttributes.Select(kvp => $"{kvp.Key}=\"{kvp.Value}\""));

                svg.Elements.Add(new SvgElement
                {
                    Content = $"<path {attributes} />"
                });
            }
        }

        public void WriteToFile(Svg svg, string path)
        {
            var sb = new StringBuilder("<svg>");
            sb.Append(string.Join(Environment.NewLine, svg.Elements.Select(e => e.Content)));
            sb.Append("</svg>");
            
            File.WriteAllText(path, sb.ToString());
        }
    }
}