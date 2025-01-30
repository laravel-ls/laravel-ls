package parser

import (
	"github.com/laravel-ls/laravel-ls/file"
	"github.com/laravel-ls/laravel-ls/treesitter"
	"github.com/laravel-ls/laravel-ls/treesitter/language"
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
	lang := language.Get(language.FiletypeToLanguage(typ))

	langTree, err := newLanguageTree(lang, []ts.Range{})
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

func (f *File) FindCaptures(lang language.Identifier, query *ts.Query, captures ...string) (treesitter.CaptureSlice, error) {
	return f.Tree.FindCaptures(lang, query, f.Src, captures...)
}

func (f *File) NodeMatchesCapture(lang language.Identifier, query *ts.Query, capture string, node *ts.Node) bool {
	captures, err := f.FindCaptures(lang, query, capture)
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
