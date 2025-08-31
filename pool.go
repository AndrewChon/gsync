package gsync

import (
	"sync"
)

// A Pool is a set of temporary objects that may be individually saved and retrieved.
type Pool[V any] struct {
	p   sync.Pool
	New func() V
}

// Get selects an arbitrary item from the Pool, removes it from the Pool, and returns it to the caller. If the Pool is
// empty and sp.New is defined, Get returns the result of calling sp.New.
func (sp *Pool[V]) Get() V {
	v := sp.p.Get()
	if v != nil {
		return v.(V)
	}

	if sp.New != nil {
		return sp.New()
	}

	var zero V
	return zero
}

// Put adds v to the Pool.
func (sp *Pool[V]) Put(v V) {
	sp.p.Put(v)
}
