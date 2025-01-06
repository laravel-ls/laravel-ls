package treesitter_test

import (
	"fmt"
	"testing"

	"github.com/laravel-ls/laravel-ls/treesitter"

	ts "github.com/tree-sitter/go-tree-sitter"
)

func TestPointInRange(t *testing.T) {
	tests := []struct {
		p     ts.Point
		r     ts.Range
		wants bool
	}{
		{ts.Point{Row: 2, Column: 4}, ts.Range{StartByte: 0, EndByte: 0, StartPoint: ts.Point{Row: 2, Column: 4}, EndPoint: ts.Point{Row: 2, Column: 10}}, true},
		{ts.Point{Row: 2, Column: 9}, ts.Range{StartByte: 0, EndByte: 0, StartPoint: ts.Point{Row: 2, Column: 4}, EndPoint: ts.Point{Row: 2, Column: 10}}, true},
		{ts.Point{Row: 2, Column: 4}, ts.Range{StartByte: 0, EndByte: 0, StartPoint: ts.Point{Row: 2, Column: 4}, EndPoint: ts.Point{Row: 3, Column: 5}}, true},
		{ts.Point{Row: 3, Column: 39}, ts.Range{StartByte: 0, EndByte: 0, StartPoint: ts.Point{Row: 2, Column: 4}, EndPoint: ts.Point{Row: 4, Column: 1}}, true},
		{ts.Point{Row: 2, Column: 4}, ts.Range{StartByte: 0, EndByte: 0, StartPoint: ts.Point{Row: 2, Column: 6}, EndPoint: ts.Point{Row: 2, Column: 10}}, false},
		{ts.Point{Row: 2, Column: 12}, ts.Range{StartByte: 0, EndByte: 0, StartPoint: ts.Point{Row: 2, Column: 6}, EndPoint: ts.Point{Row: 2, Column: 10}}, false},
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
	r := func(s_row, s_col, e_row, e_col uint) ts.Range {
		return ts.Range{
			StartPoint: ts.Point{Row: s_row, Column: s_col},
			EndPoint:   ts.Point{Row: e_row, Column: e_col},
		}
	}

	tests := []struct {
		name     string
		a        ts.Range
		b        ts.Range
		expected bool
	}{
		// Overlapping cases
		{
			name:     "Fully overlapping ranges",
			a:        r(1, 0, 2, 0),
			b:        r(1, 5, 2, 5),
			expected: true,
		},
		{
			name:     "Partially overlapping ranges",
			a:        r(1, 0, 3, 0),
			b:        r(2, 0, 4, 0),
			expected: true,
		},
		// Non-overlapping cases
		{
			name:     "Non-overlapping ranges (disjoint rows)",
			a:        r(1, 0, 2, 0),
			b:        r(3, 0, 4, 0),
			expected: false,
		},
		{
			name:     "Non-overlapping ranges (same row, disjoint columns)",
			a:        r(1, 0, 1, 5),
			b:        r(1, 6, 1, 10),
			expected: false,
		},
		// Edge cases
		{
			name:     "Adjacent ranges (overlap)",
			a:        r(1, 0, 1, 5),
			b:        r(1, 5, 1, 10),
			expected: true,
		},
		{
			name:     "Adjacent ranges (non-overlap)",
			a:        r(1, 0, 1, 5),
			b:        r(1, 6, 1, 10),
			expected: false,
		},
		{
			name:     "Same range",
			a:        r(1, 0, 2, 0),
			b:        r(1, 0, 2, 0),
			expected: true,
		},
		{
			name:     "Single-point ranges (overlapping)",
			a:        r(1, 0, 1, 0),
			b:        r(1, 0, 1, 0),
			expected: true,
		},
		{
			name:     "Single-point ranges (non-overlapping)",
			a:        r(1, 0, 1, 0),
			b:        r(1, 1, 1, 1),
			expected: false,
		},
	}

	for _, args := range tests {
		t.Run(args.name, func(t *testing.T) {
			actual := treesitter.RangeOverlap(args.a, args.b)
			if actual != args.expected {
				t.Errorf("%t is not equal to %t", actual, args.expected)
			}
		})
	}
}

func TestRangeOverlapBytes(t *testing.T) {
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
			actual := treesitter.RangeOverlapBytes(args.a, args.b)
			if actual != args.wants {
				t.Errorf("%t is not equal to %t", actual, args.wants)
			}
		})
	}
}
