package env

import (
	"path"

	"github.com/laravel-ls/laravel-ls/cache"
	"github.com/laravel-ls/laravel-ls/file"
	"github.com/laravel-ls/laravel-ls/laravel/providers/env/queries"
	"github.com/laravel-ls/laravel-ls/lsp/protocol"
	"github.com/laravel-ls/laravel-ls/provider"

	log "github.com/sirupsen/logrus"
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
}

func (p *Provider) updateRepo(logger *log.Entry, FileCache *cache.FileCache) bool {
	filePath := path.Join(p.rootPath, ".env")

	envFile, err := FileCache.Open(filePath)
	if err != nil {
		logger.WithError(err).Error("failed to open env file")
		return false
	}

	if err := p.repo.Load(envFile); err != nil {
		logger.WithError(err).Error("failed to parse env file")
		return false
	}

	return true
}

func (p *Provider) Hover(ctx provider.HoverContext) {
	node := queries.EnvCallAtPosition(ctx.File, ctx.Position)

	if node != nil {
		if !p.updateRepo(ctx.Logger, ctx.FileCache) {
			return
		}

		key := queries.GetKey(node, ctx.File.Src)
		if len(key) < 1 {
			return
		}

		content := "[undefined]"
		if meta, found := p.repo.Get(key); found {
			if len(meta.Value) < 1 {
				content = "[empty]"
			} else {
				content = meta.Value
			}
		}

		ctx.Publish(provider.Hover{
			Content: content,
		})
	}
}

// resolve env() calls to variable
func (p *Provider) ResolveDefinition(ctx provider.DefinitionContext) {
	node := queries.EnvCallAtPosition(ctx.File, ctx.Position)
	if node != nil {
		if !p.updateRepo(ctx.Logger, ctx.FileCache) {
			return
		}

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
		if !p.updateRepo(ctx.Logger, ctx.FileCache) {
			return
		}

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
	captures := queries.EnvCalls(ctx.File)

	if len(captures) > 0 {
		if !p.updateRepo(ctx.Logger, ctx.FileCache) {
			return
		}

		for _, capture := range captures {
			key := queries.GetKey(&capture.Node, ctx.File.Src)

			// Report diagnostic if key is not defined
			// and no default value is given
			if !p.repo.Exists(key) && !queries.HasDefault(&capture.Node) {
				ctx.Publish(provider.Diagnostic{
					Range:    capture.Node.Range(),
					Severity: protocol.SeverityError,
					Message:  "Environment variable is not defined",
				})
			}
		}
	}
}
