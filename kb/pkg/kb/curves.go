package kb

import (
	"fmt"
	"kb/pkg/geo"
	"kb/pkg/geosvg"
	"kb/pkg/svg"
	"strings"

	"github.com/beevik/etree"
)

// GenerateCurves generates a set of symetric quadratic curves from a set of
// points.
func GenerateCurves(
	points []*geo.Point3D,
	distance float64,
) ([]*geo.Curve3D, error) {
	curves := make([]*geo.Curve3D, len(points))

	relativeIndex := func(curr, diff, count int) int {
		index := (curr + diff) % count
		if index < 0 {
			return count + index
		} else {
			return index
		}
	}

	for curr := 0; curr < len(points); curr++ {
		prev := relativeIndex(curr, -1, len(points))
		next := relativeIndex(curr, 1, len(points))

		s1 := &geo.Segment3D{P1: points[curr], P2: points[prev]}
		s2 := &geo.Segment3D{P1: points[curr], P2: points[next]}

		start := s1.GetPoint(distance)
		end := s2.GetPoint(distance)

		curves[curr] = &geo.Curve3D{
			Start:   start,
			End:     end,
			Control: points[curr],
		}
	}

	return curves, nil
}

// TODO: We should move SVG writing into the geo package or maybe like a geosvg
// package. Either way it shouldn't be entirely here.
func WriteCurvesSVG(
	path string,
	vertices []*geo.Point3D,
	curves []*geo.Curve3D,
) error {
	// Create XML Document
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	// Set up the root SVG element
	box := geo.GetBoundingBox(vertices)

	w := (box.P2.X - box.P1.X) + 10
	h := (box.P2.Y - box.P1.Y) + 10
	root := doc.CreateElement("svg")
	root.CreateAttr("width", fmt.Sprintf("%fmm", w))
	root.CreateAttr("height", fmt.Sprintf("%fmm", h))
	root.CreateAttr("viewBox",
		fmt.Sprintf("%f %f %f %f", box.P1.X-5, box.P1.Y-5, w, h))
	root.CreateAttr("xmlns", "http://www.w3.org/2000/svg")

	g1 := root.CreateElement("g")
	g1.CreateAttr("id", "vertices")

	p1 := g1.CreateElement("path")
	p1.CreateAttr("id", "path-vertices")
	p1Data := make([]string, len(vertices)+1)
	for i := 0; i < len(vertices); i++ {
		if i == 0 {
			p1Data[i] = fmt.Sprintf("M %f,%f", vertices[i].X, vertices[i].Y)
		} else {
			p1Data[i] = fmt.Sprintf("L %f,%f", vertices[i].X, vertices[i].Y)
		}
	}
	p1Data[len(vertices)] = fmt.Sprintf("L %f,%f", vertices[0].X, vertices[0].Y)
	p1.CreateAttr("d", strings.Join(p1Data, " "))
	p1.CreateAttr("style", svg.StyleMap{
		"stroke":       "blue",
		"stroke-width": "0.2",
		"fill":         "transparent",
	}.String())

	for i, vertex := range vertices {
		c := g1.CreateElement("circle")
		c.CreateAttr("id", fmt.Sprintf("vertex-%d", i))
		c.CreateAttr("transform", fmt.Sprintf("translate(%f,%f)", vertex.X, vertex.Y))
		c.CreateAttr("style", svg.StyleMap{
			"r":            "0.5",
			"stroke":       "black",
			"stroke-width": "0.1",
			"fill":         "blue",
		}.String())
	}

	for _, curve := range curves {
		geosvg.PathCurve3D(g1, curve, svg.StyleMap{
			"stroke":       "black",
			"stroke-width": "0.1",
			"fill":         "transparent",
		})
	}

	return WriteSVGToFile(doc, path)
}
