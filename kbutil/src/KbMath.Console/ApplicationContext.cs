namespace KbMath.Console
{
    using Microsoft.Extensions.CommandLineUtils;

    public static class ApplicationContext
    {
        public static CommandLineApplication CommandLineApplication { get; } = new CommandLineApplication
        {
            Name = "Lorem Ipsum",
            FullName = "Lorem Ipsum",
            Description = "Lorem Ipsum",
        };
    }
}
