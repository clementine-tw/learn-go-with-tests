package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (secondsInHalfClock / float64(t.Second()))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / minutesInClock) + (math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock) + (math.Pi / (hoursInHalfClock / float64(t.Hour()%hoursInClock)))
}

func secondHandPoint(t time.Time) Point {
	return radiansToPoint(secondsInRadians(t))
}

func minuteHandPoint(t time.Time) Point {
	return radiansToPoint(minutesInRadians(t))
}

func hourHandPoint(t time.Time) Point {
	return radiansToPoint(hoursInRadians(t))
}

func radiansToPoint(radians float64) Point {
	x := math.Sin(radians)
	y := math.Cos(radians)
	return Point{x, y}
}
