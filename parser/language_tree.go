package parser

import (
	"fmt"

	"github.com/laravel-ls/laravel-ls/treesitter"
	"github.com/laravel-ls/laravel-ls/treesitter/debug"
	"github.com/laravel-ls/laravel-ls/treesitter/injections"
	"github.com/laravel-ls/laravel-ls/treesitter/language"

	ts "github.com/tree-sitter/go-tree-sitter"
)

type LanguageTree struct {
	parser     *ts.Parser
	tree       *ts.Tree
	language   *language.Language
	ranges     []ts.Range
	childTrees []*LanguageTree
}

func newLanguageTree(language *language.Language, ranges []ts.Range) (*LanguageTree, error) {
	parser := ts.NewParser()
	err := parser.SetLanguage(language.TSObject())
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
	query, err := treesitter.GetInjectionQuery(t.language.ID())
	if err != nil {
		if err == treesitter.ErrQueryNotFound {
			return nil
		}
		return err
	}

	childTrees := []*LanguageTree{}

	for _, injection := range injections.Query(query, t.tree.RootNode(), source) {
		lang_id := language.Identifier(injection.Language)
		print(lang_id, lang_id.Valid())
		if !lang_id.Valid() {
			continue
		}

		if injection.Combined {
			if tree := findTreeForLanguage(lang_id, childTrees); tree != nil {
				tree.ranges = append(tree.ranges, injection.Range)
				continue
			}
		}

		tree, err := newLanguageTree(lang_id.Language(), []ts.Range{injection.Range})
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
func (t LanguageTree) GetLanguageTrees(lang_id language.Identifier) []*LanguageTree {
	results := []*LanguageTree{}

	if t.language.ID() == lang_id {
		results = append(results, &t)
	}

	for _, tree := range t.childTrees {
		results = append(results, tree.GetLanguageTrees(lang_id)...)
	}
	return results
}

func (t LanguageTree) FindCaptures(lang_id language.Identifier, query *ts.Query, source []byte, captures ...string) (treesitter.CaptureSlice, error) {
	// Build a map of index name pairs.
	captureMap := map[uint]string{}

	for _, capture := range captures {
		index, ok := query.CaptureIndexForName(capture)
		if !ok {
			return nil, fmt.Errorf("capture '%s' is not present in query", capture)
		}

		captureMap[index] = capture
	}

	cursor := ts.NewQueryCursor()
	defer cursor.Close()

	results := []treesitter.Capture{}
	for _, tree := range t.GetLanguageTrees(lang_id) {
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

func (t LanguageTree) FindTags(lang_id language.Identifier, source []byte, tags ...string) (treesitter.CaptureSlice, error) {
	query, err := treesitter.GetTagsQuery(lang_id)
	if err != nil {
		return nil, err
	}
	return t.FindCaptures(lang_id, query, source, tags...)
}

// Find all trees of a given language that includes the node
func (t *LanguageTree) GetLanguageTreesWithNode(id language.Identifier, node *ts.Node) []*LanguageTree {
	results := []*LanguageTree{}

	if t.language.ID() == id && treesitter.RangeOverlap(t.Root().Range(), node.Range()) {
		results = append(results, t)
	}

	for _, tree := range t.childTrees {
		results = append(results, tree.GetLanguageTreesWithNode(id, node)...)
	}
	return results
}

func findTreeForLanguage(id language.Identifier, trees []*LanguageTree) *LanguageTree {
	for _, tree := range trees {
		if tree.language.ID() == id {
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
	str := printer.Indent(depth) + "<" + tree.language.Name()
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

	str += printer.Indent(depth) + "</" + tree.language.Name() + ">\n"

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
