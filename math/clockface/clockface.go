package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const (
	secondHandLength = 90
	clockCentreX     = 150
	clockCentreY     = 150
)

func SecondHand(w io.Writer, t time.Time) {
	p := secondHandPoint(t)
	p.X = p.X*secondHandLength + clockCentreX
	p.Y = clockCentreY - p.Y*secondHandLength
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func secondInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

func secondHandPoint(t time.Time) Point {
	raid := secondInRadians(t)
	return Point{X: math.Sin(raid), Y: math.Cos(raid)}
}

func SVGWriter(w io.Writer, time time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	SecondHand(w, time)
	io.WriteString(w, svgEnd)

}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`
const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
const svgEnd = `</svg>`
