package repository

// RouteEntry holds information for an entry for a single
// route configuration in a laravel application.
type RouteEntry struct {
	Method     string   `json:"method"`     // HTTP method (GET, POST, etc.)
	URI        string   `json:"uri"`        // URI path for the route
	Name       string   `json:"name"`       // Name of the route, if defined
	Action     string   `json:"action"`     // Action or controller handling the route
	Parameters []string `json:"parameters"` // List of parameters for the route
	File       string   `json:"filename"`   // File indicates the file where the route is defined.
	Line       int      `json:"line"`       // Line number in the file where the route is defined
}

// RouteRepository is a type alias for Repository specialized for RouteEntry.
type RouteRepository = Repository[RouteEntry]
