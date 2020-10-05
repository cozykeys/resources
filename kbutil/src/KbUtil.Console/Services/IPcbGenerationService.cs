namespace KbUtil.Console.Services
{
    using KbUtil.Lib.Models.Pcb;
    using KbUtil.Lib.PcbGeneration;
    using System.Collections.Generic;

    internal interface IPcbGenerationService
    {
        void GeneratePcb(string keyboardName, List<Switch> switches, string path, PcbGenerationOptions options = null);
    }
}
