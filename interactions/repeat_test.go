package interactions

import "testing"

func GotWithExpected(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("\ngot %v\nwant %v", got, want)
	}
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	GotWithExpected(t, repeated, expected)
}