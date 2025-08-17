package db

import "sync"

var defaultStore *Store
var once sync.Once

// Default returns the singleton in-memory store.
func Default() *Store {
	once.Do(func() {
		defaultStore = New()
		defaultStore.SeedDefault()
	})
	return defaultStore
}
