package env

import (
	"errors"

	"github.com/shufflingpixels/laravel-ls/parser"
	"github.com/shufflingpixels/laravel-ls/treesitter"
	ts "github.com/tree-sitter/go-tree-sitter"
)

type Metadata struct {
	Value  string
	Line   int
	Column int
}

type State struct {
	file *parser.File
	data map[string]Metadata
}

func (s *State) resolveVariable(node *ts.Node) string {
	id := treesitter.FirstNamedChildOfKind(node, "identifier")
	if id != nil {
		key := id.Utf8Text(s.file.Src)
		if meta, found := s.data[key]; found {
			return meta.Value
		}
	}
	// Could not resolve the variable.
	// So just return the text for the node.
	return node.Utf8Text(s.file.Src)
}

func (s *State) parseString(node *ts.Node) (string, error) {
	value := ""
	for ; node != nil; node = node.NextNamedSibling() {
		switch node.Kind() {
		case "variable":
			value += s.resolveVariable(node)
			break
		default:
			value += node.Utf8Text(s.file.Src)
			break
		}
	}
	return value, nil
}

func (s *State) parseValue(node *ts.Node) (string, error) {
	switch node.Kind() {
	case "string":
		return s.parseString(treesitter.FirstNamedChild(node))
	default:
		return node.Utf8Text(s.file.Src), nil
	}
}

func (s *State) parseAssignment(node *ts.Node) error {
	key := node.ChildByFieldName("key")
	value := node.ChildByFieldName("value")

	if key == nil {
		return errors.New("key expected")
	}

	stringValue := ""
	if value != nil {
		var err error
		stringValue, err = s.parseValue(value)
		if err != nil {
			return err
		}
	} else {
		value = key
	}

	keyName := key.Utf8Text(s.file.Src)
	point := value.Range().StartPoint
	s.data[keyName] = Metadata{
		Value:  stringValue,
		Line:   int(point.Row),
		Column: int(point.Column),
	}
	return nil
}

func (s *State) parse(node *ts.Node) error {
	for ; node != nil; node = node.NextNamedSibling() {
		if node.Kind() == "assignment" {
			err := s.parseAssignment(node)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func Parse(file *parser.File) (map[string]Metadata, error) {
	state := State{
		file: file,
		data: map[string]Metadata{},
	}

	root := file.Tree.Root()
	if root.Kind() == "document" {
		return state.data, state.parse(treesitter.FirstNamedChild(root))
	}
	return nil, errors.New("invalid document")
}
