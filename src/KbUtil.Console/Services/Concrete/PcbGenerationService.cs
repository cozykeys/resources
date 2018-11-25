namespace KbUtil.Console.Services.Concrete
{
    using System.Collections.Generic;
    using KbUtil.Lib.Models.Keyboard;
    using KbUtil.Lib.Models.Pcb;
    using KbUtil.Lib.PcbGeneration;

    internal class PcbGenerationService : IPcbGenerationService
    {
        public void GeneratePcb(List<Switch> switches, string path, PcbGenerationOptions options = null)
            => PcbGenerator.GeneratePcb(switches, path, options ?? new PcbGenerationOptions());
    }
}
