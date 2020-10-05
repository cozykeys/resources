using System.Linq;
using KbMath.Console.Models;
using KbMath.Console.Services;

namespace KbMath.Console.Commands
{
    using Microsoft.Extensions.CommandLineUtils;
    using System;
    using System.Collections.Generic;
    using System.IO;
    
    using Newtonsoft.Json;
    
    using Core.Geometry2D.Models;
    using Core.Geometry2D.Operations;

    internal class GenerateCurvesCommand
    {
        private readonly ISvgService _svgService;
        
        private readonly CommandArgument _inputPathArgument;
        private readonly CommandArgument _outputPathArgument;
        private readonly CommandArgument _distanceArgument;
        
        private readonly CommandOption _debugSvg;
        private readonly CommandOption _outputXml;
        
        private static readonly Dictionary<string, string> CurveSvgStyles = new Dictionary<string, string>
        {
            { "stroke", "black" },
            { "stroke-width", "0.1" },
            { "fill", "none" }
        };
        
        private static readonly Dictionary<string, string> SegmentSvgStyles = new Dictionary<string, string>
        {
            { "stroke", "purple" },
            { "stroke-width", "0.2" }
        };
        
        private static readonly Dictionary<string, string> VectorSvgStyles = new Dictionary<string, string>
        {
            { "r", "0.5" },
            { "stroke", "black" },
            { "stroke-width", "0.1" },
            { "fill", "purple" }
        };
        
        public GenerateCurvesCommand(ISvgService svgService)
        {
            _svgService = svgService;
            
            CommandLineApplication command = ApplicationContext.CommandLineApplication
                .Command("generate-curves", config =>
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

            _outputXml = command.Option(
                "--output-xml",
                "TODO",
                CommandOptionType.SingleValue);
        }

        private string InputPath => _inputPathArgument.Value;
        private string OutputPath => _outputPathArgument.Value;
        private string Distance => _distanceArgument.Value;
        
        private string DebugSvgPath => _debugSvg?.Value();
        private string OutputXmlPath => _outputXml?.Value();

        public int Execute()
        {
            if (!float.TryParse(Distance, out float distance))
            {
                Console.WriteLine("Distance must be a valid floating point number.");
                return -1;
            }

            var vertices = GetVertices(InputPath);

            List<Curve> curves = VectorOperations
                .GenerateCurves(vertices, distance)
                .Select(c => c.Round(3))
                .ToList();

            WriteCurvesToFile(OutputPath, curves);
            
            if (!string.IsNullOrEmpty(DebugSvgPath))
            {
                Svg svg = _svgService.CreateSvg();
                
                List<Vector> curvePoints = FlattenCurvePoints(curves).ToList();
                _svgService.Append(svg, curves, CurveSvgStyles);
                _svgService.Append(svg, curvePoints, VectorSvgStyles);
                _svgService.Append(svg, VectorOperations.GetSegments(curvePoints), SegmentSvgStyles);
                
                _svgService.WriteToFile(svg, DebugSvgPath);
            }

            if (!string.IsNullOrEmpty(OutputXmlPath))
            {
                List<string> xmlLines = new List<string>();

                for (int i = 0; i < curves.Count; ++i)
                {
                    double startX = curves[i].Start.X;
                    double startY = curves[i].Start.Y;
                    double endX = curves[i].End.X;
                    double endY = curves[i].End.Y;
                    double controlX = curves[i].Control.X;
                    double controlY = curves[i].Control.Y;

                    if (i == 0)
                        xmlLines.Add($"<AbsoluteMoveTo> <EndPoint X=\"{startX}\" Y=\"{startY}\" /> </AbsoluteMoveTo>");
                    else
                        xmlLines.Add($"<AbsoluteLineTo> <EndPoint X=\"{startX}\" Y=\"{startY}\" /> </AbsoluteLineTo>");

                    xmlLines.Add($"<AbsoluteQuadraticCurveTo> <EndPoint X=\"{endX}\" Y=\"{endY}\" /> <ControlPoint X=\"{controlX}\" Y=\"{controlY}\" /> </AbsoluteQuadraticCurveTo>");
                }
                        
                xmlLines.Add($"<AbsoluteLineTo> <EndPoint X=\"{curves[0].Start.X}\" Y=\"{curves[0].Start.Y}\" /> </AbsoluteLineTo>");
                
                string text = string.Join(Environment.NewLine, xmlLines);
                File.WriteAllText(OutputXmlPath, text);
            }

            return 0;
        }

        private static IEnumerable<Vector> FlattenCurvePoints(IEnumerable<Curve> curves)
        {
            var points = new List<Vector>();

            foreach (var curve in curves)
            {
                points.Add(curve.Start);
                points.Add(curve.Control);
                points.Add(curve.End);
            }

            return points;
        }

        private static void WriteCurvesToFile(string path, IEnumerable<Curve> curves)
        {
            File.WriteAllText(path, JsonConvert.SerializeObject(curves, Formatting.Indented));
        }
        
        private static List<Vector> GetVertices(string path)
        {
            string rawInput = File.ReadAllText(path);
            return JsonConvert.DeserializeObject<List<Vector>>(rawInput);
        }
    }
}
