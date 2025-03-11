package parser

import (
	"github.com/laravel-ls/laravel-ls/treesitter"
	"github.com/laravel-ls/laravel-ls/treesitter/injections"
	"github.com/laravel-ls/laravel-ls/treesitter/language"

	ts "github.com/tree-sitter/go-tree-sitter"
)

type LanguageTree struct {
	parser       *ts.Parser
	tree         *ts.Tree
	language     *language.Language
	ranges       []ts.Range
	childTrees   []*LanguageTree
	captureCache map[*ts.Query]treesitter.CaptureSlice
}

func newLanguageTree(language *language.Language, ranges []ts.Range) (*LanguageTree, error) {
	parser := ts.NewParser()
	err := parser.SetLanguage(language.TSObject())
	if err != nil {
		return nil, err
	}

	return &LanguageTree{
		parser:       parser,
		ranges:       ranges,
		language:     language,
		childTrees:   []*LanguageTree{},
		captureCache: map[*ts.Query]treesitter.CaptureSlice{},
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

	// invalidate all capture caches.
	langTree.captureCache = map[*ts.Query]treesitter.CaptureSlice{}

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
	// return cached result on hit.
	if cache, found := langTree.captureCache[query]; found {
		return cache.Name(captures...), nil
	}

	captureMap := query.CaptureNames()

	cursor := ts.NewQueryCursor()
	defer cursor.Close()

	results := treesitter.CaptureSlice{}
	for _, tree := range langTree.GetLanguageTrees(langID) {
		matches := cursor.Matches(query, tree.Root(), source)
		for it := matches.Next(); it != nil; it = matches.Next() {
			for _, capture := range it.Captures {

				// Just to be safe.
				if int(capture.Index) > len(captureMap) {
					continue
				}

				results = append(results, treesitter.Capture{
					Name: captureMap[capture.Index],
					Node: capture.Node,
				})
			}
		}
	}

	// update cache
	langTree.captureCache[query] = results
	return results.Name(captures...), nil
}

func (langTree LanguageTree) FindTags(langID language.Identifier, source []byte, tags ...string) (treesitter.CaptureSlice, error) {
	query, err := treesitter.GetTagsQuery(langID)
	if err != nil {
		return nil, err
	}
	return langTree.FindCaptures(langID, query, source, tags...)
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
