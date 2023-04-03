package cache

import "time"

type ICache interface {
	// Set a value with a given key and expiration time
	Set(key string, value any, expiration time.Duration) error

	// Get a value with a given key
	Get(key string) (any, error)

	// Delete a value with a given key
	Delete(key string) error

	// Flush all values from the cache
	Flush() error
}
