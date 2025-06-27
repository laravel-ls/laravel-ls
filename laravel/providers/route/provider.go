package route

import (
	"fmt"

	"github.com/laravel-ls/laravel-ls/file"
	"github.com/laravel-ls/laravel-ls/laravel/providers/route/queries"
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

func (p *Provider) Register(manager *provider.Manager) {
	manager.Register(file.TypePHP, p)
}

func (p *Provider) Init(ctx provider.InitContext) {
	p.rootPath = ctx.RootPath
}

func (p *Provider) Hover(ctx provider.HoverContext) {
	node := queries.RouteCalls(ctx.File).At(ctx.Position)
	if node == nil {
		return
	}

	service := php.GetStringContent(node, ctx.File.Src)

	repo, err := ctx.Project.Routes()
	if err != nil {
		ctx.Logger.WithError(err).Warn("failed to get repo")
		return
	}
	if r, ok := repo[service]; ok {
		ctx.Publish(provider.Hover{
			// TODO: Maybe make this into a nicer format
			Content: fmt.Sprintf("Method: %v\nURI: %v\nName: %v\nAction: %v\nParameters: %v\nFilename: %v\nLine: %v",
				r.Method, r.URI, r.Name, r.Action, r.Parameters, r.File, r.Line),
		})
	}
}

func (p *Provider) ResolveCompletion(ctx provider.CompletionContext) {
	node := queries.RouteCalls(ctx.File).At(ctx.Position)
	if node == nil {
		return
	}

	text := php.GetStringContent(node, ctx.File.Src)

	repo, err := ctx.Project.Routes()
	if err != nil {
		ctx.Logger.WithError(err).Warn("failed to get repo")
		return
	}

	for key, meta := range repo.Find(text) {
		ctx.Publish(protocol.CompletionItem{
			Label:  key,
			Detail: meta.File,
			Kind:   protocol.CompletionItemKindFile,
		})
	}
}

func (p *Provider) ResolveDefinition(ctx provider.DefinitionContext) {
	node := queries.RouteCalls(ctx.File).At(ctx.Position)

	repo, err := ctx.Project.Routes()
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
			URI: protocol.DocumentURI(meta.File),
			Range: protocol.Range{
				Start: protocol.Position{
					Line: uint32(meta.Line - 1),
				},
			},
		})
	}
}

func (p *Provider) Diagnostic(ctx provider.DiagnosticContext) {
	captures := queries.RouteCalls(ctx.File)

	if len(captures) < 1 {
		return
	}

	repo, err := ctx.Project.Routes()
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
				Severity: protocol.DiagnosticSeverityError,
				Message:  fmt.Sprintf("Route '%s' does not exist", key),
			})
		}
	}
}
