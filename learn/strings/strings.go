package strings

import (
	"fmt"
	"strings"
)

// ReplaceCondition is a custum wrapper to produce a strings.ReplaceAll function
func ReplaceCondition(original string, oldSubString string, newSubString string) string {
	return strings.Replace(original, oldSubString, newSubString, -1)
}

func Do() string{

	var b strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("ignition")
	fmt.Println(b.String())	
	return b.String()
}