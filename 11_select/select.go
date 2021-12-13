package main

import (
	"fmt"
	"net/http"
	"time"
)

// Given the URL for two servers, returns the URL of the server who responds first
func RacerZero(url_server_one, url_server_two string) string {
	// Note: this is the first implementation of the racer function
	aDuration := measureResponseTime(url_server_one)
	bDuration := measureResponseTime(url_server_two)

	if aDuration < bDuration {
		return url_server_one
	}
	return url_server_two
}

func measureResponseTime(url string) time.Duration {
	st := time.Now()
	http.Get(url)
	return time.Since(st)
}

// Given the URL for two servers, returns the URL of the server who responds first.
// How does it work?
// 		Take two urls and call the ping function on each;
// 		the ping call which returns first will then be the one whose case code will be executed
// Note:
// 		the select function allows you to wait between multiple channels,
// 		and execute the code of the one which responds first
func Racer(url_one, url_two string, timeout time.Duration) (string, error) {
	select {
	case <-ping(url_one):
		return url_one, nil
	case <-ping(url_two):
		return url_two, nil
	// time.After will return a channel after some delay, allowing us to timeout
	case <-time.After(timeout):
		return "", fmt.Errorf("timeoud out waiting for %q and %q", url_one, url_two)
	}
}

func ping(url string) chan struct{} {
	channel := make(chan struct{})
	go func() {
		http.Get(url)
		close(channel)
	}()
	return channel
}
