package treesitter

import (
	"slices"

	ts "github.com/tree-sitter/go-tree-sitter"
)

type Capture struct {
	Name string
	Node ts.Node
}

type CaptureSlice []Capture

// Name finds all captures with a given name or names.
func (captures CaptureSlice) Name(name ...string) CaptureSlice {
	nodes := CaptureSlice{}
	for _, capture := range captures {
		if slices.Contains(name, capture.Name) {
			nodes = append(nodes, capture)
		}
	}
	return nodes
}

// At finds a capture at a given position
func (captures CaptureSlice) At(position ts.Point) *ts.Node {
	for _, capture := range captures {
		if PointInRange(position, capture.Node.Range()) {
			return &capture.Node
		}
	}
	return nil
}

// In finds all captures in a given range
func (captures CaptureSlice) In(r ts.Range) []*ts.Node {
	nodes := []*ts.Node{}
	for _, capture := range captures {
		if RangeOverlap(r, capture.Node.Range()) {
			nodes = append(nodes, &capture.Node)
		}
	}
	return nodes
}
