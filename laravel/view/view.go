package view

import (
	pathpkg "path"
)

// Struct to represent a view file.
type View struct {
	// View name
	name string

	// Path to the view file
	path string
}

func NewView(path, name string) *View {
	view := &View{
		name: name,
	}
	return view.SetPath(path)
}

func (v View) Name() string {
	return v.name
}

func (v *View) SetPath(path string) *View {
	v.path = pathpkg.Clean(path)
	return v
}

func (v View) Path() string {
	return v.path
}
