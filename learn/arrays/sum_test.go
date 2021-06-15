package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("the Sum of elements in array param is 17", func(t *testing.T) {
		numbers := [6]int{1, 2, 3, 4, 5, 2}
		expected := 17
		sum := Sum(numbers)

		if expected != sum {
			t.Errorf("expected: %d, sum: %d, given: %v", expected, sum, numbers)
		}
	})
}
func TestSumAlt(t *testing.T) {
	t.Run("the Sum of elements in slice param is 17", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 2}
		expected := 17
		sum := SumAlt(numbers)

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
		output := SumAll([][]int{sliceOne, sliceTwo, sliceThree})

		// fmt.Println()
		// fmt.Printf("SumAll: %v", output)

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
		output := SumAll([][]int{sliceOne, sliceTwo, sliceThree})

		// fmt.Println()
		// fmt.Printf("SumAllAlt: %v", output)

		if !reflect.DeepEqual(expected, output) {
			t.Errorf("expected: %v, output: %v", expected, output)
		}

		fmt.Println()

	})
}



func ExampleSumAll() {
	sliceOne := []int{1, 2, 2}
	sliceTwo := []int{4, 6, 8, 6}
	sliceThree := []int{10, 11, 13}

	output := SumAll([][]int{sliceOne, sliceTwo, sliceThree})
	fmt.Println()
	fmt.Println(output)
	// Output: [5 24 34]
}
