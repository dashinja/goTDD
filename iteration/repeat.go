package iteration

func Repeat(char string, repeatCount int) string {
	var result string
	for i := 0; i < repeatCount; i++ {
		result += char
	}
	return result
}