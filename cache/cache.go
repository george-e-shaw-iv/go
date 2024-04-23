package cache

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// ErrNotInCache denotes that the provided key was not found in the cache.
var ErrNotInCache = errors.New("key not in cache")

type Cache[Key constraints.Ordered, Value comparable] interface {
	Get(k Key) (Value, error)
	Put(k Key, v Value)
}
