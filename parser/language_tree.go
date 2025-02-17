package parser

import (
	"fmt"

	"github.com/laravel-ls/laravel-ls/treesitter"
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

func (langTree *LanguageTree) Root() *ts.Node {
	return langTree.tree.RootNode()
}

func (langTree *LanguageTree) InRange(r ts.Range) bool {
	for _, c := range langTree.ranges {
		if treesitter.RangeOverlap(c, r) {
			return true
		}
	}
	return false
}

func (langTree *LanguageTree) update(edit *ts.InputEdit) {
	langTree.tree.Edit(edit)

	// Update all children
	for _, child := range langTree.childTrees {
		child.update(edit)
	}
}

func (langTree *LanguageTree) parse(source []byte) error {
	// Parse main tree first.
	err := langTree.parser.SetIncludedRanges(langTree.ranges)
	if err != nil {
		return err
	}

	langTree.tree = langTree.parser.Parse(source, langTree.tree)

	// Then parse any injections
	return langTree.parseInjections(source)
}

func (langTree *LanguageTree) parseInjections(source []byte) error {
	query, err := treesitter.GetInjectionQuery(langTree.language.ID())
	if err != nil {
		if err == treesitter.ErrQueryNotFound {
			return nil
		}
		return err
	}

	childTrees := []*LanguageTree{}

	for _, injection := range injections.Query(query, langTree.tree.RootNode(), source) {
		langID := language.Identifier(injection.Language)
		print(langID, langID.Valid())
		if !langID.Valid() {
			continue
		}

		if injection.Combined {
			if tree := findTreeForLanguage(langID, childTrees); tree != nil {
				tree.ranges = append(tree.ranges, injection.Range)
				continue
			}
		}

		tree, err := newLanguageTree(langID.Language(), []ts.Range{injection.Range})
		if err != nil {
			return err
		}

		childTrees = append(childTrees, tree)
	}

	// Parse all children
	langTree.childTrees = childTrees
	for _, child := range langTree.childTrees {
		if err := child.parse(source); err != nil {
			return err
		}
	}

	return nil
}

// Get all trees for a particular language
func (langTree LanguageTree) GetLanguageTrees(langID language.Identifier) []*LanguageTree {
	results := []*LanguageTree{}

	if langTree.language.ID() == langID {
		results = append(results, &langTree)
	}

	for _, tree := range langTree.childTrees {
		results = append(results, tree.GetLanguageTrees(langID)...)
	}
	return results
}

func (langTree LanguageTree) FindCaptures(langID language.Identifier, query *ts.Query, source []byte, captures ...string) (treesitter.CaptureSlice, error) {
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
	for _, tree := range langTree.GetLanguageTrees(langID) {
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
func (langTree *LanguageTree) GetLanguageTreesWithNode(id language.Identifier, node *ts.Node) []*LanguageTree {
	results := []*LanguageTree{}

	if langTree.language.ID() == id && treesitter.RangeOverlap(langTree.Root().Range(), node.Range()) {
		results = append(results, langTree)
	}

	for _, tree := range langTree.childTrees {
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
