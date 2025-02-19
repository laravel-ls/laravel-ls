package repository

import "strings"

// Repository is a generic type that maps string keys to values of any type.
// It essentially acts as a collection for storing and retrieving data by key.
type Repository[T any] map[string]T

// Find searches for entries in the repository whose keys start with the provided prefix.
// It returns a new map containing only the matching entries.
//
// Parameters:
//   - input: The prefix string to match keys against.
//
// Returns:
//   - A map of keys and their corresponding values that match the prefix.
func (r Repository[T]) Find(input string) map[string]T {
	result := map[string]T{}
	for k, v := range r {
		if strings.HasPrefix(k, input) {
			result[k] = v
		}
	}
	return result
}

// Get retrieves the value associated with the given key.
//
// Parameters:
//   - key: The key whose associated value is to be retrieved.
//
// Returns:
//   - value: The value corresponding to the key (zero value if not found).
//   - found: A boolean indicating whether the key was found in the repository.
func (r Repository[T]) Get(key string) (value T, found bool) {
	value, found = r[key]
	return
}

// Exists checks if a specific key exists within the repository.
//
// Parameters:
//   - key: The key to check for existence.
//
// Returns:
//   - A boolean indicating whether the key exists.
func (r Repository[T]) Exists(key string) bool {
	_, found := r.Get(key)
	return found
}

// Clear removes all entries from the repository, resetting it to an empty state.
func (r *Repository[T]) Clear() {
	*r = map[string]T{}
}
