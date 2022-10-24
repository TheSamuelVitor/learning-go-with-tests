package main

import "fmt"

// Sauda as pessoas recebendo o nome e a lingua
func Hello(name, language string) (saudacao string) {
	if name == "" {
		name = "World"
	}
	
	switch language {
	case "es":
		saudacao = "Hola, "
	case "pt":
		saudacao = "Ola, "
	case "fr":
		saudacao = "Salut, "
	default: 
		saudacao = "Hello, "
	}

	saudacao = saudacao + name
	return
}

func main() {
	fmt.Println(Hello("", "pt"))
}
