package main

import (
	"testing"
)

func ComparisonBetweenExpectedAndGot(t testing.TB, want, got int) {
	t.Helper()
	if got != want {
		t.Errorf("\ngot %v\nwant %v", got, want)
	}
}

func TestSum(t *testing.T) {

	t.Run("testing with arrays with 5 elements", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		ComparisonBetweenExpectedAndGot(t, got, want)
	})

	t.Run("testing with arrays with 3 elements", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		ComparisonBetweenExpectedAndGot(t, got, want)
	})

}
