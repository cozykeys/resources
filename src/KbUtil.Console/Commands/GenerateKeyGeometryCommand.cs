using System.IO;
using KbUtil.Lib.Models.Geometry;
using Newtonsoft.Json;

namespace KbUtil.Console.Commands
{
    using System;
    using System.Collections.Generic;
    using Microsoft.Extensions.Logging;

    using Microsoft.Extensions.CommandLineUtils;

    using KbUtil.Console.Services;
    using KbUtil.Lib.SvgGeneration;
    using KbUtil.Lib.Models.Keyboard;
    
    using KbMath.Core.Geometry2D.Models;

    internal class GenerateKeyBearingsCommand
    {
        private readonly IKeyboardDataService _keyboardDataService;
        private readonly ILogger _logger;

        private readonly CommandArgument _inputPathArgument;
        private readonly CommandArgument _outputPathArgument;
        
        private readonly CommandOption _debugSvg;

        public GenerateKeyBearingsCommand(
            IApplicationService applicationService,
            IKeyboardDataService keyboardDataService,
            ILoggerFactory loggerFactory)
        {
            _logger = loggerFactory.CreateLogger(nameof(GenerateKeyBearingsCommand));
            
            _keyboardDataService = keyboardDataService;

            Command = applicationService.CommandLineApplication
                .Command("gen-key-bearings", config =>
                {
                    config.Description = "Print the key geometry from an XML input file.";
                    config.ExtendedHelpText = "TODO";
                    config.OnExecute(() => Execute());
                });

            _inputPathArgument = Command.Argument("<input-path>", "The path to the keyboard layout data file.");
            _outputPathArgument = Command.Argument("<output-path>", "The path to the generated JSON geometry data file.");
            
            _debugSvg = Command.Option(
                "--debug-svg",
                "TODO",
                CommandOptionType.SingleValue);
        }

        public CommandLineApplication Command { get; }

        public string InputPath => _inputPathArgument.Value;
        
        public string OutputPath => _outputPathArgument.Value;
        
        private string DebugSvgPath => _debugSvg?.Value();

        private IEnumerable<Key> EnumerateKeys(Group group)
        {
            var keysInGroup = new List<Key>();
            foreach (var child in group.Children)
            {
                switch (child)
                {
                    case Key key:
                        _logger.LogDebug($"Found key {key.Name}");
                        keysInGroup.Add(key);
                        continue;
                    case Group subgroup:
                    {
                        var keysInSubgroup = EnumerateKeys(subgroup);
                        keysInGroup.AddRange(keysInSubgroup);
                        break;
                    }
                }
            }
            
            return keysInGroup;
        }
        
        private IEnumerable<Key> EnumerateKeys(Keyboard keyboard)
        {
            var keys = new List<Key>();
            foreach (var layer in keyboard.Layers)
            {
                _logger.LogDebug($"Enumerating keys on layer {layer.Name}");

                var keysInLayer = new List<Key>();
                foreach (var group in layer.Groups)
                {
                    var keysInGroup = EnumerateKeys(group);
                    keysInLayer.AddRange(keysInGroup);
                }
                keys.AddRange(keysInLayer);
            }
            return keys;
        }
        
        public int Execute()
        {
            Keyboard keyboard = _keyboardDataService.GetKeyboardData(InputPath);

            var keys = EnumerateKeys(keyboard);

            var keyGeometry = new Dictionary<string, Bearing>();

            int i = 0;
            foreach (var key in keys)
            {
                string keyName = string.IsNullOrWhiteSpace(key.Name) ? $"UnnamedKey{i++}" : key.Name;
                keyGeometry[keyName] = Util.GetAbsoluteBearing(key);
            }

            if (!string.IsNullOrEmpty(DebugSvgPath))
            {
                var svgData = new List<string>
                {
                    "<svg width=\"500mm\" height=\"500mm\" viewBox=\"0 0 500 500\" xmlns=\"http://www.w3.org/2000/svg\">"
                };

                foreach (var (name, bearing) in keyGeometry)
                {
                    svgData.Add($"  <g id=\"{name}\" transform=\"translate({bearing.Position.X},{bearing.Position.Y}) rotate({bearing.Rotation})\">");
                    svgData.Add($"    <path d=\"M -9.05 -9.05 L 9.05 -9.05 L 9.05 9.05 L -9.05 9.05 L -9.05 -9.05\"/>");
                    svgData.Add($"  </g>");
                }

                svgData.Add("</svg>");
                System.IO.File.WriteAllText(DebugSvgPath, string.Join(Environment.NewLine, svgData));
            }

            File.WriteAllText(OutputPath, JsonConvert.SerializeObject(keyGeometry, Formatting.Indented));

            return 0;
        }
    }
}
