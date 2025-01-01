package provider

import (
	"laravel-ls/file"
	"laravel-ls/parser"

	"github.com/sirupsen/logrus"
)

type Language struct {
	CompletionProviders  []CompletionProvider
	DiagnosticsProviders []DiagnosticProvider
	DefinitionProviders  []DefinitionProvider
	HoverProviders       []HoverProvider
}

type Manager struct {
	providers []Provider
	languages map[file.Type]Language
}

func NewManager() *Manager {
	return &Manager{
		languages: map[file.Type]Language{},
	}
}

func (m *Manager) Init(ctx InitContext) {
	for _, provider := range m.providers {
		provider.Init(ctx)
	}
}

func (m *Manager) Add(provider Provider) {
	provider.Register(m)
	m.providers = append(m.providers, provider)
}

func (m *Manager) Register(typ file.Type, provider any) {
	lang, ok := m.languages[typ]
	if !ok {
		m.languages[typ] = Language{
			CompletionProviders:  []CompletionProvider{},
			DiagnosticsProviders: []DiagnosticProvider{},
			DefinitionProviders:  []DefinitionProvider{},
			HoverProviders:       []HoverProvider{},
		}
		lang = m.languages[typ]
	}

	if completion, ok := provider.(CompletionProvider); ok {
		lang.CompletionProviders = append(lang.CompletionProviders, completion)
	}

	if diagnostic, ok := provider.(DiagnosticProvider); ok {
		lang.DiagnosticsProviders = append(lang.DiagnosticsProviders, diagnostic)
	}

	if definition, ok := provider.(DefinitionProvider); ok {
		lang.DefinitionProviders = append(lang.DefinitionProviders, definition)
	}

	if hover, ok := provider.(HoverProvider); ok {
		lang.HoverProviders = append(lang.HoverProviders, hover)
	}

	m.languages[typ] = lang
}

func (m *Manager) Completion(ctx CompletionContext) {
	if providers, ok := m.languages[ctx.File.Type]; ok {
		for _, provider := range providers.CompletionProviders {
			provider.ResolveCompletion(ctx)
		}
	}
}

func (m *Manager) Diagnostics(file *parser.File) []Diagnostic {
	result := []Diagnostic{}

	if providers, ok := m.languages[file.Type]; ok {
		context := DiagnosticContext{
			BaseContext: BaseContext{
				Logger: logrus.WithField("module", "diagnostic"),
				File:   file,
			},
			Publish: func(diagnostic Diagnostic) {
				result = append(result, diagnostic)
			},
		}

		for _, provider := range providers.DiagnosticsProviders {
			provider.Diagnostic(context)
		}
	}

	return result
}

func (m *Manager) ResolveDefinition(context DefinitionContext) {
	if providers, ok := m.languages[context.File.Type]; ok {
		for _, provider := range providers.DefinitionProviders {
			provider.ResolveDefinition(context)
		}
	}
}

func (m *Manager) Hover(context HoverContext) string {
	if providers, ok := m.languages[context.File.Type]; ok {
		for _, provider := range providers.HoverProviders {
			if content := provider.Hover(context); len(content) > 0 {
				return content
			}
		}
	}
	return ""
}
