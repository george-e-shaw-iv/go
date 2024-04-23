package linkedlist

import "errors"

var ErrTargetNoExist = errors.New("target does not exist")

// LinkedList is the interface that all linked list implementation should implement.
type LinkedList[T comparable] interface {
	GetFirst() *Node[T]
	GetLast() *Node[T]
	InsertFirst(nodes ...*Node[T])
	InsertLast(nodes ...*Node[T])
	InsertAfter(target T, nodes ...*Node[T]) error
	DeleteFirst()
	DeleteLast()
	Delete(target T)
	Size() int
	Search(target T) *Node[T]
	Clear()
	ToArray() []T
}

var _ LinkedList[int] = &Classic[int]{}

// Classic is a linked list that only tracks the head and is singly linked.
type Classic[T comparable] struct {
	head *Node[T]
	size int
}

func NewClassic[T comparable](nodes ...*Node[T]) *Classic[T] {
	var ll Classic[T]
	ll.InsertLast(nodes...)
	return &ll
}

func (ll *Classic[T]) InsertFirst(nodes ...*Node[T]) {
	for i := range nodes {
		nodes[i].Next = ll.head
		ll.head = nodes[i]
		ll.size++
	}
}

func (ll *Classic[T]) GetFirst() *Node[T] {
	return ll.head
}

func (ll *Classic[T]) GetLast() *Node[T] {
	if ll.head == nil {
		return nil
	}

	var cur *Node[T]
	for cur = ll.head; cur.Next != nil; cur = cur.Next {
		// This works because we're iterating until there is no node pointing next.
		continue
	}
	return cur
}

func (ll *Classic[T]) InsertLast(nodes ...*Node[T]) {
	if ll.Size() == 0 && len(nodes) != 0 {
		ll.InsertFirst(nodes[0])
		nodes = nodes[1:]
	}

	last := ll.GetLast()
	for i := range nodes {
		last.Next = nodes[i]
		last = nodes[i]
		ll.size++
	}
}

func (ll *Classic[T]) Search(target T) *Node[T] {
	for cur := ll.head; cur != nil; cur = cur.Next {
		if cur.Data == target {
			return cur
		}
	}
	return nil
}

func (ll *Classic[T]) InsertAfter(target T, nodes ...*Node[T]) error {
	after := ll.Search(target)
	if after == nil {
		return ErrTargetNoExist
	}

	for i := range nodes {
		nxt := after.Next
		after.Next = nodes[i]
		after = nodes[i]
		after.Next = nxt
		ll.size++
	}

	return nil
}

func (ll *Classic[T]) Size() int {
	return ll.size
}

func (ll *Classic[T]) Delete(target T) {
	if ll.Size() == 0 {
		return
	}

	prev := ll.head
	if prev.Data == target {
		ll.head = prev.Next
		ll.size--
	}

	for cur := prev.Next; cur != nil; cur = cur.Next {
		if cur.Data == target {
			prev.Next = cur.Next
			ll.size--
			return
		}
	}
}

func (ll *Classic[T]) DeleteFirst() {
	if ll.Size() == 0 {
		return
	}

	ll.head = ll.head.Next
	ll.size--
}

func (ll *Classic[T]) DeleteLast() {
	if ll.Size() <= 1 {
		// There is either 0 or 1 nodes in the list, either way the end state is the same.
		// The logic outside of this conditional relies on there being at least two nodes in
		// the list.
		ll.Clear()
		return
	}

	var prev, cur *Node[T]
	for cur = ll.head; cur.Next != nil; cur = cur.Next {
		prev = cur
		continue
	}
	prev.Next = nil
	ll.size--
}

func (ll *Classic[T]) Clear() {
	ll.head = nil
	ll.size = 0
}

func (ll *Classic[T]) ToArray() []T {
	var res []T
	for cur := ll.head; cur != nil; cur = cur.Next {
		res = append(res, cur.Data)
	}
	return res
}
