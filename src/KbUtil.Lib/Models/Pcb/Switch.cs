namespace KbUtil.Lib.Models.Pcb
{
    using Newtonsoft.Json;

    public class Switch
    {
        [JsonProperty("x")]
        public float X { get; set; }

        [JsonProperty("y")]
        public float Y { get; set; }

        [JsonProperty("row")]
        public int Row { get; set; }

        [JsonProperty("column")]
        public int Column { get; set; }

        [JsonProperty("rotation")]
        public float Rotation { get; set; }
    }
}
