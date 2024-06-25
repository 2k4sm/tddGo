package hello

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	prefix := englishPrefix
	switch language {
	case "spanish":
		prefix = spanishPrefix
	case "english":
		prefix = englishPrefix
	case "french":
		prefix = frenchHelloPrefix
	default:
		prefix = englishPrefix
	}

	if name == "" {
		name = "World"
	}

	return prefix + name
}
