package maps

import (
	"testing"
)

// Test betweem the language files extensions
func TestLinguagem(t *testing.T) {

	linguagens := Dictionary{
		"go": "Golang",
		"py": "Python",
		"js": "JavaScript",
		"ts": "TypeScript",
	}

	got, err := linguagens.Search("go")
	want := "Golang"

	if err != nil {
		t.Errorf("want %v\n, got %v", want, got)
	}

	if got != want {
		t.Errorf("want %q\n got %q", want, got)
	}
}
