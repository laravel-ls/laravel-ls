package parser

import (
	"testing"

	"github.com/laravel-ls/laravel-ls/treesitter/language"

	"github.com/stretchr/testify/assert"
	ts "github.com/tree-sitter/go-tree-sitter"
)

func TestLanguageTree_Parse(t *testing.T) {
	src := []byte(`
<div>
	<?php $var = 2; ?>
</div>`)

	tree, err := newLanguageTree(language.PHP.Language(), []ts.Range{})
	assert.NoError(t, err)
	assert.NoError(t, tree.parse(src))

	assert.Len(t, tree.childTrees, 1)

	assert.Equal(t, "(program (text) (php_tag) (expression_statement (assignment_expression left: (variable_name (name)) right: (integer))) (text_interpolation (text)))", tree.tree.RootNode().ToSexp())
	assert.Equal(t, language.PHP, tree.language.ID())
	assert.Equal(t, "(document (element (start_tag (tag_name)) (end_tag (tag_name))))", tree.childTrees[0].tree.RootNode().ToSexp())
}

func TestLanguageTree_UpdateThatRemovesInjectionRegion(t *testing.T) {
	src := []byte(`<div>
<?php $var = 2; ?>
</div>`)

	changedSrc := []byte(`<?php $var = 2; ?>`)

	tree, err := newLanguageTree(language.PHP.Language(), []ts.Range{})
	assert.NoError(t, err)
	assert.NoError(t, tree.parse(src))

	assert.Len(t, tree.childTrees, 1)

	assert.Equal(t, "(program (text) (php_tag) (expression_statement (assignment_expression left: (variable_name (name)) right: (integer))) (text_interpolation (text)))", tree.tree.RootNode().ToSexp())
	assert.Equal(t, language.PHP, tree.language.ID())
	assert.Equal(t, "(document (element (start_tag (tag_name)) (end_tag (tag_name))))", tree.childTrees[0].tree.RootNode().ToSexp())

	tree.update(&ts.InputEdit{
		StartByte:  0,
		OldEndByte: 5,
		NewEndByte: 0,
		StartPosition: ts.Point{
			Row:    0,
			Column: 0,
		},
		OldEndPosition: ts.Point{
			Row:    0,
			Column: 5,
		},
		NewEndPosition: ts.Point{
			Row:    0,
			Column: 0,
		},
	})

	tree.update(&ts.InputEdit{
		StartByte:  16,
		OldEndByte: 24,
		NewEndByte: 16,
		StartPosition: ts.Point{
			Row:    0,
			Column: 16,
		},
		OldEndPosition: ts.Point{
			Row:    1,
			Column: 6,
		},
		NewEndPosition: ts.Point{
			Row:    0,
			Column: 16,
		},
	})

	err = tree.parse(changedSrc)
	assert.NoError(t, err)

	assert.Len(t, tree.childTrees, 0)

	assert.Equal(t, "(program (php_tag) (expression_statement (assignment_expression left: (variable_name (name)) right: (integer))) (text_interpolation))", tree.tree.RootNode().ToSexp())
	assert.Equal(t, language.PHP, tree.language.ID())
}
