package assets

import (
	"path"

	"github.com/laravel-ls/laravel-ls/file"
	"github.com/laravel-ls/laravel-ls/laravel/providers/assets/queries"
	"github.com/laravel-ls/laravel-ls/lsp/protocol"
	"github.com/laravel-ls/laravel-ls/provider"
	"github.com/laravel-ls/laravel-ls/treesitter/php"
)

type Provider struct {
	rootPath string

	repo Repository
}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) Register(manager *provider.Manager) {
	manager.Register(file.TypePHP, p)
	manager.Register(file.TypeBlade, p)
}

func (p *Provider) Init(ctx provider.InitContext) {
	p.rootPath = ctx.RootPath
}

func (p *Provider) Hover(ctx provider.HoverContext) {
	node := queries.Assets(ctx.File).At(ctx.Position)
	if node != nil {
		filename := php.GetStringContent(node, ctx.File.Src)
		if p.repo.Exists(p.rootPath, filename) {
			ctx.Publish(provider.Hover{
				Content: path.Join("public", filename),
			})
		}
	}
}

func (p *Provider) ResolveCompletion(ctx provider.CompletionContext) {
	node := queries.Assets(ctx.File).At(ctx.Position)
	if node != nil {
		filename := php.GetStringContent(node, ctx.File.Src)

		kind := protocol.CompletionItemKindFile
		for _, file := range p.repo.Search(p.rootPath, filename) {
			ctx.Publish(protocol.CompletionItem{
				Label: file,
				Kind:  &kind,
			})
		}
	}
}

func (p *Provider) ResolveDefinition(ctx provider.DefinitionContext) {
	node := queries.Assets(ctx.File).At(ctx.Position)
	if node != nil {
		filename := php.GetStringContent(node, ctx.File.Src)
		if p.repo.Exists(p.rootPath, filename) {
			ctx.Publish(protocol.Location{
				URI: protocol.DocumentURI(path.Join(p.rootPath, "public", filename)),
			})
		}
	}
}

func (p *Provider) Diagnostic(ctx provider.DiagnosticContext) {
	for _, capture := range queries.Assets(ctx.File) {
		filename := php.GetStringContent(&capture.Node, ctx.File.Src)
		if !p.repo.Exists(p.rootPath, filename) {
			ctx.Publish(provider.Diagnostic{
				Range:    capture.Node.Range(),
				Severity: protocol.DiagnosticSeverityError,
				Message:  "Asset not found",
			})
		}
	}
}
