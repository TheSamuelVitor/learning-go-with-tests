package main

import (
	"fmt"
	"testing"
)


func TestHello(t *testing.T) {
	lista := map[string]string{
		"Samuel": "en",
	}

	fmt.Println(lista)

	got := Hello("Kayo", "es")
	want := "Hola, Kayo"

	if got != want {
		t.Errorf("\ngot %v\nwant %v", got, want)
	}
}

func TestHello2(t *testing.T) {
	got := Hello2("Kayo", "es")
	want := "Hola, Kayo"

	if got != want {
		t.Errorf("\ngot %v\nwant %v", got,want)
	}
}