namespace KbMath.Console.Models
{
    using Newtonsoft.Json;

    public class Switch
    {
        [JsonProperty("row")]
        public int Row { get; set; }
        [JsonProperty("column")]
        public int Column { get; set; }
        [JsonProperty("x")]
        public double X { get; set; }
        [JsonProperty("y")]
        public double Y { get; set; }
        [JsonProperty("rotation")]
        public double Rotation { get; set; }
    }
}