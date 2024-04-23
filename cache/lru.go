package cache

import (
	"sync"

	"github.com/george-e-shaw-iv/go/linkedlist"
	"golang.org/x/exp/constraints"
)

var _ Cache[string, string] = &LRU[string, string]{}

type kv[Key constraints.Ordered, Value comparable] struct {
	key   Key
	value Value
}

type LRU[Key constraints.Ordered, Value comparable] struct {
	ll       *linkedlist.DoubleHeadTail[kv[Key, Value]]
	data     map[Key]*linkedlist.Node[kv[Key, Value]]
	capacity int
	mu       sync.Mutex
}

func NewLRU[Key constraints.Ordered, Value comparable](capacity int) *LRU[Key, Value] {
	return &LRU[Key, Value]{
		ll:       linkedlist.NewDoubleHeadTail[kv[Key, Value]](),
		data:     make(map[Key]*linkedlist.Node[kv[Key, Value]]),
		capacity: capacity,
	}
}

func (l *LRU[Key, Value]) Get(k Key) (Value, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if _, exists := l.data[k]; !exists {
		var v Value
		return v, ErrNotInCache
	}

	// Put entry at the front of the list.
	l.ll.Delete(l.data[k].Data)
	l.ll.InsertFirst(l.data[k])

	return l.data[k].Data.value, nil
}

func (l *LRU[Key, Value]) Put(k Key, v Value) {
	l.mu.Lock()
	defer l.mu.Unlock()

	data := kv[Key, Value]{
		key:   k,
		value: v,
	}

	if n, exists := l.data[k]; exists {
		l.ll.Delete(n.Data) // Delete the existing entry.
		n.Data = data       // Overwrite the data in the existing.
		l.ll.InsertFirst(n) // Put the node with the new data at the head of the list.
		l.data[k] = n       // Save the node with the new data.
		return
	}

	n := linkedlist.NewNode(data)
	l.data[k] = n
	l.ll.InsertFirst(n)

	if l.ll.Size() > l.capacity {
		delete(l.data, l.ll.GetLast().Data.key)
		l.ll.DeleteLast()
	}
}
