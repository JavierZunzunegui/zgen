package zgen

import (
	"sync"
)

// SyncMap is a type-safe wrapper around sync.Map.
// The zero value is invalid, use [NewSyncMap].
type SyncMap[K comparable, V any] struct {
	m *sync.Map
}

func NewSyncMap[K comparable, V any]() SyncMap[K, V] {
	return SyncMap[K, V]{m: &sync.Map{}}
}

// Store adds a key/value pair
func (sm SyncMap[K, V]) Store(key K, value V) {
	sm.m.Store(key, value)
}

// Load retrieves a value for a key
func (sm SyncMap[K, V]) Load(key K) (V, bool) {
	v, ok := sm.m.Load(key)
	if !ok {
		var zero V
		return zero, false
	}
	return v.(V), true
}

// Delete removes a key
func (sm SyncMap[K, V]) Delete(key K) {
	sm.m.Delete(key)
}

// LoadOrStore returns the existing value if present, otherwise stores and returns the new one
func (sm SyncMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	actual, loaded := sm.m.LoadOrStore(key, value)
	return actual.(V), loaded
}

// Range iterates over the map
func (sm SyncMap[K, V]) Range(f func(key K, value V) bool) {
	sm.m.Range(func(k, v any) bool {
		return f(k.(K), v.(V))
	})
}
