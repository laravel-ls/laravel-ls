package view

import (
	"fmt"
	"path/filepath"

	"github.com/laravel-ls/laravel-ls/lsp/protocol"
)

func createViewCodeAction(root, filename string) protocol.CodeAction {
	relPath, _ := filepath.Rel(root, filename)
	return protocol.CodeAction{
		Title: fmt.Sprintf("Create view (%s)", relPath),
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
