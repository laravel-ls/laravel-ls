package treesitter

import (
	ts "github.com/tree-sitter/go-tree-sitter"
)

// PointInRange returns true if point p is inside range r
func PointInRange(p ts.Point, r ts.Range) bool {
	// Point can not be in range if its row is
	// before start row or after end row.
	if p.Row < r.StartPoint.Row || p.Row > r.EndPoint.Row {
		return false
	}

	// If the point is at the same row as the start row and the
	// points column is before the start column. Point is not in range
	if p.Row == r.StartPoint.Row && p.Column < r.StartPoint.Column {
		return false
	}

	// If the point is at the same row as the end row and the
	// points column is after the end column. Point is not in range
	if p.Row == r.EndPoint.Row && p.Column > r.EndPoint.Column {
		return false
	}

	// If all other tests fails. Point must be inside range
	return true
}

// RangeOverlap returns true if range a and b overlaps. (calculated using point fields)
func RangeOverlap(a, b ts.Range) bool {
	return PointInRange(a.StartPoint, b) || PointInRange(b.StartPoint, a)
}

// RangeOverlapBytes returns true if range a and b overlaps (calculated using byte fields)
func RangeOverlapBytes(a, b ts.Range) bool {
	return (a.StartByte >= b.StartByte && a.StartByte <= b.EndByte) ||
		(b.StartByte >= a.StartByte && b.StartByte <= a.EndByte)
}
