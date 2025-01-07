package parser

import (
	"fmt"

	"github.com/laravel-ls/laravel-ls/treesitter"
	"github.com/laravel-ls/laravel-ls/treesitter/debug"
	"github.com/laravel-ls/laravel-ls/treesitter/injections"

	"github.com/sirupsen/logrus"
	ts "github.com/tree-sitter/go-tree-sitter"
)

type LanguageTree struct {
	parser     *ts.Parser
	tree       *ts.Tree
	language   string
	ranges     []ts.Range
	childTrees []*LanguageTree
}

func newLanguageTree(language string, ranges []ts.Range) (*LanguageTree, error) {
	parser := ts.NewParser()
	err := parser.SetLanguage(treesitter.GetLanguage(language))
	if err != nil {
		return nil, err
	}

	return &LanguageTree{
		parser:     parser,
		ranges:     ranges,
		language:   language,
		childTrees: []*LanguageTree{},
	}, nil
}

func (t *LanguageTree) Root() *ts.Node {
	return t.tree.RootNode()
}

func (t *LanguageTree) InRange(r ts.Range) bool {
	for _, c := range t.ranges {
		if treesitter.RangeOverlap(c, r) {
			return true
		}
	}
	return false
}

func (t *LanguageTree) update(edit *ts.InputEdit) {
	t.tree.Edit(edit)

	// Update all children
	for _, child := range t.childTrees {
		child.update(edit)
	}
}

func (t *LanguageTree) parse(source []byte) error {
	// Parse main tree first.
	err := t.parser.SetIncludedRanges(t.ranges)
	if err != nil {
		return err
	}

	t.tree = t.parser.Parse(source, t.tree)

	// Then parse any injections
	return t.parseInjections(source)
}

func (t *LanguageTree) parseInjections(source []byte) error {
	injectionQuery := injections.GetQuery(t.language)

	if len(injectionQuery) < 1 {
		return nil
	}

	query, err := ts.NewQuery(t.tree.Language(), injectionQuery)
	if err != nil {
		return err
	}
	defer query.Close()

	childTrees := []*LanguageTree{}

	for _, injection := range injections.Query(query, t.tree.RootNode(), source) {
		if injection.Combined {
			if tree := findTreeForLanguage(injection.Language, childTrees); tree != nil {
				tree.ranges = append(tree.ranges, injection.Range)
				continue
			}
		}

		tree, err := newLanguageTree(injection.Language, []ts.Range{injection.Range})
		if err != nil {
			return err
		}

		childTrees = append(childTrees, tree)
	}

	// Parse all children
	t.childTrees = childTrees
	for _, child := range t.childTrees {
		if err := child.parse(source); err != nil {
			return err
		}
	}

	return nil
}

// Get all trees for a particular language
func (t LanguageTree) GetLanguageTrees(language string) []*LanguageTree {
	results := []*LanguageTree{}

	if t.language == language {
		results = append(results, &t)
	}

	for _, tree := range t.childTrees {
		results = append(results, tree.GetLanguageTrees(language)...)
	}
	return results
}

func (t LanguageTree) FindCaptures(language, pattern string, source []byte, captures ...string) (treesitter.CaptureSlice, error) {
	query, err := ts.NewQuery(treesitter.GetLanguage(language), pattern)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	// Build a map of index name pairs.
	captureMap := map[uint]string{}

	for _, capture := range captures {
		index, ok := query.CaptureIndexForName(capture)
		if !ok {
			return nil, fmt.Errorf("Capture '%s' is not present in query", capture)
		}

		captureMap[index] = capture
	}

	cursor := ts.NewQueryCursor()
	defer cursor.Close()

	results := []treesitter.Capture{}
	for _, tree := range t.GetLanguageTrees(language) {
		matches := cursor.Matches(query, tree.Root(), source)
		for it := matches.Next(); it != nil; it = matches.Next() {
			for _, capture := range it.Captures {
				name, ok := captureMap[uint(capture.Index)]
				if !ok {
					continue
				}

				results = append(results, treesitter.Capture{
					Name: name,
					Node: capture.Node,
				})
			}
		}
	}

	return results, nil
}

// Find all trees of a given language that includes the node
func (t *LanguageTree) GetLanguageTreesWithNode(language string, node *ts.Node) []*LanguageTree {
	results := []*LanguageTree{}

	logrus.Debug(t.language, t.Root().Range(), node.Range())

	if t.language == language && treesitter.RangeOverlap(t.Root().Range(), node.Range()) {
		results = append(results, t)
	}

	for _, tree := range t.childTrees {
		results = append(results, tree.GetLanguageTreesWithNode(language, node)...)
	}
	return results
}

func findTreeForLanguage(language string, trees []*LanguageTree) *LanguageTree {
	for _, tree := range trees {
		if tree.language == language {
			return tree
		}
	}
	return nil
}

func VisualizeLanguageTree(tree *LanguageTree) string {
	printer := debug.Print{
		IndentSize: 2,
	}

	return visualizeLanguageTreeInner(printer, tree, nil, 0)
}

func visualizeLanguageTreeInner(printer debug.Print, tree *LanguageTree, parent *ts.Node, depth uint) string {
	str := printer.Indent(depth) + "<" + tree.language
	if parent != nil {
		str += " " + debug.FormatNode(parent)
	}
	str += ">\n"

	str += printer.Indent(depth) + debug.FormatNode(tree.tree.RootNode()) + "\n"

	for i := uint(0); i < tree.tree.RootNode().NamedChildCount(); i++ {
		node := tree.tree.RootNode().NamedChild(i)
		if childStr, ok := visualizeChildTrees(printer, node, tree.childTrees, depth); ok {
			str += childStr
		} else {
			str += printer.Inner(node, depth+1)
		}
	}

	str += printer.Indent(depth) + "</" + tree.language + ">\n"

	return str
}

func visualizeChildTrees(printer debug.Print, node *ts.Node, trees []*LanguageTree, depth uint) (string, bool) {
	for _, childTree := range trees {
		if childTree.InRange(node.Range()) {
			return visualizeLanguageTreeInner(printer, childTree, node, depth+1), true
		}
	}
	return "", false
}
