package app

import (
	"fmt"

	"github.com/laravel-ls/laravel-ls/file"
	"github.com/laravel-ls/laravel-ls/laravel/providers/app/queries"
	"github.com/laravel-ls/laravel-ls/lsp/protocol"
	"github.com/laravel-ls/laravel-ls/provider"
	"github.com/laravel-ls/laravel-ls/treesitter/php"
)

type Provider struct {
	rootPath string
}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) Init(ctx provider.InitContext) {
	p.rootPath = ctx.RootPath
}

func (p *Provider) Register(manager *provider.Manager) {
	manager.Register(file.TypePHP, p)
	manager.Register(file.TypeBlade, p)
}

func (p *Provider) Hover(ctx provider.HoverContext) {
	node := queries.AppCalls(ctx.File).At(ctx.Position)
	if node == nil {
		return
	}

	service := php.GetStringContent(node, ctx.File.Src)

	repo, err := ctx.Project.AppBindings()
	if err != nil {
		ctx.Logger.WithError(err).Warn("failed to get repo")
		return
	}
	if r, ok := repo[service]; ok {
		ctx.Publish(provider.Hover{
			Content: r.Class,
		})
	}
}

func (p *Provider) ResolveCompletion(ctx provider.CompletionContext) {
	node := queries.AppCalls(ctx.File).At(ctx.Position)
	if node == nil {
		return
	}

	text := php.GetStringContent(node, ctx.File.Src)

	repo, err := ctx.Project.AppBindings()
	if err != nil {
		ctx.Logger.WithError(err).Warn("failed to get repo")
		return
	}

	for key, meta := range repo.Find(text) {
		ctx.Publish(protocol.CompletionItem{
			Label:  key,
			Detail: meta.Class,
			Kind:   protocol.CompletionItemKindClass,
		})
	}
}

func (p *Provider) ResolveDefinition(ctx provider.DefinitionContext) {
	node := queries.AppCalls(ctx.File).At(ctx.Position)

	repo, err := ctx.Project.AppBindings()
	if err != nil {
		ctx.Logger.WithError(err).Warn("failed to get repo")
		return
	}

	if node == nil {
		return
	}

	key := php.GetStringContent(node, ctx.File.Src)
	if meta, found := repo.Get(key); found {
		ctx.Publish(protocol.Location{
			URI: meta.Path,
			Range: protocol.Range{
				Start: protocol.Position{
					Line: meta.Line - 1,
				},
			},
		})
	}
}

func (p *Provider) Diagnostic(ctx provider.DiagnosticContext) {
	captures := queries.AppCalls(ctx.File)

	if len(captures) < 1 {
		return
	}

	repo, err := ctx.Project.AppBindings()
	if err != nil {
		ctx.Logger.WithError(err).Warn("failed to get repo")
		return
	}

	for _, capture := range captures {
		key := php.GetStringContent(&capture.Node, ctx.File.Src)

		// Report diagnostic if key is not defined
		// and no default value is given
		if !repo.Exists(key) {
			ctx.Publish(provider.Diagnostic{
				Range:    capture.Node.Range(),
				Severity: protocol.SeverityError,
				Message:  fmt.Sprintf("application binding '%s' does not exist", key),
			})
		}
	}
}
