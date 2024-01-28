package main

const (
	englishPrefix = "Hello, "
	spanish       = "Spanish"
	spanishPrefix = "Hola, "
	french        = "French"
	frenchPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchPrefix
	case "Spanish":
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}
