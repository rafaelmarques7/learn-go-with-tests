package clockface

import (
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := Point{X: 150, Y: 150 - 90}
	got := SecondHand(tm)

	if got != want {
		t.Errorf(" got %v, want %v", got, want)
	}
}

func TestSecondHandAt15Second(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 15, 0, time.UTC)

	want := Point{X: 150 + 90, Y: 150}
	got := SecondHand(tm)

	if got != want {
		t.Errorf(" got %v, want %v", got, want)
	}
}

func TestSecondHandAt30Second(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := Point{X: 150, Y: 150 + 90}
	got := SecondHand(tm)

	if got != want {
		t.Errorf(" got %v, want %v", got, want)
	}
}
