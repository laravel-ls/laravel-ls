package route

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/laravel-ls/laravel-ls/utils/repository"
	"github.com/laravel-ls/protocol"
)

// Follow format from - https://github.com/laravel/vs-code-extension/blob/v1.0.11/src/features/route.ts#L192-L207
func formatCompetionItem(key string, meta repository.RouteEntry) protocol.CompletionItem {
	return protocol.CompletionItem{
		Label:  key,
		Kind:   protocol.CompletionItemKindEnum,
		Detail: fmt.Sprintf("%s\n\n[%s] %s", meta.Action, meta.Method, meta.URI),
	}
}

func formatLocation(meta repository.RouteEntry) protocol.Location {
	return protocol.Location{
		URI: protocol.DocumentURI(meta.File),
		Range: protocol.Range{
			Start: protocol.Position{
				Line: uint32(meta.Line),
			},
			End: protocol.Position{
				Line: uint32(meta.Line),
			},
		},
	}
}

func formatHoverContent(rootPath string, meta repository.RouteEntry) string {
	content := []string{}

	if meta.Action == "Closure" {
		content = append(content, "[Closure]")
	} else {
		content = append(content, meta.Action)
	}

	if relativePath, err := filepath.Rel(rootPath, meta.File); err == nil {
		content = append(content, fmt.Sprintf("[%s](%s)", relativePath, meta.File))
	}

	// Follow format from - https://github.com/laravel/vs-code-extension/blob/v1.0.11/src/features/route.ts#L110-L115
	return strings.Join(content, "\n\n")
}
