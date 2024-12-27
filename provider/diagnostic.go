package provider

import (
	"github.com/shufflingpixels/laravel-ls/lsp/protocol"
	ts "github.com/tree-sitter/go-tree-sitter"
)

type Diagnostic struct {
	Range    ts.Range
	Severity protocol.DiagnosticSeverity
	Message  string
}

type DiagnosticPublisher func(Diagnostic)

type DiagnosticContext struct {
	BaseContext
	Publish DiagnosticPublisher
}

// Interface that plugins that supports diagnostics can implement
type DiagnosticProvider interface {
	Diagnostic(DiagnosticContext)
}
