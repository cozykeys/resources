namespace KbUtil.Lib.Models.Keyboard
{
    using KbUtil.Lib.Models.Attributes;
    using System.Collections.Generic;

    [GroupChild]
    public class Key : Element
    {
        public IEnumerable<Legend> Legends { get; set; }

        public int Row { get; set; }
        public int Column { get; set; }

        /// TODO: Strongly type these instead of using string
        public string Fill { get; set; }
        public string Stroke { get; set; }
    }
}
