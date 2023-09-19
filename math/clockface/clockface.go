package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	return Point{}
}

func secondInRadians(t time.Time) float64 {
	return float64(t.Second()) *  math.Pi / 30 
}