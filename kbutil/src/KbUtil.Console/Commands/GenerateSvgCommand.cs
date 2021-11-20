namespace KbUtil.Console.Commands
{
    using Microsoft.Extensions.CommandLineUtils;

    using KbUtil.Console.Services;
    using KbUtil.Lib.Models.Keyboard;
    using KbUtil.Lib.SvgGeneration;

    internal class GenerateSvgCommand
    {
        private readonly IKeyboardDataService _keyboardDataService;
        private readonly ISvgGenerationService _svgGenerationService;

        private readonly CommandArgument _inputPathArgument;
        private readonly CommandArgument _outputPathArgument;

        private readonly CommandOption _keycapOverlayOption;
        private readonly CommandOption _visualSwitchCutoutsOption;
        private readonly CommandOption _keycapLegendsOption;
        private readonly CommandOption _squashOption;

        public GenerateSvgCommand(
            IApplicationService applicationService,
            IKeyboardDataService keyboardDataService,
            ISvgGenerationService svgGenerationService)
        {
            _keyboardDataService = keyboardDataService;
            _svgGenerationService = svgGenerationService;

            CommandLineApplication command = ApplicationContext.CommandLineApplication
                .Command("gen-svg", config =>
                {
                    config.Description = "Generate an SVG file from an XML input file.";
                    config.OnExecute(() => Execute());
                });

            _inputPathArgument = command.Argument("<input-path>", "The path to the keyboard layout data file.");
            _outputPathArgument = command.Argument("<output-path>", "The path to the generated SVG file.");

            _visualSwitchCutoutsOption = command.Option(
                "--visual-switch-cutouts",
                "Write switch cutouts with thicker path strokes.",
                CommandOptionType.NoValue);

            _keycapOverlayOption = command.Option(
                "--keycap-overlays",
                "Write keycap overlays to the generated SVG file.",
                CommandOptionType.NoValue);

            _keycapLegendsOption = command.Option(
                "--keycap-legends",
                "Write keycap legends to the generated SVG file.",
                CommandOptionType.NoValue);

            _squashOption = command.Option(
                "--squash",
                "Squash all layers into a single SVG file.",
                CommandOptionType.NoValue);
        }

        public string InputPath => _inputPathArgument.Value;

        public string OutputPath => _outputPathArgument.Value;

        public bool WriteKeycapOverlays => _keycapOverlayOption != null && _keycapOverlayOption.HasValue();

        public bool WriteKeycapLegends => _keycapLegendsOption != null && _keycapLegendsOption.HasValue();

        public bool SquashLayers => _squashOption != null && _squashOption.HasValue();

        public bool WriteVisualSwitchCutouts => _visualSwitchCutoutsOption != null && _visualSwitchCutoutsOption.HasValue();

        public int Execute()
        {
            Keyboard keyboard = _keyboardDataService.GetKeyboardData(InputPath);

            var generationOptions = new SvgGenerationOptions
            {
                EnableKeycapOverlays = WriteKeycapOverlays,
                EnableLegends = WriteKeycapLegends,
                EnableVisualSwitchCutouts = WriteVisualSwitchCutouts,
                SquashLayers = SquashLayers
            };

            _svgGenerationService.GenerateSvg(keyboard, OutputPath, generationOptions);

            return 0;
        }
    }
}
