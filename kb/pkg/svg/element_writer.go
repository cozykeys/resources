package svg

import (
	"fmt"
	"kb/pkg/models"
	"strings"

	"github.com/beevik/etree"
)

func writeTransform(e *etree.Element, element models.KeyboardElement) {
	transformStrings := []string{}

	xOffset := element.GetXOffset()
	yOffset := element.GetYOffset()

	// TODO: Handle case when parent is a stack
	////        if (element.Parent is Stack)
	////        {
	////            switch (((Stack) element.Parent).Orientation)
	////            {
	////                case StackOrientation.Horizontal:
	////                    xOffset += SvgGeneration.Util.GetOffsetInStack((Stack)element.Parent, element);
	////                    break;
	////                case StackOrientation.Vertical:
	////                    yOffset += SvgGeneration.Util.GetOffsetInStack((Stack)element.Parent, element);
	////                    break;
	////                default:
	////                    throw new ArgumentOutOfRangeException();
	////            }
	////        }

	if xOffset != 0.0 || yOffset != 0.0 {
		transformStrings = append(transformStrings, fmt.Sprintf("translate(%f,%f)", xOffset, yOffset))
	}

	if element.GetRotation() != 0.0 {
		transformStrings = append(transformStrings, fmt.Sprintf("rotate(%f)", element.GetRotation()))
	}

	if len(transformStrings) > 0 {
		e.CreateAttr("transform", strings.Join(transformStrings, " "))
	}
}

func writeDebugOverlay(e *etree.Element, element models.KeyboardElement) {
	// Write border (With Margin)
	{
		h := element.GetHeight() + (element.GetMargin() * 2)
		w := element.GetWidth() + (element.GetMargin() * 2)
		path := e.CreateElement("path")
		path.CreateAttr("id", fmt.Sprintf("%sDebugOverlayMargin", element.GetName()))
		path.CreateAttr("d", fmt.Sprintf("M -%f,-%f h %f v %f h -%f v -%f h %f", w/2, h/2, w, h, w, h, w))
		path.CreateAttr("style", "fill:none;stroke:#00ff00;stroke-width:0.1")
	}

	// Write border (Without Margin)
	{
		h := element.GetHeight()
		w := element.GetWidth()
		path := e.CreateElement("path")
		path.CreateAttr("id", fmt.Sprintf("%sDebugOverlayBorder", element.GetName()))
		path.CreateAttr("d", fmt.Sprintf("M -%f,-%f h %f v %f h -%f v -%f h %f", w/2, h/2, w, h, w, h, w))
		path.CreateAttr("style", "fill:none;stroke:#00ffff;stroke-width:0.1")
	}

	// Write center axes
	{
		path := e.CreateElement("path")
		path.CreateAttr("id", fmt.Sprintf("%sCenterAxes", element.GetName()))
		path.CreateAttr("d", "m 0,0 h 5 h -10 h 5 v 5 v -10 v 5")
		path.CreateAttr("style", "fill:none;stroke:#00ffff;stroke-width:0.1")
	}
}
