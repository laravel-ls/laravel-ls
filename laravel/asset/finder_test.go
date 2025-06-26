//go:build !windows

package asset_test

import (
	"testing"

	"github.com/laravel-ls/laravel-ls/laravel/asset"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func Test_Search(t *testing.T) {
	fs := afero.NewMemMapFs()
	fs.Create("/var/www/project/public/css/app.css")
	fs.Create("/var/www/project/public/js/app.js")
	fs.Create("/var/www/project/public/image/logo.png")

	finder := asset.NewFinder(fs, "/var/www/project")
	result := finder.Search("app")
	require.Len(t, result, 2)
	require.Equal(t, "/var/www/project/public/css/app.css", result[0])
	require.Equal(t, "/var/www/project/public/js/app.js", result[1])
}

func Test_Exists(t *testing.T) {
	fs := afero.NewMemMapFs()
	fs.Create("/var/www/project/public/css/app.css")
	fs.Create("/var/www/project/public/js/app.js")
	fs.Create("/var/www/project/public/image/logo.png")

	finder := asset.NewFinder(fs, "/var/www/project")
	require.True(t, finder.Exists("css/app.css"))
	require.True(t, finder.Exists("js/app.js"))
	require.False(t, finder.Exists("robots.txt"))
}
