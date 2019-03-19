namespace KbMath.Console.Commands
{
    using Microsoft.Extensions.CommandLineUtils;
    using System;

    internal class GenerateSwitchBearingsCommand
    {
        private readonly CommandArgument _inputPathArgument;
        
        public GenerateSwitchBearingsCommand()
        {
            CommandLineApplication command = ApplicationContext.CommandLineApplication
                .Command("generate-switch-bearings", config =>
                {
                    config.Description = "TODO";
                    config.ExtendedHelpText = "TODO";
                    config.OnExecute(() => Execute());
                });
            
            _inputPathArgument = command.Argument("<input-path>", "TODO");
        }

        private string InputPath => _inputPathArgument.Value;

        public int Execute()
        {
            throw new NotImplementedException();
        }
    }
}
