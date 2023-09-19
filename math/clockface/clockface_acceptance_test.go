package clockface_test

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	"github.com/wuziqiu666/learngo/math/clockface"
)

type SVG struct {
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

func TestSVGWriterAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
	b := bytes.Buffer{}
	clockface.SVGWriter(&b, tm)
	svg := SVG{}
	xml.Unmarshal(b.Bytes(), &svg)
	want := Line{150, 150, 150, 60}
	for _, line := range svg.Line {
		//TODO: 继续写
		if line == want {
			return
		}
	}
	t.Errorf("Expect to find the second hand with x2 of %+v, and y2 of %+v, in the SVG output to %v", x2, y2, b.String())
}
