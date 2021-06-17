package test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/dashinja/learn-go-with-tests/arrays"
)

func TestSum(t *testing.T) {
	t.Run("the Sum of elements in array param is 17", func(t *testing.T) {
		numbers := [6]int{1, 2, 3, 4, 5, 2}
		expected := 17
		sum := arrays.Sum(numbers)

		if expected != sum {
			t.Errorf("expected: %d, sum: %d, given: %v", expected, sum, numbers)
		}
	})
}
func TestSumAlt(t *testing.T) {
	t.Run("the Sum of elements in slice param is 17", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 2}
		expected := 17
		sum := arrays.SumAlt(numbers)

		if expected != sum {
			t.Errorf("expected: %d, sum: %d, given: %v", expected, sum, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("SumAll returns new slice containing totals for each slice passed in", func(t *testing.T) {
		sliceOne := []int{1, 2, 2}
		sliceTwo := []int{4, 6, 8, 6}
		sliceThree := []int{10, 11, 13}

		expected := []int{5, 24, 34}
		output := arrays.SumAll([][]int{sliceOne, sliceTwo, sliceThree})

		if !reflect.DeepEqual(expected, output) {
			t.Errorf("expected: %v, output: %v", expected, output)
		}
	})
}
func TestSumAllAlt(t *testing.T) {
	t.Run("SumAllAlt returns new slice containing totals for each slice passed in", func(t *testing.T) {
		sliceOne := []int{1, 2, 2}
		sliceTwo := []int{4, 6, 8, 6}
		sliceThree := []int{10, 11, 13}

		expected := []int{5, 24, 34}
		output := arrays.SumAllAlt(sliceOne, sliceTwo, sliceThree)

		if !reflect.DeepEqual(expected, output) {
			t.Errorf("expected: %v, output: %v", expected, output)
		}

		fmt.Println()

	})
}
func TestSumAllTails(t *testing.T) {
	t.Run("SumAll returns new slice containing totals for each slice passed in", func(t *testing.T) {
		sliceOne := []int{1, 2, 2}
		sliceTwo := []int{4, 6, 8, 6}
		sliceThree := []int{10, 11, 13}

		expected := []int{4, 20, 24}
		output, _ := arrays.SumAllTails(sliceOne, sliceTwo, sliceThree)

		if !reflect.DeepEqual(expected, output) {
			t.Errorf("expected: %v, output: %v", expected, output)
		}
	})

	t.Run("Passing an empty slice results in instructional error", func(t *testing.T) {

		expected := errors.New("must not pass an empty slice")
		_, err := arrays.SumAllTails([]int{})

		if expected.Error() != err.Error() {
			t.Errorf("expected error output: %q\n actual error output: %q\n", expected, err)
		}
	})
}

func ExampleSumAll() {
	sliceOne := []int{1, 2, 2}
	sliceTwo := []int{4, 6, 8, 6}
	sliceThree := []int{10, 11, 13}

	output := arrays.SumAll([][]int{sliceOne, sliceTwo, sliceThree})
	fmt.Println()
	fmt.Println(output)
	// Output: [5 24 34]
}

func BenchmarkSumAllTails(b *testing.B) {
	sliceOne := []int{1, 2, 2}
	sliceTwo := []int{4, 6, 8, 6}
	sliceThree := []int{10, 11, 13}

	for i := 0; i < b.N; i++ {
		arrays.SumAll([][]int{sliceOne, sliceTwo, sliceThree})
	}
}
