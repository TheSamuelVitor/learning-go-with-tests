package main

import (
	"testing"
)

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestHelloPasses(t *testing.T) {
	const englishHelloPrefix = "Hello, "
	const name = "Samuel"

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello(name)
		want := englishHelloPrefix + name
		assertCorrectMessage(t, got, want)
	})

	t.Run("see what happens when name is empty", func(t *testing.T) {
		got := Hello("")
		want := englishHelloPrefix + "World"
		assertCorrectMessage(t, got, want)
	})
}
