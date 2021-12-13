package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("runs safely using linear programming", func(t *testing.T) {
		counter := NewCounter()

		counter.Increment()
		counter.Increment()
		counter.Increment()

		expected := 3

		assertCounter(t, counter, expected)
	})

	t.Run("runs safely using concurrent programming", func(t *testing.T) {
		counter := NewCounter()
		wantedCount := 1000

		// the WaitGroup, by using with the .Wait() method,
		// will wait for a number of goroutines to finish.
		// The number of goroutines to wait for is specified by the .Add() method,
		// each goroutine should call the .Done() method to singla it is finished.
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Increment()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, counter *Counter, want int) {
	t.Helper()
	if counter.Value() != want {
		t.Errorf("got %d, want %d", counter.Value(), want)
	}
}

func NewCounter() *Counter {
	return &Counter{}
}
