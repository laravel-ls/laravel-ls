package parser

import (
	"github.com/laravel-ls/laravel-ls/file"
	"github.com/laravel-ls/laravel-ls/treesitter"
	"github.com/laravel-ls/laravel-ls/utils"

	ts "github.com/tree-sitter/go-tree-sitter"
)

type File struct {
	Type file.Type

	Tree *LanguageTree

	// Content
	Src utils.Buffer
}

func Parse(content []byte, typ file.Type) (*File, error) {
	langTree, err := newLanguageTree(treesitter.FiletypeToLanguage(typ), []ts.Range{})
	if err != nil {
		return nil, err
	}

	file := &File{
		Type: typ,
		Tree: langTree,
	}
	return file, file.SetContent(content)
}

// Set the file contents.
func (f *File) SetContent(src []byte) error {
	f.Src = src
	return f.Tree.parse(f.Src)
}

// Update part of file contents
func (f *File) Update(start, end ts.Point, src []byte) error {
	edit := treesitter.CalculateEdit(start, end, f.Src, src)

	f.Tree.update(edit)

	f.Src.Update(edit.StartByte, edit.OldEndByte, src)

	return f.Tree.parse(f.Src)
}

// Get the node at a given position
// func (f *File) GetNodeAt(point ts.Point) *ts.Node {
// 	// return GetNodeAt(f.Tree.RootNode(), point)
// 	return GetNodeAt(f.Tree.tree.RootNode(), point)
// }

func (f *File) Query(pattern string) (*ts.Query, *ts.QueryError) {
	return ts.NewQuery(f.Tree.tree.Language(), pattern)
}

func (f *File) FindCaptures(language, pattern string, captures ...string) (CaptureSlice, error) {
	return f.Tree.FindCaptures(language, pattern, f.Src, captures...)
}

func (f *File) NodeMatchesCapture(language, pattern, capture string, node *ts.Node) bool {
	captures, err := f.FindCaptures(language, pattern, capture)
	if err != nil {
		return false
	}

	for _, found := range captures {
		if found.Node.Id() == node.Id() {
			return true
		}
	}
	return false
}
