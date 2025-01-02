package treesitter_test

import (
	"fmt"
	"testing"

	"laravel-ls/treesitter"

	ts "github.com/tree-sitter/go-tree-sitter"
)

func TestPointInRange(t *testing.T) {
	tests := []struct {
		p     ts.Point
		r     ts.Range
		wants bool
	}{
		{ts.Point{Row: 2, Column: 4}, ts.Range{StartByte: 0, EndByte: 0, StartPoint: ts.Point{Row: 2, Column: 4}, EndPoint: ts.Point{Row: 3, Column: 5}}, true},
		{ts.Point{Row: 3, Column: 39}, ts.Range{StartByte: 0, EndByte: 0, StartPoint: ts.Point{Row: 2, Column: 4}, EndPoint: ts.Point{Row: 4, Column: 1}}, true},
		{ts.Point{Row: 2, Column: 4}, ts.Range{StartByte: 0, EndByte: 0, StartPoint: ts.Point{Row: 2, Column: 6}, EndPoint: ts.Point{Row: 3, Column: 5}}, false},
		{ts.Point{Row: 2, Column: 4}, ts.Range{StartByte: 0, EndByte: 0, StartPoint: ts.Point{Row: 3, Column: 0}, EndPoint: ts.Point{Row: 3, Column: 5}}, false},
		{ts.Point{Row: 13, Column: 29}, ts.Range{StartByte: 0, EndByte: 0, StartPoint: ts.Point{Row: 10, Column: 0}, EndPoint: ts.Point{Row: 12, Column: 39}}, false},
	}

	for i, args := range tests {
		t.Run(fmt.Sprintf("#%d", i+1), func(t *testing.T) {
			actual := treesitter.PointInRange(args.p, args.r)
			if actual != args.wants {
				t.Errorf("%t is not equal to %t", actual, args.wants)
			}
		})
	}
}

