using System;
using System.Collections.Generic;
using System.IO;
using KbUtil.Console.Models;
using Newtonsoft.Json;

namespace KbUtil.Console.Commands
{
    using Microsoft.Extensions.CommandLineUtils;

    using KbUtil.Console.Services;
    using KbUtil.Lib.Models.Keyboard;

    internal class GenerateSwitchDataCommand
    {
        private readonly IKeyboardDataService _keyboardDataService;

        private readonly CommandArgument _inputPathArgument;
        private readonly CommandArgument _outputPathArgument;

        public GenerateSwitchDataCommand(
            IApplicationService applicationService,
            IKeyboardDataService keyboardDataService)
        {
            _keyboardDataService = keyboardDataService;

            CommandLineApplication command = ApplicationContext.CommandLineApplication
                .Command("gen-switches", config =>
                {
                    config.Description = "Generate a switch data JSON file from an XML input file.";
                    config.OnExecute(() => Execute());
                });

            _inputPathArgument = command.Argument("<input-path>", "The path to the keyboard layout data file.");
            _outputPathArgument = command.Argument("<output-path>", "The path to the generated JSON file.");
        }

        public string InputPath => _inputPathArgument.Value;

        public string OutputPath => _outputPathArgument.Value;

        public int Execute()
        {
            Keyboard keyboard = _keyboardDataService.GetKeyboardData(InputPath);

            var keys = GetKeys(keyboard);
            
            var switchGeometry = GetAbsoluteGeometry(keys);
            
            File.WriteAllText(OutputPath, JsonConvert.SerializeObject(switchGeometry, Formatting.Indented));

            return 0;
        }

        private IEnumerable<Key> GetKeys(Keyboard keyboard)
        {
            var keys = new List<Key>();
            
            foreach (Layer layer in keyboard.Layers)
            {
                foreach (Group group in layer.Groups)
                {
                    keys.AddRange(GetKeysInGroup(group));
                }
            }

            return keys;
        }

        private IEnumerable<Key> GetKeysInGroup(Group group)
        {
            var keys = new List<Key>();
            
            foreach (Element child in group.Children)
            {
                switch (child)
                {
                    case Key key:
                        keys.Add(key);
                        break;
                    case Group group1:
                        keys.AddRange(GetKeysInGroup(group1));
                        break;
                }
            }

            return keys;
        }

        private IEnumerable<Switch> GetAbsoluteGeometry(IEnumerable<Key> keys)
        {
            var results = new List<Switch>();
            
            foreach (Key key in keys)
            {
                
                var tokens = key.Name.Split("Key");
                int column = int.Parse(tokens[0].Replace("Column", "")) - 1;
                int row = int.Parse(tokens[1]) - 1;

                var x = key.XOffset;
                var y = key.YOffset;
                if (key.Parent is Stack stack)
                {
                    switch (stack.Orientation)
                    {
                        case StackOrientation.Horizontal:
                            x += GetOffsetInStack(stack, key);
                            break;
                        case StackOrientation.Vertical:
                            y += GetOffsetInStack(stack, key);
                            break;
                        default:
                            throw new ArgumentOutOfRangeException();
                    }
                }

                var rotation = key.Rotation;
                var parent = key.Parent;
                
                while (parent != null && !(parent is Keyboard))
                {
                    var magnitude = Math.Pow(Math.Pow(x, 2) + Math.Pow(y, 2), 0.5);
                    if (Math.Abs(parent.Rotation) > 0.01)
                    {
                        rotation = rotation - parent.Rotation;
                    }
                    x = (float) (parent.XOffset + magnitude * Math.Sin(Radians(rotation)));
                    y = (float) (parent.YOffset + magnitude * Math.Cos(Radians(rotation)));
                    rotation = 0.0f;

                    parent = parent.Parent;
                }
                
                Switch sw = new Switch
                {
                    Column = column,
                    Row = row,
                    X = x,
                    Y = y,
                    Rotation = (int) rotation,
                    DiodePosition = "bottom"
                };
                
                results.Add(sw);
            }

            return results;
        }

        private static float Radians(float degrees)
        {
            return (float) ((Math.PI / 180) * degrees);
        }

        private static float GetOffsetInStack(Stack stack, Element element)
        {
            switch (stack.Orientation)
            {
                case StackOrientation.Vertical:
                    float stackHeight = default;
                    float dy = float.MaxValue;

                    foreach (Element child in stack.Children)
                    {
                        if (child == element)
                        {
                            dy = stackHeight + child.Height / 2 + child.Margin;
                        }

                        stackHeight += child.Height + child.Margin * 2;
                    }

                    return -(stackHeight / 2 - dy);

                case StackOrientation.Horizontal:
                    float stackWidth = default;
                    float dx = float.MaxValue;

                    foreach (Element child in stack.Children)
                    {
                        if (child == element)
                        {
                            dx = stackWidth + child.Width / 2 + child.Margin;
                        }

                        stackWidth += child.Width + child.Margin * 2;
                    }

                    return -(stackWidth / 2 - dx);

                default:
                    throw new NotSupportedException();
            }
        }
    }
}
