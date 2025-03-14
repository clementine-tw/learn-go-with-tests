package clockface

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"
)

// An awesome website can translate SVG for us:
// https://xml-to-go.github.io/
type Svg struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 60},
		},
		{
			simpleTime(0, 0, 30),
			Line{150, 150, 150, 240},
		},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, c.time)

			svg := Svg{}
			err := xml.Unmarshal(b.Bytes(), &svg)
			if err != nil {
				t.Fatalf("unmarshal xml failed: %v", err)
			}

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the second hand line %+v, "+
					"in the SVG lines %+v", c.line, svg)
			}
		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 70},
		},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, c.time)

			svg := Svg{}
			err := xml.Unmarshal(b.Bytes(), &svg)
			if err != nil {
				t.Fatalf("unmarshal xml failed: %v", err)
			}

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the minute hand line %+v, "+
					"in the SVG lines %+v", c.line, svg)
			}
		})
	}
}

func TestSVGWriterHourHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(6, 0, 0), Line{150, 150, 150, 210}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, c.time)

			svg := Svg{}
			if err := xml.Unmarshal(b.Bytes(), &svg); err != nil {
				t.Fatalf("Unmarshal xml file error: %v", err)
			}

			if !containsLine(c.line, svg.Line) {
				t.Fatalf("Expected to find hour hand line %+v, "+
					"in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func containsLine(l Line, ls []Line) bool {
	for _, line := range ls {
		if l == line {
			return true
		}
	}
	return false
}
