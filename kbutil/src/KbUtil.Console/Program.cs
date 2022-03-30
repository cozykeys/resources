namespace KbUtil.Console
{
    using System;
    using Microsoft.Extensions.CommandLineUtils;
    using Microsoft.Extensions.DependencyInjection;
    using Microsoft.Extensions.Logging;

    using KbUtil.Console.Commands;
    using KbUtil.Console.Services;
    using KbUtil.Console.Services.Concrete;

    public static class Program
    {
        public static int Main(string[] args)
        {
            ServiceProvider serviceProvider = new ServiceCollection()
                .AddServices()
                .BuildServiceProvider();

            CreateCommands(serviceProvider);

            var loggerFactory = serviceProvider.GetService<ILoggerFactory>();
            Microsoft.Extensions.Logging.ILogger logger = loggerFactory.CreateLogger(nameof(Program));

            try
            {
                logger.LogInformation("Running...");
                return ApplicationContext.CommandLineApplication.Execute(args);
            }
            catch (Exception ex)
            {
                logger.LogCritical($"Unhandled {ex}");
                return -1;
            }
        }

        private static IServiceCollection AddServices(this IServiceCollection serviceCollection)
        {
            var loggerFactory = LoggerFactory.Create(builder => builder
                    .AddConsole());

            serviceCollection.AddSingleton(loggerFactory);
            serviceCollection.AddSingleton<IAssemblyService, AssemblyService>();
            serviceCollection.AddSingleton<IEnvironmentService, EnvironmentService>();
            serviceCollection.AddSingleton<IApplicationService, ApplicationService>();
            serviceCollection.AddSingleton<IKeyboardDataService, KeyboardDataService>();
            serviceCollection.AddSingleton<ISwitchDataService, SwitchDataService>();
            serviceCollection.AddSingleton<IFileService, FileService>();
            serviceCollection.AddSingleton<ISvgGenerationService, SvgGenerationService>();

            // From KbMath
            serviceCollection.AddSingleton<ISvgService, SvgService>();

            return serviceCollection;
        }

        private static void CreateCommands(IServiceProvider serviceProvider)
        {
            ActivatorUtilities.CreateInstance<GenerateSvgCommand>(serviceProvider);
            ActivatorUtilities.CreateInstance<GenerateKeyBearingsCommand>(serviceProvider);

            // From KbMath
            ActivatorUtilities.CreateInstance<ExpandVerticesCommand>(serviceProvider);
            ActivatorUtilities.CreateInstance<ExpandVerticesCommand2>(serviceProvider);
            ActivatorUtilities.CreateInstance<GenerateCurvesCommand>(serviceProvider);
            ActivatorUtilities.CreateInstance<DrawSvgPathCommand>(serviceProvider);
            ActivatorUtilities.CreateInstance<DrawSvgHolesCommand>(serviceProvider);
            ActivatorUtilities.CreateInstance<DrawSwitchesCommand>(serviceProvider);
            ActivatorUtilities.CreateInstance<ScratchCommand>(serviceProvider);
        }
    }
}
