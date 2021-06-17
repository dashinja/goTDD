package starter

import (
	"fmt"

	"github.com/dashinja/learn-go-with-tests/learn/integers"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const mandarinHelloPrefix = "Ni hao, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	case "Mandarin":
		prefix = mandarinHelloPrefix
	
	default:
		prefix = englishHelloPrefix
	}

	return 
}

func main() {
	fmt.Println(Hello("world", ""))

	fmt.Printf("Answer: %d", integers.Add(2, 2))
	
}
