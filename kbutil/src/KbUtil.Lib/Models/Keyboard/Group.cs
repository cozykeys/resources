namespace KbUtil.Lib.Models.Keyboard
{
    using KbUtil.Lib.Models.Attributes;
    using System.Collections.Generic;

    [GroupChild]
    public class Group : Element
    {
        public IEnumerable<Element> Children { get; set; }
    }
}
