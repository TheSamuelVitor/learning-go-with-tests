package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{}
	handler := http.HandlerFunc(server.ServerHTTP)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
