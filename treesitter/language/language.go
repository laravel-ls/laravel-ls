package language

import (
	"unsafe"

	ts "github.com/tree-sitter/go-tree-sitter"
)

// Language is a wrapper object that makes is easier to handle tree-sitter language objects.
type Language struct {
	id    Identifier
	tsObj *ts.Language
}

func NewLanguage(id Identifier, impl unsafe.Pointer) *Language {
	return &Language{
		id:    id,
		tsObj: ts.NewLanguage(impl),
	}
}

// ID get the identifier for this language
func (l Language) ID() Identifier {
	return l.id
}

// Name get the name of this language
func (l Language) Name() string {
	return l.id.String()
}

// TSObject get the tree-sitter object.
func (l Language) TSObject() *ts.Language {
	return l.tsObj
}

// Query create a new Query object for this language.
func (l Language) Query(source string) (*ts.Query, *ts.QueryError) {
	return ts.NewQuery(l.tsObj, source)
}
