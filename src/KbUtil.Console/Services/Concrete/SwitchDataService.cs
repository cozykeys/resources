namespace KbUtil.Console.Services.Concrete
{
    using KbUtil.Lib.Models.Pcb;
    using Newtonsoft.Json;
    using System.Collections.Generic;
    using System.IO;

    internal class SwitchDataService : ISwitchDataService
    {
        public List<Switch> GetSwitchData(string inputPath)
            => JsonConvert
                .DeserializeObject<List<Switch>>(File.ReadAllText(inputPath));
    }
}
