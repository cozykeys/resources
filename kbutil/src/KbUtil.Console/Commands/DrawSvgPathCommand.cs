namespace KbUtil.Console.Commands
{
    using Microsoft.Extensions.CommandLineUtils;
    using System.Collections.Generic;
    using System.IO;
    
    using Newtonsoft.Json;
    
    using KbUtil.Lib.Geometry2D.Models;
    using Models;
    using Services;

    internal class DrawSvgPathCommand
    {
        private readonly ISvgService _svgService;
        
        private readonly CommandArgument _inputPathArgument;
        private readonly CommandArgument _outputPathArgument;
        
        private static readonly Dictionary<string, string> PathStyle = new Dictionary<string, string>
        {
            { "stroke", "black" },
            { "stroke-width", "0.1" },
            { "fill", "none" }
        };
        
        public DrawSvgPathCommand(ISvgService svgService)
        {
            _svgService = svgService;
            
            CommandLineApplication command = ApplicationContext.CommandLineApplication
                .Command("draw-svg-path", config =>
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
            IEnumerable<Curve> curves = GetCurves(InputPath);

            Svg svg = _svgService.CreateSvg();
            _svgService.DrawPath(svg, curves, PathStyle);
            _svgService.WriteToFile(svg, OutputPath);

            return 0;
        }

        private static IEnumerable<Curve> GetCurves(string path)
        {
            string rawInput = File.ReadAllText(path);
            return JsonConvert.DeserializeObject<List<Curve>>(rawInput);
        }
    }
}
