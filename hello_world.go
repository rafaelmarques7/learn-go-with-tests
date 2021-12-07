package main

import "fmt"

const helloPrefixEnglish = "Hello, "
const helloPrefixSpanish = "Hola, "
const helloPrefixFrench = "Bonjour, "
const helloSuffix = "!"

// domain function - this facilitates testing
func SayHello(name string, language string) string {
	// when no name is available, say hello to the world
	if name == "" {
		name = "World"
	}

	// select a prefix based on the language
	helloPrefix := greetingPrefix(language)

	// build the hello string and return
	return helloPrefix + name + helloSuffix
}

// this is an internal function, so it is named with a lower case
func greetingPrefix(language string) (helloPrefix string) {
	// choose the prefix based on the selected language
	switch language {
	case "Spanish":
		helloPrefix = helloPrefixSpanish
	case "French":
		helloPrefix = helloPrefixFrench
	default:
		helloPrefix = helloPrefixEnglish
	}

	// the function signature is set to return `(helloPrefix string)`, 
	// so this is what is returned when we call `return` without any arguments
	return
}

// function with side effects (printing to console) - not covered by tests
func main() {
	fmt.Println(SayHello("", ""))
}