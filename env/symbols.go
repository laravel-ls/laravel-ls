package env

type Variable struct {
	Value  string
	Line   uint32
	Column uint32
}

type SymbolTable map[string]Variable
