namespace KbUtil.Console.Services
{
    using KbUtil.Lib.Models.Pcb;
    using System.Collections.Generic;

    public interface ISwitchDataService
    {
        List<Switch> GetSwitchData(string inputPath);
    }
}
