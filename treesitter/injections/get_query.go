package injections

import (
	_ "embed"

	"laravel-ls/treesitter"
)

//go:embed queries/php.scm
var phpQuery string

//go:embed queries/blade.scm
var bladeQuery string

// Get the injection query based on language
func GetQuery(language string) string {
	switch language {
	case treesitter.LanguagePhp:
		return phpQuery
	case treesitter.LanguageBlade:
		return bladeQuery
	}
	return ""
}
