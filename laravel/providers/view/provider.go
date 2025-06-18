package view

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/laravel-ls/laravel-ls/file"
	"github.com/laravel-ls/laravel-ls/laravel/providers/view/queries"
	"github.com/laravel-ls/laravel-ls/laravel/view"
	"github.com/laravel-ls/laravel-ls/lsp/protocol"
	"github.com/laravel-ls/laravel-ls/provider"
	"github.com/laravel-ls/laravel-ls/treesitter/php"
	"github.com/spf13/afero"
)

type Provider struct {
	rootPath string
	finder   *view.Finder
}

func NewProvider() *Provider {
	return &Provider{
		finder: view.NewFinder(afero.NewOsFs()),
	}
}

func (p *Provider) Init(ctx provider.InitContext) {
	p.rootPath = ctx.RootPath
	p.finder.AddLocation(path.Join(p.rootPath, "resources/views"))
	p.finder.RegisterExtensions(".blade.php")
}

func (p *Provider) Register(manager *provider.Manager) {
	manager.Register(file.TypePHP, p)
	manager.Register(file.TypeBlade, p)
}

// resolve view() calls to view files.
func (p *Provider) ResolveDefinition(ctx provider.DefinitionContext) {
	node := queries.ViewNames(ctx.File).At(ctx.Position)

	if node != nil {
		name := php.GetStringContent(node, ctx.File.Src)

		if len(name) < 1 {
			return
		}

		fullPath, found := p.finder.Find(name)

		ctx.Logger.Debugf("%s %v", fullPath, found)

		if found {
			ctx.Publish(protocol.Location{
				URI: protocol.DocumentURI(fullPath),
			})
		}
	}
}

func (p *Provider) ResolveCompletion(ctx provider.CompletionContext) {
	node := queries.ViewNames(ctx.File).At(ctx.Position)

	if node != nil {
		input := php.GetStringContent(node, ctx.File.Src)

		kind := protocol.CompletionItemKindFile
		for _, result := range p.finder.Search(input) {
			rel, err := filepath.Rel(p.rootPath, result.Path())
			if err != nil {
				return
			}

			ctx.Publish(protocol.CompletionItem{
				Label:  result.Name(),
				Detail: rel,
				Kind:   &kind,
			})
		}
	}
}

func (p *Provider) Diagnostic(ctx provider.DiagnosticContext) {
	// Find all view calls in the file.
	for _, capture := range queries.ViewNames(ctx.File) {
		name := php.GetStringContent(&capture.Node, ctx.File.Src)

		// Report diagnostic if view does not exist.
		if _, found := p.finder.Find(name); !found {
			ctx.Publish(provider.Diagnostic{
				Range:    capture.Node.Range(),
				Severity: protocol.DiagnosticSeverityError,
				Message:  "View not found",
			})
		}
	}
}

func (p *Provider) Hover(ctx provider.HoverContext) {
	node := queries.ViewNames(ctx.File).At(ctx.Position)

	if node != nil {
		name := php.GetStringContent(node, ctx.File.Src)
		if len(name) < 1 {
			return
		}

		if viewPath, found := p.finder.Find(name); found {
			rel, err := filepath.Rel(p.rootPath, viewPath)
			if err != nil {
				return
			}

			ctx.Publish(provider.Hover{
				Content: rel,
			})
		}
	}
}

func (p *Provider) ResolveCodeAction(ctx provider.CodeActionContext) {
	nodes := queries.ViewNames(ctx.File).In(ctx.Range)

	for _, node := range nodes {
		name := php.GetStringContent(node, ctx.File.Src)
		if len(name) < 1 {
			return
		}

		ctx.Logger.Debug(name)

		if _, found := p.finder.Find(name); !found {
			for _, filename := range p.finder.PossibleFiles(name) {
				filename, _ = filepath.Rel(p.rootPath, filename)
				kind := protocol.CodeActionRefactor
				ctx.Publish(protocol.CodeAction{
					Title: fmt.Sprintf("Create view (%s)", filename),
					Kind:  &kind,
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
				})
			}
		}
	}
}
