package main

import "fmt"

func main() {
	fmt.Println(Hello("world", ""))
}

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "

	languageSpanish = "Spanish"
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	var prefix string
	switch language {
	case languageSpanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix + name
}
