package language

// Identifier is a type that uniquely identifies a language
type Identifier string

// InvalidIdentifier special identifier that is used to flag
// that the language is invalid.
const InvalidIdentifier Identifier = "invalid"

// Language get the language object for this identifier
// Syntactic sugar for Get(id)
func (id Identifier) Language() *Language {
	return Get(id)
}

// String syntactic sugar for string(id)
func (id Identifier) String() string {
	return string(id)
}

func (id Identifier) Valid() bool {
	return id != InvalidIdentifier
}
