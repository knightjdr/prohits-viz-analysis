package circheatmap

import (
	"fmt"
	"math"

	"github.com/knightjdr/prohits-viz-analysis/pkg/color"
	"github.com/knightjdr/prohits-viz-analysis/pkg/float"
	customMath "github.com/knightjdr/prohits-viz-analysis/pkg/math"
)

// Circle has properties for drawing one circle in a circular heatmap.
type Circle struct {
	Attribute string
	Color     string
	Max       float64
	Min       float64
	Radius    float64
	Thickness float64
	Values    []float64
}

// Segment has properties for a single  circle segment
type Segment struct {
	A    SegmentPath
	B    SegmentPath
	C    SegmentPath
	D    SegmentPath
	Fill string
}

// SegmentPath contains x and y coordinates for drawing the path
type SegmentPath struct {
	Arc int
	X   float64
	Y   float64
}

func writeCircles(c *CircHeatmapSVG, writeString func(string)) {
	reformatted := reformatCircHeatmapData(c)

	for _, circle := range reformatted {
		writeCircle(circle, writeString)
	}
}

func reformatCircHeatmapData(c *CircHeatmapSVG) []Circle {
	reformatted := make([]Circle, len(c.Legend))

	space := c.Dimensions.Thickness / 4
	for i, legendItem := range c.Legend {
		attribute := legendItem.Attribute
		reformatted[i] = Circle{
			Attribute: attribute,
			Color:     legendItem.Color,
			Max:       legendItem.Max,
			Min:       legendItem.Min,
			Radius:    c.Dimensions.Radius - (float64(i) * (c.Dimensions.Thickness + space)),
			Thickness: c.Dimensions.Thickness,
			Values:    make([]float64, len(c.Plot.Readouts)),
		}

		for j, readout := range c.Plot.Readouts {
			reformatted[i].Values[j] = float64(readout.Segments[attribute])
		}
	}

	return reformatted
}

func writeCircle(c Circle, writeString func(string)) {
	radii := calculateRadii(c.Radius, c.Thickness)
	colors := createColourRange(c)

	segments := defineSegments(colors, radii)

	for _, segment := range segments {
		drawSegment(segment, radii, writeString)
	}
}

func createGradient(gradientColor string) []color.Space {
	gradient := color.InitializeGradient()
	gradient.ColorSpace = gradientColor
	gradient.NumColors = 101

	return gradient.CreateColorGradient()
}

func calculateRadii(radius, thickness float64) map[string]float64 {
	return map[string]float64{
		"inner": math.Floor(radius - thickness),
		"outer": radius,
	}
}

func createColourRange(c Circle) []string {
	colorGradient := createGradient(c.Color)
	convertValueToColorIndex := float.GetRange(c.Min, c.Max, 0, 100)

	colors := make([]string, len(c.Values))
	for i, value := range c.Values {
		index := int(convertValueToColorIndex(value))
		colors[i] = colorGradient[index].Hex
	}

	return colors
}

func defineSegments(colors []string, radii map[string]float64) []Segment {
	numSegments := len(colors)
	segments := make([]Segment, numSegments)

	var cumulativePercent float64
	last := map[string][]float64{
		"inner": {radii["inner"], 0},
		"outer": {radii["outer"], 0},
	}
	arc := 0
	if numSegments < 2 {
		arc = 1
	}
	percent := customMath.Round(1/float64(numSegments), 0.000001)
	for i, color := range colors {
		cumulativePercent += percent
		innerPoint := percentToCoordinate(cumulativePercent, radii["inner"])
		outerPoint := percentToCoordinate(cumulativePercent, radii["outer"])
		start := map[string][]float64{
			"inner": last["inner"],
			"outer": last["outer"],
		}
		last["inner"] = innerPoint
		last["outer"] = outerPoint
		segments[i] = Segment{
			A: SegmentPath{
				X: start["outer"][0],
				Y: start["outer"][1],
			},
			B: SegmentPath{
				Arc: arc,
				X:   outerPoint[0],
				Y:   outerPoint[1],
			},
			C: SegmentPath{
				X: innerPoint[0],
				Y: innerPoint[1],
			},
			D: SegmentPath{
				Arc: arc,
				X:   start["inner"][0],
				Y:   start["inner"][1],
			},
			Fill: color,
		}
	}

	return segments
}

func drawSegment(segment Segment, radii map[string]float64, writeString func(string)) {
	path := fmt.Sprintf(
		"M %f %f A %f %f 0 %d 1 %f %f L %f %f A %f %f 0 %d 0 %f %f Z",
		segment.A.X,
		segment.A.Y,
		radii["outer"],
		radii["outer"],
		segment.B.Arc,
		segment.B.X,
		segment.B.Y,
		segment.C.X,
		segment.C.Y,
		radii["inner"],
		radii["inner"],
		segment.D.Arc,
		segment.D.X,
		segment.D.Y,
	)

	writeString("\t\t<g transform=\"scale(0.85)\">\n")
	writeString(fmt.Sprintf(
		"\t\t\t<path d=\"%s\" fill=\"%s\" stroke=\"#f5f5f5\" strokeLinejoin=\"round\" strokeWidth=\"2\"/>\n",
		path,
		segment.Fill,
	))
	writeString("\t\t</g>\n")
}
