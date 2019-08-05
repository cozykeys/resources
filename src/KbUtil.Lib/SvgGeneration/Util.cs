namespace KbUtil.Lib.SvgGeneration
{
    using System;
    
    using KbMath.Core.Geometry2D.Extensions;
    using KbMath.Core.Geometry2D.Models;
    using KbUtil.Lib.Models.Geometry;
    using KbUtil.Lib.Models.Keyboard;

    public static class Util
    {
        public static Bearing GetAbsoluteBearing(Element element)
        {
            var childPosition = new Vector(element.XOffset, element.YOffset);
            float childRotation = element.Rotation;

            Element child = element;
            var parent = element.Parent;
            while (!(parent is Keyboard))
            {
                if (parent is Stack stack)
                {
                    var offset = Util.GetOffsetInStack(stack, child);
                    switch (stack.Orientation)
                    {
                        case StackOrientation.Horizontal:
                            childPosition = new Vector(childPosition.X + offset, childPosition.Y);
                            break;
                        case StackOrientation.Vertical:
                            childPosition = new Vector(childPosition.X, childPosition.Y + offset);
                            break;
                        default:
                            throw new ArgumentOutOfRangeException();
                    }
                }
                
                var parentPosition = new Vector(parent.XOffset, parent.YOffset);
                
                var parentRotationAngleRadians = -parent.Rotation.ToRadians();
                var rotationMatrix = new Matrix2x2(new double[2,2]
                {
                    { Math.Cos(parentRotationAngleRadians), -Math.Sin(parentRotationAngleRadians) },
                    { Math.Sin(parentRotationAngleRadians), Math.Cos(parentRotationAngleRadians) }
                });
                
                childPosition = (rotationMatrix * childPosition) + parentPosition;
                childRotation = childRotation + parent.Rotation;

                child = parent;
                parent = parent.Parent;
            }

            return new Bearing { Rotation = childRotation, Position = childPosition};
        }
        
        public static float GetOffsetInStack(Stack stack, Element element)
        {
            switch (stack.Orientation)
            {
                case StackOrientation.Vertical:
                    float stackHeight = default;
                    float dy = float.MaxValue;

                    foreach (Element child in stack.Children)
                    {
                        if (child == element)
                        {
                            dy = stackHeight + child.Height / 2 + child.Margin;
                        }

                        stackHeight += child.Height + child.Margin * 2;
                    }

                    return -(stackHeight / 2 - dy);

                case StackOrientation.Horizontal:
                    float stackWidth = default;
                    float dx = float.MaxValue;

                    foreach (Element child in stack.Children)
                    {
                        if (child == element)
                        {
                            dx = stackWidth + child.Width / 2 + child.Margin;
                        }

                        stackWidth += child.Width + child.Margin * 2;
                    }

                    return -(stackWidth / 2 - dx);

                default:
                    throw new NotSupportedException();
            }
        }
    }
}

