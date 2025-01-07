package php

import (
	ts "github.com/tree-sitter/go-tree-sitter"
)

// Get content from a string or encapsed_string node.
func GetStringContent(node *ts.Node, source []byte) string {
	if (node.Kind() == "string" || node.Kind() == "encapsed_string") && node.NamedChildCount() > 0 {
		contentNode := node.NamedChild(0)
		if contentNode.Kind() == "string_content" {
			return contentNode.Utf8Text(source)
		}
	}
	return ""
}
