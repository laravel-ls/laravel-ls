package env

import (
	"fmt"
	"path"

	"github.com/laravel-ls/laravel-ls/cache"
	"github.com/laravel-ls/laravel-ls/file"
	"github.com/laravel-ls/laravel-ls/laravel/providers/env/queries"
	"github.com/laravel-ls/laravel-ls/lsp/protocol"
	"github.com/laravel-ls/laravel-ls/provider"
	"github.com/laravel-ls/laravel-ls/treesitter/php"

	log "github.com/sirupsen/logrus"
)

type Provider struct {
	rootPath string

	// Repository for key,value pairs in .env
	repo Repository

	// Repository for key,value pairs in .env.example
	exampleRepo Repository
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

func updateRepoFile(logger *log.Entry, FileCache *cache.FileCache, filename string, repo *Repository) bool {
	envFile, err := FileCache.Open(filename)
	if err != nil {
		logger.WithField("filename", filename).
			WithError(err).Error("failed to open env file")
		return false
	}

	if err := repo.Load(envFile); err != nil {
		logger.WithField("filename", filename).
			WithError(err).Error("failed to parse env file")
		return false
	}

	return true
}

func (p *Provider) updateRepo(logger *log.Entry, FileCache *cache.FileCache) bool {
	filename := path.Join(p.rootPath, ".env")

	// example file is optional, so don't return false if it fails.
	updateRepoFile(logger, FileCache, filename+".example", &p.exampleRepo)

	return updateRepoFile(logger, FileCache, filename, &p.repo)
}

func (p *Provider) ResolveCodeAction(ctx provider.CodeActionContext) {
	nodes := queries.EnvCalls(ctx.File).In(ctx.Range)

	if len(nodes) > 0 && !p.updateRepo(ctx.Logger, ctx.FileCache) {
		return
	}

	for _, node := range nodes {
		key := php.GetStringContent(node, ctx.File.Src)
		if len(key) < 1 {
			return
		}

		if _, found := p.repo.Get(key); !found {
			uri := "file://" + path.Join(p.rootPath, ".env")
			envFile := ctx.FileCache.Get(path.Join(p.rootPath, ".env"))
			line := int(envFile.Tree.Root().EndPosition().Row)

			if meta, found := p.exampleRepo.Get(key); found {
				text := fmt.Sprintf("%s=%s", key, meta.Value)
				ctx.Publish(codeAction(uri, "Copy value from .env.example", line, text))
			}
			ctx.Publish(codeAction(uri, "Add value to .env file", line, key+"="))
		}
	}
}

func (p *Provider) Hover(ctx provider.HoverContext) {
	node := queries.EnvCalls(ctx.File).At(ctx.Position)

	if node != nil {
		if !p.updateRepo(ctx.Logger, ctx.FileCache) {
			return
		}

		key := php.GetStringContent(node, ctx.File.Src)
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
	node := queries.EnvCalls(ctx.File).At(ctx.Position)
	if node != nil {
		if !p.updateRepo(ctx.Logger, ctx.FileCache) {
			return
		}

		key := php.GetStringContent(node, ctx.File.Src)
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
	node := queries.EnvCalls(ctx.File).At(ctx.Position)

	if node != nil {
		if !p.updateRepo(ctx.Logger, ctx.FileCache) {
			return
		}

		text := php.GetStringContent(node, ctx.File.Src)

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
			key := php.GetStringContent(&capture.Node, ctx.File.Src)

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
