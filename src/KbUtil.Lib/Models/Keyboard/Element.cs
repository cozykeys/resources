namespace KbUtil.Lib.Models.Keyboard
{
    using System.Collections.Generic;

    public abstract class Element
    {
        public string Name { get; set; }
        public float XOffset { get; set; }
        public float YOffset { get; set; }
        public float Rotation { get; set; }
        public virtual float Height { get; set; }
        public virtual float Width { get; set; }
        public float Margin { get; set; }
        public Element Parent { get; set; }
        public IDictionary<string, Constant> Constants { get; set; }
        public bool Debug { get; set; }

        // FIXME: This isn't fully supported by all of the deserializers/writers
        // yet but it's still useful.
        public bool Visible { get; set; } = true;
    }
}
