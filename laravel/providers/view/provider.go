package view

import (
	"path"

	"github.com/laravel-ls/laravel-ls/file"
	"github.com/laravel-ls/laravel-ls/laravel"
	"github.com/laravel-ls/laravel-ls/laravel/providers/view/queries"
	"github.com/laravel-ls/laravel-ls/lsp/protocol"
	"github.com/laravel-ls/laravel-ls/provider"
)

type Provider struct {
	rootPath string
	fs       filesystem
}

func NewProvider() *Provider {
	return &Provider{
		fs: filesystem{
			// TODO: should get the actual directories
			// from current laravel configuration.
			paths: []string{"resources/views"},
		},
	}
}

func (p *Provider) Init(ctx provider.InitContext) {
	p.rootPath = ctx.RootPath
}

func (p *Provider) Register(manager *provider.Manager) {
	manager.Register(file.TypePHP, p)
	manager.Register(file.TypeBlade, p)
}

// resolve view() calls to view files.
func (p *Provider) ResolveDefinition(ctx provider.DefinitionContext) {
	node := queries.ViewNameAtPosition(ctx.File, ctx.Position)

	if node != nil {
		name := queries.GetViewName(node, ctx.File.Src)

		if len(name) < 1 {
			return
		}

		viewFile := laravel.ViewFromName(name)

		ctx.Logger.Info(viewFile)

		fullPath, found := p.fs.find(p.rootPath, viewFile.Filename())

		ctx.Logger.Infof("%s %v", fullPath, found)

		if found {
			ctx.Publish(protocol.Location{
				URI: path.Join(p.rootPath, fullPath),
			})
		}
	}
}

func (p *Provider) ResolveCompletion(ctx provider.CompletionContext) {
	node := queries.ViewNameAtPosition(ctx.File, ctx.Position)

	ctx.Logger.Debug(node)

	if node != nil {
		text := queries.GetViewName(node, ctx.File.Src)

		ctx.Logger.Debug(text)

		results, err := p.fs.search(p.rootPath, text)
		if err != nil {
			ctx.Logger.WithError(err).Error("failed to search view files")
			return
		}

		for _, result := range results {
			ctx.Publish(protocol.CompletionItem{
				Label:  result.Name(),
				Detail: result.Path(),
				Kind:   protocol.CompletionItemKindFile,
			})
		}
	}
}

func (p *Provider) Diagnostic(ctx provider.DiagnosticContext) {
	// Find all view calls in the file.
	for _, capture := range queries.ViewNames(ctx.File) {
		name := queries.GetViewName(&capture.Node, ctx.File.Src)

		// Report diagnostic if view does not exist.
		if !p.fs.exists(p.rootPath, name) {
			ctx.Publish(provider.Diagnostic{
				Range:    capture.Node.Range(),
				Severity: protocol.SeverityError,
				Message:  "View not found",
			})
		}
	}
}

func (p *Provider) Hover(ctx provider.HoverContext) {
	node := queries.ViewNameAtPosition(ctx.File, ctx.Position)

	if node != nil {
		name := queries.GetViewName(node, ctx.File.Src)
		if len(name) < 1 {
			return
		}

		if view, found := p.fs.findView(p.rootPath, name); found {
			ctx.Publish(provider.Hover{
				Content: view.Path(),
			})
		}
	}
}
