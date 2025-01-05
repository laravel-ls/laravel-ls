package server

import (
	"laravel-ls/lsp/protocol"

	ts "github.com/tree-sitter/go-tree-sitter"
)

// func getNodeAt(file *parser.File, pos protocol.Position) *ts.Node {
// 	point := ts.Point{
// 		Row:    uint(pos.Line),
// 		Column: uint(pos.Character),
// 	}
// 	return file.GetNodeAt(point)
// }

func toTSPoint(pos protocol.Position) ts.Point {
	return ts.Point{
		Row:    uint(pos.Line),
		Column: uint(pos.Character),
	}
}

func FromTSPoint(point ts.Point) protocol.Position {
	return protocol.Position{
		Line:      int(point.Row),
		Character: int(point.Column),
	}
}

func toTSRange(r protocol.Range) ts.Range {
	return ts.Range{
		StartPoint: toTSPoint(r.Start),
		EndPoint:   toTSPoint(r.End),
	}
}
