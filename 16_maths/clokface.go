package clock

import (
	// "fmt"
	"math"
	"time"
)

const secondHandLength = 90
const clockCentreX = 150
const clockCentreY = 150

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	// the x y position is determined by the cos and sin respectively
	degree := SecondsInRadians(t)

	delta_coeff_x := math.Cos(degree)
	delta_coeff_y := math.Sin(degree)

	x := clockCentreX + secondHandLength*delta_coeff_x
	y := clockCentreY + secondHandLength*delta_coeff_y

	return Point{x, y}
}

func SecondsInRadians(t time.Time) float64 {
	// The seconds hand rotates 1/60 of a complete rotation, that is (1/60) * 2 * PI (rad) every second.
	// We are going to measure the angle from: the OX axis, in clockwise rotation.
	// Thus, 15 seconds corresponds to 0 deg, 30 sec to 90 deg, 60 sec to 270 deg
	// given that the seconds hand starts at the top (OY axis), there is an offset of -90 degree

	offset := -math.Pi / 2.0
	degree := offset + float64(float64(1)/float64(60)*float64(t.Second()*2))*math.Pi

	return degree
}
