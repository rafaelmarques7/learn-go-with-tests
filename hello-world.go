package main

import "fmt"

// domain function - this facilitates testing
func sayHello() string {
	return "Hello World!"
}

func sayHelloToAnyone(name string) string {
	return "Hello, " + name + "!"
}

// function with side effects (printing to console) - not covered by tests
func main() {
	fmt.Println(sayHello())
}