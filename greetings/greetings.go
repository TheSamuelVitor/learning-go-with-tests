package greetings

// Greets someone
func Hello(name, language string) string {
	if language == "es" {
		return "Hola, " + name
	} else if language == "fr" {
		return "Salut, " + name
	} else if language == "en" {
		return "Hello, " + name
	} else if language == "pt" {
		return "Ola, " + name
	}

	return "Idioma nao encontrado"
}
