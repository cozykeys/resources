using Newtonsoft.Json;

namespace KbUtil.Console.Models
{
    public class Switch
    {
        [JsonProperty("row")]
        public int Row { get; set; }
        
        [JsonProperty("column")]
        public int Column { get; set; }
        
        [JsonProperty("x")]
        public float X { get; set; }
        
        [JsonProperty("y")]
        public float Y { get; set; }
        
        [JsonProperty("rotation")]
        public int Rotation { get; set; }
        
        [JsonProperty("diode_position")]
        public string DiodePosition  { get; set; }
    }
}