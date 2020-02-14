namespace KbUtil.Lib.SvgGeneration.Internal
{
    using System;
    using System.IO;
    using System.Xml;
    using KbUtil.Lib.Models.Keyboard;
    using KbUtil.Lib.SvgGeneration.Internal.Path;

    internal class GroupWriter : IElementWriter<Group>
    {
        public SvgGenerationOptions GenerationOptions { get; set; }

        public void Write(XmlWriter writer, Group group)
        {
            if (!group.Visible)
                return;

            writer.WriteStartElement("g");

            var elementWriter = new ElementWriter { GenerationOptions = GenerationOptions };

            elementWriter.WriteAttributes(writer, group);
            WriteAttributes(writer, group);

            elementWriter.WriteSubElements(writer, group);
            WriteSubElements(writer, group);

            writer.WriteEndElement();
        }

        public void WriteAttributes(XmlWriter writer, Group group)
        {
        }

        public void WriteSubElements(XmlWriter writer, Group group)
        {
            foreach (Element child in group.Children)
            {
                switch (child)
                {
                    case var _ when child is Keyboard:
                        throw new InvalidDataException("Keyboard is not a valid child type.");
                    case Key key when child is Key:
                        var keyWriter = new KeyWriter { GenerationOptions = GenerationOptions };
                        keyWriter.Write(writer, key);
                        break;
                    case Spacer spacer when child is Spacer:
                        var spacerWriter = new SpacerWriter { GenerationOptions = GenerationOptions };
                        spacerWriter.Write(writer, spacer);
                        break;
                    case Stack stack when child is Stack:
                        var stackWriter = new StackWriter { GenerationOptions = GenerationOptions };
                        stackWriter.Write(writer, stack);
                        break;
                    case Models.Path.Path path when child is Models.Path.Path:
                        var pathWriter = new PathWriter { GenerationOptions = GenerationOptions };
                        pathWriter.Write(writer, path);
                        break;
                    case Circle circle when child is Circle:
                        var holeWriter = new CircleWriter { GenerationOptions = GenerationOptions };
                        holeWriter.Write(writer, circle);
                        break;
                    case Text text when child is Text:
                        var textWriter = new TextWriter { GenerationOptions = GenerationOptions };
                        textWriter.Write(writer, text);
                        break;
                    case Group subGroup when child is Group:
                        Write(writer, subGroup);
                        break;
                    default:
                        throw new NotSupportedException();
                }
            }
        }
    }
}
