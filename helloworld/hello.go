package main

import "fmt"

func main() {
	SayHello("Phil", "")
}

const (
	french  = "French"
	spanish = "Spanish"

	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
)

func SayHello(word, language string) string {
	if word == "" {
		return fmt.Sprint("Hello, world")
	}
	return greetingPrefix(language) + word
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix

	}
	return
}
