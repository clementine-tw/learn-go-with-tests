package clockface

import (
	"fmt"
	"io"
	"time"
)

const (
	clockCentreX     = 150
	clockCentreY     = 150
	hourHandLength   = 60
	minuteHandLength = 80
	secondHandLength = 90

	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

func SVGWriter(w io.Writer, t time.Time) {
	_, _ = io.WriteString(w, svgStart)
	_, _ = io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	_, _ = io.WriteString(w, svgEnd)
}

// SecondHand is the unit vector of the second hand of an analogue clock at
// time `t`, represented as a Point
func secondHand(w io.Writer, t time.Time) {
	p := makeHand(secondHandPoint(t), secondHandLength)
	fmt.Fprintf(w, `<line x1="%d" y1="%d" x2="%.3f" y2="%.3f" `+
		`style="fill:none;stroke:#f00;stroke-width:3px;"/>`, clockCentreX, clockCentreY, p.X, p.Y)
}

func minuteHand(w io.Writer, t time.Time) {
	p := makeHand(minuteHandPoint(t), minuteHandLength)
	fmt.Fprintf(w, `<line x1="%d" y1="%d" x2="%.3f" y2="%.3f" `+
		`style="fill:none;stroke:#000;stroke-width:3px"/>`, clockCentreX, clockCentreY, p.X, p.Y)
}

func hourHand(w io.Writer, t time.Time) {
	p := makeHand(hourHandPoint(t), hourHandLength)
	fmt.Fprintf(w, `<line x1="%d" y1="%d" x2="%.3f" y2="%.3f" `+
		`style="fill:none;stroke:#000;stroke-width:3px"/>`, clockCentreX, clockCentreY, p.X, p.Y)
}

func makeHand(p Point, handLength float64) Point {
	p = Point{p.X * handLength, p.Y * handLength}
	p = Point{p.X, -p.Y}
	return Point{p.X + clockCentreX, p.Y + clockCentreY}
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
