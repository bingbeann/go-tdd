package hello

import "fmt"

var greetingPrefix = map[string]string{
	"English": "Hello",
	"Spanish": "Hola",
	"French":  "Bonjour",
}

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	if language == "" {
		language = "English"
	}

	prefix := greetingPrefix[language]

	return fmt.Sprintf("%s, %s", prefix, name)
}
