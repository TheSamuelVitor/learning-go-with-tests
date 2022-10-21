package main

import (
	"fmt"
)


// Greets someone 
func Hello(name, language string) string {
	if language == "es" {
		return "Hola, " + name
	} else if language == "fr" {
		return "Bonjour, " + name
	}

	return "Hello, " + name
}

func Hello2(name, language string) string {
	return fmt.Sprintf("%s, %s", greeting(language), name)
}

var greetings = map[string]string{
	"es": "Hola",
	"fr": "Bonjour",
	"pt": "Ola",
}

func greeting(language string) string {
	fmt.Scanln(&language)
	greeting, existis := greetings[language]

	if existis {
		return greeting
	}

	return "Hello"
}

func main() {
	fmt.Println(Hello2("Samuel", "en"))
}