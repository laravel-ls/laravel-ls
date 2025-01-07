package parser

import (
	"github.com/laravel-ls/laravel-ls/treesitter"
	ts "github.com/tree-sitter/go-tree-sitter"
)

type Capture struct {
	Name string
	Node ts.Node
}

type CaptureSlice []Capture

// Find a capture at a given position
func (captures CaptureSlice) At(position ts.Point) *ts.Node {
	for _, capture := range captures {
		if treesitter.PointInRange(position, capture.Node.Range()) {
			return &capture.Node
		}
	}
	return nil
}

// Find all captures in a given range
func (captures CaptureSlice) In(r ts.Range) []*ts.Node {
	nodes := []*ts.Node{}
	for _, capture := range captures {
		if treesitter.RangeOverlap(r, capture.Node.Range()) {
			nodes = append(nodes, &capture.Node)
		}
	}
	return nodes
}
