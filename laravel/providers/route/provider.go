package route

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/laravel-ls/laravel-ls/file"
	"github.com/laravel-ls/laravel-ls/laravel/providers/route/queries"
	"github.com/laravel-ls/laravel-ls/provider"
	"github.com/laravel-ls/laravel-ls/treesitter/php"
	"github.com/laravel-ls/protocol"
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

	text := php.GetStringContent(node, ctx.File.Src)
	if len(text) < 1 {
		return
	}

	repo, err := ctx.Project.Routes()
	if err != nil {
		ctx.Logger.WithError(err).Warn("failed to get repo")
		return
	}

	route, ok := repo.Get(text)
	if !ok {
		return
	}

	content := []string{}

	if route.Action == "Closure" {
		content = append(content, "[Closure]")
	} else {
		content = append(content, route.Action)
	}

	if relativePath, err := filepath.Rel(p.rootPath, route.File); err == nil {
		content = append(content, fmt.Sprintf("[%s](%s)", relativePath, route.File))
	}

	// Follow format from - https://github.com/laravel/vs-code-extension/blob/v1.0.11/src/features/route.ts#L110-L115
	ctx.Publish(provider.Hover{Content: strings.Join(content, "\n\n")})
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
		// Follow format from - https://github.com/laravel/vs-code-extension/blob/v1.0.11/src/features/route.ts#L192-L207
		ctx.Publish(protocol.CompletionItem{
			Label:  key,
			Kind:   protocol.CompletionItemKindEnum,
			Detail: fmt.Sprintf("%s\n\n[%s] %s", meta.Action, meta.Method, meta.URI),
		})
	}
}

func (p *Provider) ResolveDefinition(ctx provider.DefinitionContext) {
	node := queries.RouteCalls(ctx.File).At(ctx.Position)
	if node == nil {
		return
	}

	route := php.GetStringContent(node, ctx.File.Src)
	if len(route) < 1 {
		return
	}

	repo, err := ctx.Project.Routes()
	if err != nil {
		ctx.Logger.WithError(err).Warn("failed to get repo")
		return
	}

	if meta, found := repo.Get(route); found {
		ctx.Publish(protocol.Location{
			URI: protocol.DocumentURI(meta.File),
			// TODO: Maybe refactor this into a helper function
			Range: protocol.Range{
				Start: protocol.Position{
					Line: uint32(meta.Line),
				},
				End: protocol.Position{
					Line: uint32(meta.Line),
				},
			},
		})
	}
}

func (p *Provider) Diagnostic(ctx provider.DiagnosticContext) {
	node := queries.RouteCalls(ctx.File)
	if len(node) < 1 {
		return
	}

	repo, err := ctx.Project.Routes()
	if err != nil {
		ctx.Logger.WithError(err).Warn("failed to get repo")
		return
	}

	for _, capture := range node {
		text := php.GetStringContent(&capture.Node, ctx.File.Src)
		if len(text) < 1 || repo.Exists(text) {
			continue
		}

		// Follow format from:
		// https://github.com/laravel/vs-code-extension/blob/v1.0.11/src/features/route.ts#L137-L142
		// https://github.com/laravel/vs-code-extension/blob/main/src/diagnostic/index.ts#L3-L14
		ctx.Publish(provider.Diagnostic{
			Range:    capture.Node.Range(),
			Severity: protocol.DiagnosticSeverityWarning,
			Message:  fmt.Sprintf("Route [%s] not found", text),
		})
	}
}

func (p *Provider) ResolveCodeAction(ctx provider.CodeActionContext) {
	nodes := queries.RouteCalls(ctx.File).In(ctx.Range)
	if len(nodes) < 1 {
		return
	}

	repo, err := ctx.Project.Routes()
	if err != nil {
		ctx.Logger.WithError(err).Warn("failed to get repo")
		return
	}

	routeFilename := path.Join(p.rootPath, "routes/web.php")
	routeFile, err := ctx.FileCache.Open(routeFilename)
	if err != nil {
		ctx.Logger.WithError(err).Warn("failed to parse routes/web.php file")
		return
	}

	for _, node := range nodes {
		text := php.GetStringContent(node, ctx.File.Src)
		if len(text) < 1 {
			return
		}

		if _, found := repo.Get(text); !found {
			uri := protocol.DocumentURI("file://" + routeFilename)
			line := routeFile.Tree.Root().EndPosition().Row

			code := fmt.Sprintf("\nRoute::get('/', function() {\n    return view('');\n})->name('%s');\n", text)

			ctx.Publish(codeAction(uri, "Add to routes file (web.php)", line, code))
		}
	}
}
