package env

import (
	"path"

	"github.com/shufflingpixels/laravel-ls/file"
	"github.com/shufflingpixels/laravel-ls/laravel/providers/env/queries"
	"github.com/shufflingpixels/laravel-ls/lsp/protocol"
	"github.com/shufflingpixels/laravel-ls/provider"
)

type Provider struct {
	rootPath string
	repo     Repository
}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) Register(manager *provider.Manager) {
	manager.Register(file.TypePHP, p)
}

func (p *Provider) Init(ctx provider.InitContext) {
	p.rootPath = ctx.RootPath

	filePath := path.Join(p.rootPath, ".env")

	envFile, err := ctx.FileCache.Open(filePath)
	if err != nil {
		ctx.Logger.WithError(err).Error("failed to open env file")
	}

	if err := p.repo.Load(envFile); err != nil {
		ctx.Logger.WithError(err).Error("failed to parse env file")
	}
}

func (p *Provider) Hover(ctx provider.HoverContext) string {
	node := queries.EnvCallAtPosition(ctx.File, ctx.Position)

	if node != nil {
		key := queries.GetKey(node, ctx.File.Src)
		if len(key) < 1 {
			return ""
		}

		if meta, found := p.repo.Get(key); found {
			if len(meta.Value) < 1 {
				return "[empty]"
			}
			return meta.Value
		}
	}
	return ""
}

// resolve env() calls to variable
func (p *Provider) ResolveDefinition(ctx provider.DefinitionContext) {
	node := queries.EnvCallAtPosition(ctx.File, ctx.Position)
	if node != nil {
		key := queries.GetKey(node, ctx.File.Src)
		if meta, found := p.repo.Get(key); found {
			ctx.Publish(protocol.Location{
				URI: path.Join(p.rootPath, ".env"),
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      meta.Line,
						Character: meta.Column,
					},
				},
			})
		}
	}
}

func (p *Provider) ResolveCompletion(ctx provider.CompletionContext) {
	node := queries.EnvCallAtPosition(ctx.File, ctx.Position)

	if node != nil {
		text := queries.GetKey(node, ctx.File.Src)

		for key, meta := range p.repo.Find(text) {
			ctx.Publish(protocol.CompletionItem{
				Label:  key,
				Detail: meta.Value,
				Kind:   protocol.CompletionItemKindConstant,
			})
		}
	}
}

func (p *Provider) Diagnostic(ctx provider.DiagnosticContext) {
	// Find all env calls in the file.
	for _, capture := range queries.EnvCalls(ctx.File) {
		key := queries.GetKey(&capture.Node, ctx.File.Src)

		// Report diagnostic if key is not defined.
		if !p.repo.Exists(key) {
			ctx.Publish(provider.Diagnostic{
				Range:    capture.Node.Range(),
				Severity: protocol.SeverityError,
				Message:  "Environment variable is not defined",
			})
		}
	}
}
