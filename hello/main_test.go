package hello

import (
	"testing"
)

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("\ngot %v\nwant %v", got, want)
	}
}

func TestHelloLanguages(t *testing.T) {
	nome := "Samuel"

	t.Run("saying Hello in english", func(t *testing.T) {
		got := Hello(nome, "en")
		want := "Hello, " + nome
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hello in spanish", func(t *testing.T) {
		got := Hello(nome, "es")
		want := "Hola, " + nome
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hello in portuguese", func(t *testing.T) {
		got := Hello(nome, "pt")
		want := "Ola, " + nome
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hello in french", func(t *testing.T) {
		got := Hello(nome, "fr")
		want := "Salut, " + nome
		assertCorrectMessage(t, got, want)
	})

	t.Run("not specifying the language", func(t *testing.T) {
		got := Hello(nome, "")
		want := "Hello, " + nome
		assertCorrectMessage(t, got, want)
	})

	t.Run("not specifying the name", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
