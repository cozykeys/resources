﻿namespace KbUtil.Console.Commands
{
    using KbUtil.Console.Services;
    using KbUtil.Lib.PcbGeneration;
    using Microsoft.Extensions.CommandLineUtils;

    internal class GeneratePcbCommand
    {
        private readonly ISwitchDataService _switchDataService;
        private readonly IPcbGenerationService _pcbGenerationService;

        private readonly CommandArgument _keyboardNameArgument;
        private readonly CommandArgument _inputPathArgument;
        private readonly CommandArgument _outputPathArgument;

        public GeneratePcbCommand(
            IApplicationService applicationService,
            ISwitchDataService switchDataService,
            IPcbGenerationService pcbGenerationService)
        {
            _switchDataService = switchDataService;
            _pcbGenerationService = pcbGenerationService;

            CommandLineApplication command = ApplicationContext.CommandLineApplication
                .Command("gen-pcb", config =>
                {
                    config.Description = "Generate a Kicad PCB file from an XML input file.";
                    config.OnExecute(() => Execute());
                });
            
            _keyboardNameArgument = command.Argument("<keyboard-name>", "The keyboard name (Used to look up templates).");
            _inputPathArgument = command.Argument("<input-path>", "The path to the keyboard data file.");
            _outputPathArgument = command.Argument("<output-path>", "The path to the generated PCB file.");
        }

        public string KeyboardName => _keyboardNameArgument.Value;

        public string InputPath => _inputPathArgument.Value;

        public string OutputPath => _outputPathArgument.Value;

        public int Execute()
        {
            var switches = _switchDataService.GetSwitchData(InputPath);

            var generationOptions = new PcbGenerationOptions();

            _pcbGenerationService.GeneratePcb(KeyboardName, switches, OutputPath, generationOptions);

            return 0;
        }
    }
}
