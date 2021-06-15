package arrays

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("the Sum of elements in array param is 17", func(t *testing.T) {
		numbers := [6]int {1, 2, 3, 4, 5, 2}
		expected := 17
		sum := Sum(numbers)

		if expected != sum {
			t.Errorf("expected: %d, sum: %d, given: %v", expected, sum, numbers)
		}
	})
}
func TestSumAlt(t *testing.T) {
	t.Run("the Sum of elements in slice param is 17", func(t *testing.T) {
		numbers := []int {1, 2, 3, 4, 5, 2}
		expected := 17
		sum := SumAlt(numbers)

		if expected != sum {
			t.Errorf("expected: %d, sum: %d, given: %v", expected, sum, numbers)
		}
	})
}