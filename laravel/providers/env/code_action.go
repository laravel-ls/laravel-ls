package env

import (
	"github.com/laravel-ls/laravel-ls/lsp/protocol"
)

func codeAction(uri protocol.DocumentURI, title string, line int, text string) protocol.CodeAction {
	kind := protocol.CodeActionQuickFix
	return protocol.CodeAction{
		Title: title,
		Kind:  &kind,
		Edit: &protocol.WorkspaceEdit{
			Changes: map[protocol.DocumentURI][]protocol.TextEdit{
				uri: {
					{
						Range: protocol.Range{
							Start: protocol.Position{Line: uint32(line), Character: 0},
							End:   protocol.Position{Line: uint32(line), Character: uint32(len(text))},
						},
						NewText: text,
					},
				},
			},
		},
	}
}
