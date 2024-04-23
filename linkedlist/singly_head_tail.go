package linkedlist

var _ LinkedList[int] = &SingleHeadTail[int]{}

// SingleHeadTail is a single linked list that tracks both the head and the tail.
type SingleHeadTail[T comparable] struct {
	head, tail *Node[T]
	size       int
}

// NewSingleHeadTail returns a single linked list prepopulated with any nodes passed as
// arguments (linked in order of parameter index).
func NewSingleHeadTail[T comparable](nodes ...*Node[T]) *SingleHeadTail[T] {
	ll := SingleHeadTail[T]{}
	ll.InsertLast(nodes...)
	return &ll
}

func (ll *SingleHeadTail[T]) GetFirst() *Node[T] {
	return ll.head
}

func (ll *SingleHeadTail[T]) GetLast() *Node[T] {
	return ll.tail
}

func (ll *SingleHeadTail[T]) Search(target T) *Node[T] {
	for cur := ll.head; cur != nil; {
		if cur.Data == target {
			return cur
		}
		cur = cur.Next
	}

	return nil
}

func (ll *SingleHeadTail[T]) ToArray() []T {
	var arr []T
	for cur := ll.head; cur != nil; cur = cur.Next {
		arr = append(arr, cur.Data)
	}
	return arr
}

func (ll *SingleHeadTail[T]) InsertFirst(nodes ...*Node[T]) {
	for i := range nodes {
		nodes[i].Next = ll.head
		ll.head = nodes[i]

		// This is the "seed node" case. If the list is empty, head and tail will be the same.
		if ll.Size() == 0 {
			ll.tail = nodes[i]
		}
		ll.size++
	}
}

func (ll *SingleHeadTail[T]) InsertLast(nodes ...*Node[T]) {
	for i := range nodes {
		// This is the "seed node" case. If the list is empty, head and tail will be the same, which
		// InsertFirst already takes care of.
		if ll.Size() == 0 {
			ll.InsertFirst(nodes[i])
			continue
		}

		ll.tail.Next = nodes[i]
		ll.tail = nodes[i]
		ll.size++
	}
}

func (ll *SingleHeadTail[T]) InsertAfter(target T, nodes ...*Node[T]) error {
	after := ll.Search(target)
	if after == nil {
		return ErrTargetNoExist
	}

	ll.size++

	rightSide := after.Next
	for i := range nodes {
		after.Next = nodes[i]
		after = nodes[i]
	}
	after.Next = rightSide

	return nil
}

func (ll *SingleHeadTail[T]) DeleteFirst() {
	if ll.Size() == 0 {
		return
	}

	defer func() {
		ll.size--
	}()

	// This is the "one element" case.
	if ll.Size() == 1 {
		ll.head = nil
		ll.tail = nil
		return
	}

	nxt := ll.head.Next
	ll.head = nxt
}

func (ll *SingleHeadTail[T]) DeleteLast() {
	if ll.Size() == 0 {
		return
	}

	defer func() {
		ll.size--
	}()

	// They're both pointing to the same node.
	if ll.Size() == 1 {
		ll.head = nil
		ll.tail = nil
		return
	}

	var prev *Node[T]

	// We have an implicit guarantee that the first cur from ll.head will be non-nil
	// because we've already checked if the tail is non-nil above. There should never
	// be a case where the tail is non-nil and the head is
	for cur := ll.head; cur.Next != nil; {
		nxt := cur.Next
		prev = cur
		cur = nxt
	}

	// cur should be pointing at the last node, prev should be pointing at second to last
	prev.Next = nil
	ll.tail = prev
}

func (ll *SingleHeadTail[T]) Size() int {
	return ll.size
}

func (ll *SingleHeadTail[T]) Delete(target T) {
	if ll.Size() == 0 {
		return
	}

	if ll.head.Data == target {
		ll.DeleteFirst()
		return
	}

	prev := ll.head
	for cur := prev.Next; cur != nil; cur = cur.Next {
		if cur.Data == target {
			prev.Next = cur.Next
			ll.size--
			return
		}
	}
}

func (ll *SingleHeadTail[T]) Clear() {
	ll.head = nil
	ll.tail = nil
	ll.size = 0
}
