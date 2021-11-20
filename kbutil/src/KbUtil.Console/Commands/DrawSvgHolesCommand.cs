using System.Globalization;
using System.Linq;
using KbUtil.Lib.Geometry2D.Extensions;

namespace KbUtil.Console.Commands
{
    using Microsoft.Extensions.CommandLineUtils;
    using System.Collections.Generic;
    using System.IO;
    
    using Newtonsoft.Json;
    
    using KbUtil.Lib.Geometry2D.Models;
    using Models;
    using Services;

    internal class DrawSvgHolesCommand
    {
        private readonly ISvgService _svgService;
        
        private readonly CommandArgument _inputPathArgument;
        private readonly CommandArgument _outputPathArgument;
        
        private static readonly Dictionary<string, string> HoleStyle = new Dictionary<string, string>
        {
            { "r", (3.563 / 2.0).ToString(CultureInfo.InvariantCulture) },
            { "stroke", "black" },
            { "stroke-width", "0.1" },
            { "fill", "none" }
        };
        
        public DrawSvgHolesCommand(ISvgService svgService)
        {
            _svgService = svgService;
            
            CommandLineApplication command = ApplicationContext.CommandLineApplication
                .Command("draw-svg-holes", config =>
                {
                    config.Description = "TODO";
                    config.OnExecute(() => Execute());
                });
            
            _inputPathArgument = command.Argument("<input-path>", "TODO");
            _outputPathArgument = command.Argument("<output-path>", "TODO");
        }

        private string InputPath => _inputPathArgument.Value;
        private string OutputPath => _outputPathArgument.Value;

        public int Execute()
        {
            IEnumerable<Segment> segments = GetSegments(InputPath);

            var midpoints = segments.Select(segment => segment.Midpoint());

            Svg svg = _svgService.CreateSvg();
            _svgService.Append(svg, midpoints, HoleStyle);
            _svgService.WriteToFile(svg, OutputPath);

            return 0;
        }

        private static IEnumerable<Segment> GetSegments(string path)
        {
            string rawInput = File.ReadAllText(path);
            return JsonConvert.DeserializeObject<List<Segment>>(rawInput);
        }
    }
}
