package repository

// ConfigEntry holds information for an entry for a single
// configuration key,value pair in a laravel application.
type ConfigEntry struct {
	// Value holds the actual value.
	Value any `json:"value"`

	// File specifies the file path where the configuration entry is defined.
	File string `json:"file"`

	// Line indicates the line number in the file where the configuration entry is defined.
	Line int `json:"line"`
}

// ConfigRepository is a type alias for Repository specialized for ConfigEntry.
type ConfigRepository = Repository[ConfigEntry]
