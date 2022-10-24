package main

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

	t.Run("saying Hello in enlish", func(t *testing.T) {
		got := Hello("Samuel", "en")
		want := "Hello, " + nome
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hello in spanish", func(t *testing.T) {
		got := Hello("Samuel", "es")
		want := "Hola, " + nome
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hello in portuguese", func(t *testing.T) {
		got := Hello("Samuel", "pt")
		want := "Ola, " + nome
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hello in french", func(t *testing.T) {
		got := Hello("Samuel", "fr")
		want := "Salut, " + nome
		assertCorrectMessage(t, got, want)
	})
}
