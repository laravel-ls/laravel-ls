package route

import "github.com/laravel-ls/protocol"

func codeAction(uri protocol.DocumentURI, title string, line uint, text string) protocol.CodeAction {
	return protocol.CodeAction{
		Title: title,
		Kind:  protocol.CodeActionQuickFix,
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
