package test

import (
	"fmt"
	"testing"

	"github.com/dashinja/learn-go-with-tests/learn/strings"
)

func TestReplaceCondition(t *testing.T) {
	t.Run("replace Byron with Byronicus", func(t *testing.T) {
		expected := "Byronicus"
		output := strings.ReplaceCondition("Byron", "on", "onicus")

		if expected != output {
			t.Errorf("expected: %q, output: %q", expected, output)
		}
	})
}

func TestDo(t *testing.T) {
	expected := "3...2...1...ignition"
	output := strings.Do()

	if expected != output {
		t.Errorf("expected: %q, output: %q", expected, output)
	}
}

func ExampleReplaceCondition() {
	result := strings.ReplaceCondition("Hakuna Mufasa", "Mufasa", "Matata")

	fmt.Println(result)
	//Output: Hakuna Matata
}

func BenchmarkReplaceCondition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.ReplaceCondition("How many licks does it take to get to the center of a tootsie pop?", "licks", "kicks")
	}
}
