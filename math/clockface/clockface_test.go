package clockface

import (
	"math"
	"testing"
	"time"
)

const (
	equalityThreshold = 1e-7
)

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := secondHandPoint(test.time)
			if !roughlyEqualPoint(got, test.point) {
				t.Errorf("got %v, want %v", got, test.point)
			}
		})
	}
}

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		t       time.Time
		radians float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 45), math.Pi * 3 / 2},
		{simpleTime(0, 0, 7), math.Pi * 7 / 30},
	}

	for _, test := range cases {
		t.Run(testName(test.t), func(t *testing.T) {
			got := secondsInRadians(test.t)
			if got != test.radians {
				t.Errorf("got %v, want %v", got, test.radians)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minuteHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("got %v, want %v", got, c.point)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		t       time.Time
		radians float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, test := range cases {
		t.Run(testName(test.t), func(t *testing.T) {
			got := minutesInRadians(test.t)
			if got != test.radians {
				t.Errorf("got %v, want %v", got, test.radians)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hourHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("got %v, want %v", got, c.point)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time    time.Time
		radians float64
	}{
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hoursInRadians(c.time)
			if !roughlyEqualFloat64(got, c.radians) {
				t.Errorf("got %v, want %v", got, c.radians)
			}
		})
	}
}

func simpleTime(hour, minute, second int) time.Time {
	return time.Date(2025, time.February, 1, hour, minute, second, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func roughlyEqualFloat64(a, b float64) bool {
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}
