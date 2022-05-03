package geosvg

import (
	"fmt"
	"kb/pkg/geo"
	"kb/pkg/svg"

	"github.com/beevik/etree"
)

func PathCurve3D(
	parent *etree.Element,
	c *geo.Curve3D,
	style svg.StyleMap,
) {
	path := parent.CreateElement("path")
	path.CreateAttr("id", "TODO")
	path.CreateAttr("d", PathDataCurve3D(c))
	path.CreateAttr("style", style.String())
}

func PathDataCurve3D(c *geo.Curve3D) string {
	////public void Append(Svg svg, IEnumerable<Curve> curves, Dictionary<string, string> styleOverrides)
	////{
	////	var styles = new Dictionary<string, string>
	////	{
	////		{ "stroke", styleOverrides.ContainsKey("stroke") ? styleOverrides["stroke"] : "black" },
	////		{ "stroke-width", styleOverrides.ContainsKey("stroke-width") ? styleOverrides["stroke-width"] : "0.1" },
	////		{ "fill", styleOverrides.ContainsKey("fill") ? styleOverrides["fill"] : "black" }
	////	};
	////
	////	string style = string.Join(" ", styles.Select(kvp => $"{kvp.Key}=\"{kvp.Value}\""));
	////
	////	foreach (Curve curve in curves)
	////	{
	////		string pathData = $"M {curve.Start.X},{curve.Start.Y} Q {curve.Control.X},{curve.Control.Y}, {curve.End.X},{curve.End.Y}";
	return fmt.Sprintf("M %f,%f Q %f,%f, %f,%f", c.Start.X, c.Start.Y,
		c.Control.X, c.Control.Y, c.End.X, c.End.Y)
	////		svg.Elements.Add(new SvgElement
	////		{
	////			Content = $"<path {style} d=\"{pathData}\" />"
	////		});
	////	}
	////}
}
