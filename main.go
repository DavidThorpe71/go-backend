package main

import (
	"log"
	"net/http"
)

// InMemoryPlayerStore yadda yadda
type InMemoryPlayerStore struct{}

// GetPlayerScore yaddad
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerStore) RecordWin(name string) {}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("Could not listen on port 5000 %v", err)
	}
}
