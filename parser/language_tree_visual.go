package parser

import (
	"github.com/laravel-ls/laravel-ls/treesitter/debug"
	ts "github.com/tree-sitter/go-tree-sitter"
)

func (langTree *LanguageTree) Visualize() string {
	printer := debug.Print{
		IndentSize: 2,
	}

	return langTree.visualizeLanguageTreeInner(printer, nil, 0)
}

func (langTree *LanguageTree) visualizeLanguageTreeInner(printer debug.Print, parent *ts.Node, depth uint) string {
	str := printer.Indent(depth) + "<" + langTree.language.Name()
	if parent != nil {
		str += " " + debug.FormatNode(parent)
	}
	str += ">\n"

	str += printer.Indent(depth) + debug.FormatNode(langTree.tree.RootNode()) + "\n"

	for i := uint(0); i < langTree.tree.RootNode().NamedChildCount(); i++ {
		node := langTree.tree.RootNode().NamedChild(i)
		if childStr, ok := langTree.visualizeChildTrees(printer, node, depth); ok {
			str += childStr
		} else {
			str += printer.Inner(node, depth+1)
		}
	}

	str += printer.Indent(depth) + "</" + langTree.language.Name() + ">\n"

	return str
}

func (langTree *LanguageTree) visualizeChildTrees(printer debug.Print, node *ts.Node, depth uint) (string, bool) {
	for _, childTree := range langTree.childTrees {
		if childTree.InRange(node.Range()) {
			return childTree.visualizeLanguageTreeInner(printer, node, depth+1), true
		}
	}
	return "", false
}
