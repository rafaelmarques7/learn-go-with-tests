package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep(d time.Duration)
}

type SleeperDefault struct{}

func (d *SleeperDefault) Sleep(duration time.Duration) {
	time.Sleep(duration)
}

func Countdown(buffer io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep(1 * time.Second)
		fmt.Fprintln(buffer, i)
	}
	sleeper.Sleep(1 * time.Second)
	fmt.Fprintf(buffer, "Go!")
}

func main() {
	sleeper := &SleeperDefault{}
	Countdown(os.Stdout, sleeper)
}
