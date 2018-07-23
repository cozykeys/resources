namespace KbUtil.Console.Commands
{
    using KbUtil.Console.Services;
    using KbUtil.Lib.PcbGeneration;
    using Microsoft.Extensions.CommandLineUtils;

    internal class GeneratePcbCommand
    {
        private readonly ISwitchDataService _switchDataService;
        private readonly IPcbGenerationService _pcbGenerationService;

        private readonly CommandArgument _inputPathArgument;
        private readonly CommandArgument _outputPathArgument;

        public GeneratePcbCommand(
            IApplicationService applicationService,
            ISwitchDataService switchDataService,
            IPcbGenerationService pcbGenerationService)
        {
            _switchDataService = switchDataService;
            _pcbGenerationService = pcbGenerationService;

            Command = applicationService.CommandLineApplication
                .Command("gen-pcb", config =>
                {
                    config.Description = "Generate a Kicad PCB file from an XML input file.";
                    config.ExtendedHelpText = "TODO";
                    config.OnExecute(() => Execute());
                });

            _inputPathArgument = Command.Argument("<input-path>", "The path to the keyboard data file.");
            _outputPathArgument = Command.Argument("<output-path>", "The path to the generated PCB file.");
        }

        public CommandLineApplication Command { get; }

        public string InputPath => _inputPathArgument.Value;

        public string OutputPath => _outputPathArgument.Value;

        public int Execute()
        {
            var switches = _switchDataService.GetSwitchData(InputPath);

            var generationOptions = new PcbGenerationOptions();

            _pcbGenerationService.GeneratePcb(switches, OutputPath, generationOptions);

            return 0;
        }
    }
}
