package test

import (
	"fmt"
	"testing"

	"github.com/dashinja/learn-go-with-tests/learn/integers"
)
func ExampleAdd() {
	sum := integers.Add(1, 5)
	fmt.Println(sum)
	//Output: 6
}

func TestAdd(t *testing.T)  {
	t.Helper()
	assertCorrectMessage := func(t testing.TB, expected, sum int) {
		if (expected != sum) {
			t.Errorf("expected %d, sum %d", expected, sum)
		}
	};
	
	
	t.Run("adds 2+5 appropriately", func(t *testing.T) {
		got := integers.Add(2, 5)
		want := 7
		assertCorrectMessage(t, got, want)
	})
}