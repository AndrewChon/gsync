package gsync

import (
	"sync"
)

// Map is a generic wrapper for sync.Map. K and V must implement comparable (see CompareAndDelete and CompareAndSwap).
type Map[K comparable, V comparable] struct {
	m sync.Map
}

// Clear deletes all the entries in the Map, resulting in an empty Map.
func (gm *Map[K, V]) Clear() {
	gm.m.Clear()
}

// CompareAndDelete deletes the entry for a key if its value is equal to the old value.
func (gm *Map[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return gm.m.CompareAndDelete(key, old)
}

// CompareAndSwap swaps the old and new values for a key if the value stored in the map is equal to the old value.
func (gm *Map[K, V]) CompareAndSwap(key K, old, new V) (swapped bool) {
	return gm.m.CompareAndSwap(key, old, new)
}

// Delete deletes the value for a key.
func (gm *Map[K, V]) Delete(key K) {
	gm.m.Delete(key)
}

// Load returns the value stored in the map for a key or nil if no value is present. The ok result indicates whether the
// value was found in the Map.
func (gm *Map[K, V]) Load(key K) (value V, ok bool) {
	v, ok := gm.m.Load(key)
	if !ok {
		return
	}

	value = v.(V)
	return
}

// LoadAndDelete deletes the value for a key, returning the previous value, if any. The loaded result indicates whether
// the key was present in the Map.
func (gm *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := gm.m.LoadAndDelete(key)
	if !loaded {
		return
	}

	value = v.(V)
	return
}

// LoadOrStore returns the existing value for the key if present. If no value exists, it stores and returns the provided
// value. The loaded result indicates whether the value was loaded.
func (gm *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	v, loaded := gm.m.LoadOrStore(key, value)
	if !loaded {
		actual = value
		return
	}

	actual = v.(V)
	return
}

// Range calls f sequentially for each key and value present in the Map. If f returns false, Range stops iterating over
// the Map.
func (gm *Map[K, V]) Range(f func(key K, value V) bool) {
	gm.m.Range(func(key, value interface{}) bool {
		return f(key.(K), value.(V))
	})
}

// Store sets the value for a key.
func (gm *Map[K, V]) Store(key K, value V) {
	gm.m.Store(key, value)
}

// Swap swaps the value for a key and returns the previous value, if any. The loaded result indicates whether the key
// was present in the Map.
func (gm *Map[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	pv, loaded := gm.m.Swap(key, value)
	if !loaded {
		return
	}

	previous = pv.(V)
	return
}
