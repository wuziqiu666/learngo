package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func secondInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondInRadians(t))
}

func minuteInRadians(t time.Time) float64 {
	return (secondInRadians(t) / 60) + math.Pi/(30/float64(t.Minute()))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minuteInRadians(t))
}

func angleToPoint(angle float64) Point {
	return Point{X: math.Sin(angle), Y: math.Cos(angle)}
}
