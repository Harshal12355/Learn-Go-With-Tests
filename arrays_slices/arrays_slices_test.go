package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum of array with 5 values", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := SumArray(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("sum of slice with 5 values", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		got := SumSlice(nums)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, nums)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum of slice with 5 values", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		// Go does not allow you to use equality operaters to compare two slices, so 
		// instead you can you reflect.DeepEqual() to compare *ANY TWO VARIABLES* and check if they are the same 
		// reflect.DeepEqual() is not type safe
		if !reflect.DeepEqual(got, want) { 
			t.Errorf("got %v want %v", got, want)
		}
	})

}

func TestSumAllTails(t *testing.T) {

	// there was repeating code, so a variable that acts as a 
	// function was made inside the function instead
	// of creating a whole new function to accomodate 
	// for this and then it was used.
	checkSums := func(t testing.TB, got, want []int) { // assigned a function to a vairable 
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}

