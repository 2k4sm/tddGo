package arrsli

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum of list of integers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("SumAll: empty slice passed.", func(t *testing.T) {
		slice := []int{}

		got := SumAll(slice)
		want := []int{0}

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v, given %v", got, want, slice)
		}
	})

	t.Run("SumAll: only one slice passed.", func(t *testing.T) {
		slice1 := []int{1, 2, 3, 4, 5}

		got := SumAll(slice1)
		want := []int{15}

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v, given %v", got, want, slice1)
		}
	})

	t.Run("SumAll: 5 slices passed.", func(t *testing.T) {
		slice1 := []int{1, 2, 3, 4, 5}
		slice2 := []int{5, 1, 3, 1, 5}
		slice3 := []int{1, 2, 4, 5}
		slice4 := []int{1, 3, 4, 5}
		slice5 := []int{1, 2, 1, 2}

		got := SumAll(slice1, slice2, slice3, slice4, slice5)
		want := []int{15, 15, 12, 13, 6}

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v, \n given:- \n%v\n%v\n%v\n%v\n%v", got, want, slice1, slice2, slice3, slice4, slice5)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("SumAllTails: empty slice passed.", func(t *testing.T) {
		slice := []int{}

		got := SumAllTails(slice)
		want := []int{0}

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v, given %v", got, want, slice)
		}
	})

	t.Run("SumAllTails: slice with 1 element passed.", func(t *testing.T) {
		slice := []int{3}

		got := SumAllTails(slice)
		want := []int{0}

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v, given %v", got, want, slice)
		}
	})

	t.Run("SumAllTails: only one slice passed.", func(t *testing.T) {
		slice1 := []int{1, 2, 3, 4, 5}

		got := SumAllTails(slice1)
		want := []int{14}

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v, given %v", got, want, slice1)
		}
	})

	t.Run("SumAllTails: 5 slices passed.", func(t *testing.T) {
		slice1 := []int{1, 2, 3, 4, 5}
		slice2 := []int{5, 1, 3, 1, 5}
		slice3 := []int{1, 2, 4, 5}
		slice4 := []int{1, 3, 4, 5}
		slice5 := []int{1, 2, 1, 2}

		got := SumAllTails(slice1, slice2, slice3, slice4, slice5)
		want := []int{14, 10, 11, 12, 5}

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v, \ngiven: \n%v\n%v\n%v\n%v\n%v", got, want, slice1, slice2, slice3, slice4, slice5)
		}
	})
}
