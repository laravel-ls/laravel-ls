package view

import (
	"fmt"

	"github.com/laravel-ls/laravel-ls/lsp/protocol"
)

func createViewCodeAction(filename string) protocol.CodeAction {
	return protocol.CodeAction{
		Title: fmt.Sprintf("Create view (%s)", filename),
		Edit: &protocol.WorkspaceEdit{
			DocumentChanges: []protocol.DocumentChangeOperation{
				protocol.CreateFile{
					URI: protocol.DocumentURI("file://" + filename),
					ResourceOperation: protocol.ResourceOperation{
						Kind: "create",
					},
				},
			},
		},
	}
}
