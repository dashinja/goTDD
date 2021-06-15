package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeat 7 times", func(t *testing.T) {
		expected := "aaaaaaa"
		repeated := Repeat("a", 7)

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}

func ExampleRepeat() {
	result := Repeat("b", 10)
	fmt.Println(result)
	//Output: bbbbbbbbbb
}