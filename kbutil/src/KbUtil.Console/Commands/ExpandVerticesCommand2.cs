namespace KbUtil.Console.Commands
{
    using Microsoft.Extensions.CommandLineUtils;
    using System;
    using System.Collections.Generic;
    using System.IO;
    using System.Linq;
    using Microsoft.Extensions.Logging;

    using Newtonsoft.Json;
    
    using Models;
    using Services;
    using KbUtil.Lib.Geometry2D.Extensions;
    using KbUtil.Lib.Geometry2D.Models;
    using KbUtil.Lib.Geometry2D.Operations;

    using GeometRi;

    internal class ExpandVerticesCommand2
    {
        private readonly ILogger _logger;
        private readonly ISvgService _svgService;
        
        private readonly CommandArgument _inputPathArgument;
        private readonly CommandArgument _outputPathArgument;
        private readonly CommandArgument _distanceArgument;
        
        private readonly CommandOption _debugSvg;
        
        private static readonly Dictionary<string, string> InitialSegmentSvgStyles = new Dictionary<string, string>
        {
            { "stroke", "blue" },
            { "stroke-width", "0.2" }
        };
        
        private static readonly Dictionary<string, string> InitialVertexSvgStyles = new Dictionary<string, string>
        {
            { "r", "0.5" },
            { "stroke", "black" },
            { "stroke-width", "0.1" },
            { "fill", "blue" }
        };
        
        private static readonly Dictionary<string, string> ExpandedSegmentSvgStyles = new Dictionary<string, string>
        {
            { "stroke", "purple" },
            { "stroke-width", "0.2" }
        };
        
        private static readonly Dictionary<string, string> ExpandedVertexSvgStyles = new Dictionary<string, string>
        {
            { "r", "0.5" },
            { "stroke", "black" },
            { "stroke-width", "0.1" },
            { "fill", "purple" }
        };
        
        public ExpandVerticesCommand2(ILoggerFactory loggerFactory, ISvgService svgService)
        {
            _logger = loggerFactory.CreateLogger(nameof(ExpandVerticesCommand2));

            _svgService = svgService;
            
            CommandLineApplication command = ApplicationContext.CommandLineApplication
                .Command("expand-vertices2", config =>
                {
                    config.Description = "Expand list of vertices; this version uses the GeometRi library to avoid issues with vertical lines that the expand-vertices command has.";
                    config.OnExecute(() => Execute());
                });
            
            _inputPathArgument = command.Argument("<input-path>", "TODO");
            _outputPathArgument = command.Argument("<output-path>", "TODO");
            _distanceArgument = command.Argument("<distance>", "TODO");
            
            _debugSvg = command.Option(
                "--debug-svg",
                "TODO",
                CommandOptionType.SingleValue);
        }

        private string InputPath => _inputPathArgument.Value;
        private string OutputPath => _outputPathArgument.Value;
        private string Distance => _distanceArgument.Value;

        private string DebugSvgPath => _debugSvg?.Value();
        
        public int Execute()
        {
            _logger.LogInformation("ExpandVerticesCommand2.Execute");

            if (!double.TryParse(Distance, out double distance))
            {
                Console.WriteLine("Distance must be a valid floating point number.");
                return -2;
            }

            // Parse JSON input file
            List<Point3d> inputVertices = GetVertices(InputPath)
                .Select(v => new Point3d(v.X, v.Y, 0.0))
                .ToList();

            _logger.LogInformation("Parsed input vertices:");
            foreach (Point3d v in inputVertices) {
                _logger.LogInformation($" - ({v.X}, {v.Y})");
            }
           
            // First, turn the list of vertices into a list of line segments.
            List<Segment3d> inputSegments = GetSegments(inputVertices).ToList();

            _logger.LogInformation("Converted vertices to segments:");
            foreach (Segment3d s in inputSegments) {
                _logger.LogInformation($" - [({s.P1.X}, {s.P1.Y}), ({s.P2.X}, {s.P2.Y})]");
            }
            
            // Next, expand the segments into parallel lines that are a given distance away.
            List<Line3d> expandedLines = ExpandSegments(inputSegments, distance)
                .ToList();

            _logger.LogInformation("Expanded segments into lines:");
            foreach (Line3d l in expandedLines) {
                _logger.LogInformation($" - Direction: {l.Direction}, Point: {l.Point}");
            }
            
            // Finally, get the intersection points of the adjacent lines.
            List<Point3d> expandedVertices = GetIntersectionPoints(expandedLines)
                .ToList();
            
            WriteVerticesToFile(OutputPath, expandedVertices);
            WriteVerticesToConsole(inputVertices, expandedVertices);

            if (!string.IsNullOrEmpty(DebugSvgPath))
            {
                Svg svg = _svgService.CreateSvg();
                _svgService.Append(
                    svg,
                    inputVertices.Select(p => new Vector(p.X, p.Y)),
                    InitialVertexSvgStyles);
                _svgService.Append(
                    svg,
                    inputSegments.Select(s => new Segment(
                        new Vector(s.P1.X, s.P1.Y),
                        new Vector(s.P2.X, s.P2.Y)
                    )),
                    InitialSegmentSvgStyles);
                _svgService.Append(
                    svg,
                    expandedVertices.Select(p => new Vector(p.X, p.Y)),
                    ExpandedVertexSvgStyles);
                
                _svgService.Append(
                    svg,
                    GetSegments(expandedVertices).Select(s => new Segment(
                        new Vector(s.P1.X, s.P1.Y),
                        new Vector(s.P2.X, s.P2.Y)
                    ))
                    .ToList(),
                    ExpandedSegmentSvgStyles);
                
                _svgService.WriteToFile(svg, DebugSvgPath);
            }

            return 0;
        }

        private static void WriteVerticesToConsole(List<Point3d> inputVertices, List<Point3d> expandedVertices)
        {
            Console.WriteLine("Expansion Details:");

            for (int i = 0; i < inputVertices.Count(); ++i)
            {
                var ix = inputVertices[i].X;
                var iy = inputVertices[i].Y;
                var ox = expandedVertices[i].X;
                var oy = expandedVertices[i].Y;
                Console.WriteLine($"  ({ix},{iy}) => ({ox}, {oy})");
            }
        }

        private static IEnumerable<Line3d> ExpandSegments(IEnumerable<Segment3d> segments, double distance)
        {
            var lines = new List<Line3d>();

            Segment3d PerpendicularClockwise(Segment3d s)
            {
                var dx = s.P2.X - s.P1.X;
                var dy = s.P2.Y - s.P1.Y;
                var v = new Vector3d(dy, -dx, 0.0).Normalized * distance;
                return new Segment3d(s.P1.Translate(v), s.P2.Translate(v));
            }

            /*
            Segment3d PerpendicularCounterClockwise(Segment3d s)
            {
                var dx = s.P2.X - s.P1.X;
                var dy = s.P2.Y - s.P1.Y;
                var v = new Vector3d(-dy, dx, 0.0).Normalized * distance;
                return new Segment3d(s.P1.Translate(v), s.P2.Translate(v));
            }
            */

            foreach (Segment3d segment in segments)
            {
                lines.Add(PerpendicularClockwise(segment).ToLine);
            }

            return lines;
        }

        private static IEnumerable<Point3d> GetIntersectionPoints(IEnumerable<Line3d> lines)
        {
            var intersectionPoints = new List<Point3d>();

            var lineList = lines.ToList();
            int lineCount = lineList.Count();
            for (int curr = 0; curr < lineCount; ++curr)
            {
                int prev = curr == 0 ? lineCount - 1 : curr - 1;
                var l1 = lineList[curr];
                var l2 = lineList[prev];
                var i = l1.IntersectionWith(l2);
                if (i == null)
                    throw new Exception($"ERROR: {curr}");
                var p = i as Point3d;
                if (p == null)
                    throw new Exception("ERROR");
                intersectionPoints.Add(p);
            }

            return intersectionPoints;
        }

        private static List<Vector> GetVertices(string path)
        {
            string rawInput = File.ReadAllText(path);
            return JsonConvert.DeserializeObject<List<Vector>>(rawInput);
        }
        
        private static void WriteVerticesToFile(string path, IEnumerable<Point3d> points)
        {
            var vertices = points.Select(p => new Vector(p.X, p.Y)).ToList();
            File.WriteAllText(path, JsonConvert.SerializeObject(vertices, Formatting.Indented));
        }

        public static IEnumerable<Segment3d> GetSegments(IEnumerable<Point3d> points)
        {
            var segments = new List<Segment3d>();

            var vertexList = points.ToList();
            int vertexCount = vertexList.Count();

            if (vertexCount < 2)
            {
                throw new InvalidOperationException();
            }
            
            for (int curr = 0; curr < vertexCount; ++curr)
            {
                int next = (curr + 1) % vertexCount;
                segments.Add(new Segment3d(vertexList[curr], vertexList[next]));
            }

            return segments;
        }
    }
}
