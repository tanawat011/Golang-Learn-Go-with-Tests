package Hello

const thai = "thai"
const french = "french"

const englishHelloPrefix = "Hello, "
const thaiHelloPrefix = "สวัสดี, "
const frenchHelloPrefix = "Salut, "

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case thai:
		prefix = thaiHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
