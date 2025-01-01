package injections

import (
	ts "github.com/tree-sitter/go-tree-sitter"
)

type Capture struct {
	Language string
	Combined bool
	Range    ts.Range
}

func Query(query *ts.Query, node *ts.Node, source []byte) []Capture {
	captures := []Capture{}

	captureContent, ok := query.CaptureIndexForName("injection.content")
	if !ok {
		return captures
	}

	cursor := ts.NewQueryCursor()
	defer cursor.Close()

	matches := cursor.Matches(query, node, source)
	for match := matches.Next(); match != nil; match = matches.Next() {
		lang, combined := getInfo(query, match.PatternIndex)

		if lang == "unknown" {
			continue
		}

		for _, c := range match.Captures {

			if c.Index != uint32(captureContent) {
				continue
			}

			captures = append(captures, Capture{
				Language: lang,
				Combined: combined,
				Range:    c.Node.Range(),
			})
		}
	}
	return captures
}

func getInfo(query *ts.Query, patterIndex uint) (string, bool) {
	lang := "unknown"
	combined := false

	for _, prop := range query.PropertySettings(patterIndex) {
		switch prop.Key {
		case "injection.language":
			lang = *prop.Value
		case "injection.combined":
			combined = true
		}
	}
	return lang, combined
}
