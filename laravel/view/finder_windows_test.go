//go:build windows

package view_test

import (
	"testing"

	"github.com/laravel-ls/laravel-ls/laravel/view"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func Test_Find(t *testing.T) {
	fs := afero.NewMemMapFs()
	fs.Create("D:\\code\\www\\project\\resources\\views\\index.blade.php")
	fs.Create("D:\\code\\www\\project\\resources\\views\\partials/header.php")
	fs.Create("D:\\code\\www\\project\\other\\module1\\component.php")

	finder := view.NewFinder(fs)
	finder.AddLocation("D:\\code\\www\\project\\resources\\views")
	finder.AddLocation("D:\\code\\www\\project\\other")

	path, found := finder.Find("index")
	require.True(t, found)
	require.Equal(t, "D:\\code\\www\\project\\resources\\views\\index.blade.php", path)

	path, found = finder.Find("partials.header")
	require.True(t, found)
	require.Equal(t, "D:\\code\\www\\project\\resources\\views\\partials\\header.php", path)

	path, found = finder.Find("module1.component")
	require.True(t, found)
	require.Equal(t, "D:\\code\\www\\project\\other\\module1\\component.php", path)

	path, found = finder.Find("does.not.exist")
	require.False(t, found)
	require.Equal(t, "", path)
}

func Test_Search(t *testing.T) {
	fs := afero.NewMemMapFs()
	fs.Create("D:\\code\\www\\project\\resources\\views\\index.blade.php")
	fs.Create("D:\\code\\www\\project\\resources\\views\\partials\\header.php")
	fs.Create("D:\\code\\www\\project\\resources\\views\\contact.php")
	fs.Create("D:\\code\\www\\project\\resources\\views\\components\\card\\index.blade.php")
	fs.Create("D:\\code\\www\\project\\other\\components\\badge\\index.php")
	fs.Create("D:\\code\\www\\project\\other\\module1\\component.php")

	finder := view.NewFinder(fs)
	finder.AddLocation("D:\\code\\www\\project\\resources\\views")
	finder.AddLocation("D:\\code\\www\\project\\other")

	views := finder.Search("components")
	require.Len(t, views, 2)
	require.Equal(t, *view.NewView("D:\\code\\www\\project\\resources\\views\\components\\card\\index.blade.php", "components.card.index"), views[0])
	require.Equal(t, *view.NewView("D:\\code\\www\\project\\other\\components\\badge\\index.php", "components.badge.index"), views[1])
}
