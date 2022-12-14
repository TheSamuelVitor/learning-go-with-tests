package main

import (
	"reflect"
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

}

func TestSumAll(t *testing.T) {
	t.Run("testing with multiple arrays", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0,9})
		want := []int{3,9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("\ngot %v\nwant %v", got, want)
		}
	})
}