package env

type Variable struct {
	Value  string
	Line   int
	Column int
}

type SymbolTable map[string]Variable
