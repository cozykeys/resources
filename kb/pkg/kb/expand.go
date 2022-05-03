package kb

import (
	"fmt"
	"log"
	"strings"

	"kb/pkg/geo"
	"kb/pkg/svg"

	"github.com/beevik/etree"
)

func ConvertPointsToSegments(points []*geo.Point3D) ([]*geo.Segment3D, error) {
	if len(points) < 2 {
		return nil, fmt.Errorf("must have at least 2 points to convert to segments")
	}

	segments := make([]*geo.Segment3D, len(points))
	for curr := 0; curr < len(points); curr++ {
		next := (curr + 1) % len(points)
		segments[curr] = &geo.Segment3D{
			P1: points[curr],
			P2: points[next],
		}
	}

	return segments, nil
}

func ExpandSegments(segments []*geo.Segment3D, distance float64) ([]*geo.Line3D, error) {
	lines := make([]*geo.Line3D, len(segments))

	for i, s := range segments {
		lines[i] = s.PerpendicularClockwise(distance).ToLine()
	}

	return lines, nil
}

func GetIntersectionPoints(lines []*geo.Line3D) ([]*geo.Point3D, error) {
	points := make([]*geo.Point3D, len(lines))

	for curr := 0; curr < len(lines); curr++ {
		prev := curr - 1
		if curr == 0 {
			prev = len(lines) - 1
		}

		l1 := lines[curr]
		l2 := lines[prev]

		intersection := l1.IntersectionWith(l2)
		if intersection == nil {
			return nil, fmt.Errorf("TODO: intersection")
		}

		_, ok := intersection.(*geo.Line3D)
		if ok {
			log.Printf("intersection is a line")
		}

		p, ok := intersection.(*geo.Point3D)
		if !ok {
			return nil, fmt.Errorf("TODO: type cast %v", intersection)
		}

		points[curr] = p
	}

	return points, nil
}

// TODO: We should move SVG writing into the geo package or maybe like a geosvg
// package. Either way it shouldn't be entirely here.
func WriteExpandSVG(path string, before, after []*geo.Point3D) error {
	// Create XML Document
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	// Set up the root SVG element
	box := geo.GetBoundingBox(append(before, after...))

	w := (box.P2.X - box.P1.X) + 10
	h := (box.P2.Y - box.P1.Y) + 10
	root := doc.CreateElement("svg")
	root.CreateAttr("width", fmt.Sprintf("%fmm", w))
	root.CreateAttr("height", fmt.Sprintf("%fmm", h))
	root.CreateAttr("viewBox",
		fmt.Sprintf("%f %f %f %f", box.P1.X-5, box.P1.Y-5, w, h))
	root.CreateAttr("xmlns", "http://www.w3.org/2000/svg")

	g1 := root.CreateElement("g")
	g1.CreateAttr("id", "before")

	p1 := g1.CreateElement("path")
	p1.CreateAttr("id", "path-before")
	p1Data := make([]string, len(before)+1)
	for i := 0; i < len(before); i++ {
		if i == 0 {
			p1Data[i] = fmt.Sprintf("M %f,%f", before[i].X, before[i].Y)
		} else {
			p1Data[i] = fmt.Sprintf("L %f,%f", before[i].X, before[i].Y)
		}
	}
	p1Data[len(before)] = fmt.Sprintf("L %f,%f", before[0].X, before[0].Y)
	p1.CreateAttr("d", strings.Join(p1Data, " "))
	p1.CreateAttr("style", svg.StyleMap{
		"stroke":       "blue",
		"stroke-width": "0.2",
		"fill":         "transparent",
	}.String())

	for i, vertex := range before {
		c := g1.CreateElement("circle")
		c.CreateAttr("id", fmt.Sprintf("vertex-before-%d", i))
		c.CreateAttr("transform", fmt.Sprintf("translate(%f,%f)", vertex.X, vertex.Y))
		c.CreateAttr("style", svg.StyleMap{
			"r":            "0.5",
			"stroke":       "black",
			"stroke-width": "0.1",
			"fill":         "blue",
		}.String())
	}

	g2 := root.CreateElement("g")
	g2.CreateAttr("id", "after")

	p2 := g1.CreateElement("path")
	p2.CreateAttr("id", "path-after")
	p2Data := make([]string, len(after)+1)
	for i := 0; i < len(after); i++ {
		if i == 0 {
			p2Data[i] = fmt.Sprintf("M %f,%f", after[i].X, after[i].Y)
		} else {
			p2Data[i] = fmt.Sprintf("L %f,%f", after[i].X, after[i].Y)
		}
	}
	p2Data[len(after)] = fmt.Sprintf("L %f,%f", after[0].X, after[0].Y)
	p2.CreateAttr("d", strings.Join(p2Data, " "))
	p2.CreateAttr("style", svg.StyleMap{
		"stroke":       "purple",
		"stroke-width": "0.2",
		"fill":         "transparent",
	}.String())

	for i, vertex := range after {
		c := g1.CreateElement("circle")
		c.CreateAttr("id", fmt.Sprintf("vertex-after-%d", i))
		c.CreateAttr("transform", fmt.Sprintf("translate(%f,%f)", vertex.X, vertex.Y))
		c.CreateAttr("style", svg.StyleMap{
			"r":            "0.5",
			"stroke":       "black",
			"stroke-width": "0.1",
			"fill":         "purple",
		}.String())
	}

	return WriteSVGToFile(doc, path)
}
