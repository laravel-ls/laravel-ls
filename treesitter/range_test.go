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

func TestRangeOverlap(t *testing.T) {
	tests := []struct {
		a     ts.Range
		b     ts.Range
		wants bool
	}{
		{ts.Range{StartByte: 40, EndByte: 50}, ts.Range{StartByte: 30, EndByte: 100}, true},
		{ts.Range{StartByte: 30, EndByte: 100}, ts.Range{StartByte: 40, EndByte: 50}, true},

		{ts.Range{StartByte: 12, EndByte: 14}, ts.Range{StartByte: 14, EndByte: 100}, true},
		{ts.Range{StartByte: 14, EndByte: 100}, ts.Range{StartByte: 12, EndByte: 14}, true},

		{ts.Range{StartByte: 99, EndByte: 100}, ts.Range{StartByte: 14, EndByte: 100}, true},
		{ts.Range{StartByte: 14, EndByte: 100}, ts.Range{StartByte: 99, EndByte: 100}, true},

		{ts.Range{StartByte: 0, EndByte: 10}, ts.Range{StartByte: 101, EndByte: 110}, false},
		{ts.Range{StartByte: 101, EndByte: 110}, ts.Range{StartByte: 0, EndByte: 10}, false},

		{ts.Range{StartByte: 99, EndByte: 100}, ts.Range{StartByte: 200, EndByte: 310}, false},
		{ts.Range{StartByte: 200, EndByte: 310}, ts.Range{StartByte: 99, EndByte: 100}, false},
	}

	for i, args := range tests {
		t.Run(fmt.Sprintf("#%d", i+1), func(t *testing.T) {
			actual := treesitter.RangeOverlap(args.a, args.b)
			if actual != args.wants {
				t.Errorf("%t is not equal to %t", actual, args.wants)
			}
		})
	}
}
