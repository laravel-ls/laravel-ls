package server

import (
	"github.com/laravel-ls/protocol"

	ts "github.com/tree-sitter/go-tree-sitter"
)

func toTSPoint(pos protocol.Position) ts.Point {
	return ts.Point{
		Row:    uint(pos.Line),
		Column: uint(pos.Character),
	}
}

func FromTSPoint(point ts.Point) protocol.Position {
	return protocol.Position{
		Line:      uint32(point.Row),
		Character: uint32(point.Column),
	}
}

func toTSRange(r protocol.Range) ts.Range {
	return ts.Range{
		StartPoint: toTSPoint(r.Start),
		EndPoint:   toTSPoint(r.End),
	}
}
