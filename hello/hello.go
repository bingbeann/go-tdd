package main

import "fmt"

func main() {
	fmt.Println(Hello("world", ""))
}

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "

	languageSpanish = "Spanish"
	languageFrench  = "French"
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := getGreetingPrefix(language)

	return prefix + name
}

func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case languageSpanish:
		prefix = spanishHelloPrefix
	case languageFrench:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
