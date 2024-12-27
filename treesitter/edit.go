package treesitter

import (
	"bytes"

	ts "github.com/tree-sitter/go-tree-sitter"
)

// Calculate a treesitter InputEdit object.
func CalculateEdit(start, end ts.Point, oldSrc []byte, newSrc []byte) *ts.InputEdit {
	startIndex := PointToByteOffset(start, oldSrc)
	oldEndIndex := PointToByteOffset(end, oldSrc)

	return &ts.InputEdit{
		StartByte:      startIndex,
		OldEndByte:     oldEndIndex,
		NewEndByte:     startIndex + uint(len(newSrc)),
		StartPosition:  start,
		OldEndPosition: end,
		NewEndPosition: CalculateNewEndPoint(start, newSrc),
	}
}

// Get the byte count at an specific row.
func GetRowByte(row uint, src []byte) uint {
	current := uint(0)

	// Find the start of the row by iterating through the newlines
	for i, b := range src {
		if b == '\n' {
			current++
			if current == row {
				return uint(i) + 1
			}
		}
	}
	return 0
}

// Get the byte offset for a point.
func PointToByteOffset(point ts.Point, src []byte) uint {
	// Total byte offset = row + column offset
	return GetRowByte(point.Row, src) + point.Column
}

// Calculate a new endpoint
func CalculateNewEndPoint(start ts.Point, src []byte) ts.Point {
	lines := bytes.Split(src, []byte("\n"))
	if len(lines) == 1 {
		// Single line edit: only update the column
		return ts.Point{
			Row:    start.Row,
			Column: start.Column + uint(len(lines[0])),
		}
	}
	// Multi-line edit: update row and column
	return ts.Point{
		Row:    start.Row + uint(len(lines)-1),
		Column: uint(len(lines[len(lines)-1])),
	}
}
