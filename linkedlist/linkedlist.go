package linkedlist

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var ErrTargetNoExist = errors.New("target does not exist")

type Node[T constraints.Ordered] struct {
	Data       T
	next, prev *Node[T]
}

func NewNode[T constraints.Ordered](val T) *Node[T] {
	return &Node[T]{
		Data: val,
	}
}

// Singly is a singly linked list (only next pointers, not prev pointers).
type Singly[T constraints.Ordered] struct {
	head, tail *Node[T]
}

// NewSingly returns a singly linked list prepopulated with any nodes passed as
// arguments (linked in order of parameter index).
func NewSingly[T constraints.Ordered](nodes ...*Node[T]) *Singly[T] {
	ll := Singly[T]{}
	ll.InsertLast(nodes...)
	return &ll
}

func (ll *Singly[T]) GetFirst() *Node[T] {
	return ll.head
}

func (ll *Singly[T]) GetLast() *Node[T] {
	return ll.tail
}

func (ll *Singly[T]) Search(target T) *Node[T] {
	for cur := ll.head; cur != nil; {
		if cur.Data == target {
			return cur
		}
		cur = cur.next
	}

	return nil
}

func (ll *Singly[T]) ToArray() []T {
	var arr []T
	for cur := ll.head; cur != nil; cur = cur.next {
		arr = append(arr, cur.Data)
	}
	return arr
}

func (ll *Singly[T]) InsertFirst(nodes ...*Node[T]) {
	for i := range nodes {
		nodes[i].next = ll.head
		ll.head = nodes[i]

		// This is the "seed node" case. If the list is empty, head and tail will be the same.
		if ll.tail == nil {
			ll.tail = nodes[i]
		}
	}
}

func (ll *Singly[T]) InsertLast(nodes ...*Node[T]) {
	for i := range nodes {
		ll.tail.next = nodes[i]
		ll.tail = nodes[i]

		// This is the "seed node" case. If the list is empty, head and tail will be the same.
		if ll.head == nil {
			ll.head = nodes[i]
		}
	}
}

func (ll *Singly[T]) InsertAfter(target T, nodes ...*Node[T]) error {
	after := ll.Search(target)
	if after == nil {
		return ErrTargetNoExist
	}

	rightSide := after.next
	for i := range nodes {
		after.next = nodes[i]
		after = nodes[i]
	}
	after.next = rightSide

	return nil
}

func (ll *Singly[T]) DeleteFirst() {
	if ll.head == nil {
		return
	}

	// This is the "one element" case.
	if ll.head == ll.tail {
		ll.head = nil
		ll.tail = nil
		return
	}

	nxt := ll.head.next
	ll.head = nxt
}

func (ll *Singly[T]) DeleteLast() {
	if ll.tail == nil {
		return
	}

	// This is the "one element" case.
	if ll.head == ll.tail {
		ll.head = nil
		ll.tail = nil
		return
	}

	var prev *Node[T]

	// We have an implicit guarantee that the first cur from ll.head will be non-nil
	// because we've already checked if the tail is non-nil above. There should never
	// be a case where the tail is non-nil and the head is
	for cur := ll.head; cur.next != nil; {
		nxt := cur.next
		prev = cur
		cur = nxt
	}

	// cur should be pointing at the last node, prev should be pointing at second to last
	prev.next = nil
}

func (ll *Singly[T]) Delete(target T) {

}
