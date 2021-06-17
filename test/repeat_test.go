package test

import (
	"fmt"
	"testing"

	"github.com/dashinja/learn-go-with-tests/iteration"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeat 7 times", func(t *testing.T) {
		expected := "aaaaaaa"
		repeated := iteration.Repeat("a", 7)

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.Repeat("a", 7)
	}
}

func ExampleRepeat() {
	result := iteration.Repeat("b", 10)
	fmt.Println(result)
	//Output: bbbbbbbbbb
}