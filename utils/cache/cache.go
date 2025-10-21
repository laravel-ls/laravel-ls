package cache

// Cache is a simple generic datatype that caches key-value pairs.
type Cache[T comparable] struct {
	items map[string]T
}

func New[T comparable]() *Cache[T] {
	return &Cache[T]{
		items: map[string]T{},
	}
}

// Get a value from the cache by its key.
// It returns the value and a boolean indicating
// if the key was found in the cache.
func (c Cache[T]) Get(key string) (T, bool) {
	v, hit := c.items[key]
	return v, hit
}

// Set a value in the cache.
// If the key already exists, its value is updated.
func (c *Cache[T]) Set(key string, value T) {
	c.items[key] = value
}

// Remember gets a value from the cache or generate the value using the provided
// callback if the key is not found. The callback function takes the
// key as input and returns a value and an error. If the callback
// succeeds (non-nil error), the value is cached and returned.
func (c *Cache[T]) Remember(key string, callback func(key string) (T, error)) (T, error) {
	if value, hit := c.Get(key); hit {
		return value, nil
	}

	value, err := callback(key)
	if err == nil {
		c.Set(key, value)
	}
	return value, err
}

func (c *Cache[T]) Forget(key string) {
	delete(c.items, key)
}

// Items returns the entire map of cached items.
func (c Cache[T]) Items() map[string]T {
	return c.items
}

// Clear removes all items from the cache,
// resetting it to an empty state.
func (c *Cache[T]) Clear() {
	c.items = map[string]T{}
}
