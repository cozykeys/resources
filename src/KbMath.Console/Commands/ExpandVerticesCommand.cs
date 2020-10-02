namespace KbMath.Console.Commands
{
    using Microsoft.Extensions.CommandLineUtils;
    using System;
    using System.Collections.Generic;
    using System.IO;
    using System.Linq;

    using Newtonsoft.Json;
    
    using Models;
    using Services;
    using Core.Geometry2D.Extensions;
    using Core.Geometry2D.Models;
    using Core.Geometry2D.Operations;

    internal class ExpandVerticesCommand
    {
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
        
        public ExpandVerticesCommand(ISvgService svgService)
        {
            _svgService = svgService;
            
            CommandLineApplication command = ApplicationContext.CommandLineApplication
                .Command("expand-vertices", config =>
                {
                    config.Description = "TODO";
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
            if (!double.TryParse(Distance, out double distance))
            {
                Console.WriteLine("Distance must be a valid floating point number.");
                return -2;
            }

            List<Vector> inputVertices = GetVertices(InputPath);
           
            // First, turn the list of vertices into a list of line segments.
            List<Segment> inputSegments = VectorOperations.GetSegments(inputVertices).ToList();
            
            // Next, expand the segments into parallel lines that are a given distance away.
            List<Line> expandedLines = ExpandSegments(inputSegments, distance).ToList();
            
            // Finally, get the intersection points of the adjacent lines.
            List<Vector> expandedVertices = GetIntersectionPoints(expandedLines).ToList();
            
            WriteVerticesToFile(OutputPath, expandedVertices);
            WriteVerticesToConsole(inputVertices, expandedVertices);

            if (!string.IsNullOrEmpty(DebugSvgPath))
            {
                Svg svg = _svgService.CreateSvg();
                _svgService.Append(svg, inputVertices, InitialVertexSvgStyles);
                _svgService.Append(svg, inputSegments, InitialSegmentSvgStyles);
                _svgService.Append(svg, expandedVertices, ExpandedVertexSvgStyles);
                
                List<Segment> expandedSegments = VectorOperations.GetSegments(expandedVertices).ToList();
                _svgService.Append(svg, expandedSegments, ExpandedSegmentSvgStyles);
                
                _svgService.WriteToFile(svg, DebugSvgPath);
            }

            return 0;
        }

        private static void WriteVerticesToConsole(List<Vector> inputVertices, List<Vector> expandedVertices)
        {
            Console.WriteLine("Expansion Details:");

            for (int i = 0; i < inputVertices.Count(); ++i)
            {
                Console.WriteLine($"  {inputVertices[i]} => {expandedVertices[i]}");
            }
        }

        private static IEnumerable<Line> ExpandSegments(IEnumerable<Segment> segments, double distance)
        {
            var lines = new List<Line>();

            foreach (Segment segment in segments)
            {
                var line = segment.ToLine();

                bool InLeftHemisphere(double theta) => ((0.5 * Math.PI) < theta) && (theta < (1.5 * Math.PI));

                lines.Add(InLeftHemisphere(segment.Theta())
                    ? line.Parallel(distance)
                    : line.Parallel(-distance));
            }

            return lines;
        }

        private static IEnumerable<Vector> GetIntersectionPoints(IEnumerable<Line> lines)
        {
            var intersectionPoints = new List<Vector>();

            // Convert this to a list to avoid multiple enumerations.
            var lineList = lines.ToList();
            
            int lineCount = lineList.Count();
            for (int curr = 0; curr < lineCount; ++curr)
            {
                int prev = curr == 0 ? lineCount - 1 : curr - 1;
                intersectionPoints.Add(lineList[curr].Intersection(lineList[prev]));
            }

            return intersectionPoints;
        }

        private static List<Vector> GetVertices(string path)
        {
            string rawInput = File.ReadAllText(path);
            return JsonConvert.DeserializeObject<List<Vector>>(rawInput);
        }
        
        private static void WriteVerticesToFile(string path, IEnumerable<Vector> vertices)
        {
            File.WriteAllText(path, JsonConvert.SerializeObject(vertices, Formatting.Indented));
        }
    }
}
