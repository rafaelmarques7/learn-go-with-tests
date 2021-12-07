package main

import "fmt"

const helloPrefix = "Hello, "
const helloSuffix = "!"

// domain function - this facilitates testing
func SayHello(name string) string {
	// when no name is available, say hello to the world
	if name == "" {
		name = "World"
	}
	return helloPrefix + name + helloSuffix
}

// function with side effects (printing to console) - not covered by tests
func main() {
	fmt.Println(SayHello(""))
}