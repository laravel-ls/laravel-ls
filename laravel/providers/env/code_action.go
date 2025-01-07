package env

import (
	"github.com/laravel-ls/laravel-ls/lsp/protocol"
)

func codeAction(uri, title string, line int, text string) protocol.CodeAction {
	return protocol.CodeAction{
		Title: title,
		Kind:  protocol.CodeActionQuickFix,
		Edit: &protocol.WorkspaceEdit{
			Changes: map[string][]protocol.TextEdit{
				uri: {
					{
						Range: protocol.Range{
							Start: protocol.Position{Line: line, Character: 0},
							End:   protocol.Position{Line: line, Character: len(text)},
						},
						NewText: text,
					},
				},
			},
		},
	}
}
