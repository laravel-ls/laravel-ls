package repository

// AppEntry represents an entry for an application binding in laravel.
// It holds static information about where a binding is defined in the application
type AppEntry struct {
	// Class specifies what class this binding is bound to.
	Class string `json:"class"`

	// Path specifies the file path where the binding is defined.
	Path string `json:"path"`

	// Line indicates the line number in the file where the app entry is declared.
	Line int `json:"line"`
}

// AppRepository is a type alias for Repository specialized for AppEntry.
type AppRepository = Repository[AppEntry]
