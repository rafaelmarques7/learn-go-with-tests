package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy := &SpySleeper{}

		Countdown(buffer, spy)

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("", func(t *testing.T) {
		spy := &SpyCountdownOps{}

		Countdown(spy, spy)

		expectedCalls := []string{
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(spy.Calls, expectedCalls) {
			t.Errorf("Unexpected sequence, expected %v, got %v", spy.Calls, expectedCalls)
		}
	})
}

const write = "write"
const sleep = "sleep"

type SpyCountdownOps struct {
	Calls []string
}

func (s *SpyCountdownOps) Sleep(d time.Duration) {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOps) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpySleeper struct {
	NumberOfCalls int
}

func (s *SpySleeper) Sleep(d time.Duration) {
	s.NumberOfCalls++
}
