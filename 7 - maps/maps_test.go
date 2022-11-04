package maps

import "testing"

// Compares between the got and the want
func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q\nwant %q", got, want)
	}
}

// Tests the search function
func TestSearch(t *testing.T) {
	
	dictionary := Dictionary{
		"test":   "this is just a test",
		"golang": "programming language created by google",
	}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
	
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("teste")
		want := "could not find the word"
	
		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), want)
	})
}

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
		t.Errorf("language is not in dictionary")
	}

	assertStrings(t, got, want)
}

