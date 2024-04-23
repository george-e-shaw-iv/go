package linkedlist

var _ LinkedList[int] = &DoubleHeadTail[int]{}

type DoubleHeadTail[T comparable] struct {
	head, tail *Node[T]
	size       int
}

func NewDoubleHeadTail[T comparable](nodes ...*Node[T]) *DoubleHeadTail[T] {
	var d DoubleHeadTail[T]
	d.InsertLast(nodes...)
	return &d
}

func (d *DoubleHeadTail[T]) Clear() {
	d.head = nil
	d.tail = nil
	d.size = 0
}

func (d *DoubleHeadTail[T]) Delete(target T) {
	n := d.Search(target)
	if n == nil {
		return
	}

	if d.size == 1 {
		d.Clear()
		return
	}

	prev := n.Prev
	next := n.Next

	if prev != nil {
		prev.Next = next
	} else {
		// We just deleted the head
		d.head = next
		d.head.Prev = nil
	}

	if next != nil {
		next.Prev = prev
	} else {
		// We just deleted the tail
		d.tail = prev
		d.tail.Next = nil
	}

	d.size--
}

func (d *DoubleHeadTail[T]) DeleteFirst() {
	if d.head == nil {
		return
	}

	if d.size == 1 {
		d.Clear()
		return
	}

	d.head = d.head.Next
	d.head.Prev = nil
	d.size--
}

func (d *DoubleHeadTail[T]) DeleteLast() {
	if d.tail == nil {
		return
	}

	if d.size == 1 {
		d.Clear()
		return
	}

	d.tail = d.tail.Prev
	d.tail.Next = nil
	d.size--
}

func (d *DoubleHeadTail[T]) GetFirst() *Node[T] {
	return d.head
}

func (d *DoubleHeadTail[T]) GetLast() *Node[T] {
	return d.tail
}

func (d *DoubleHeadTail[T]) InsertAfter(target T, nodes ...*Node[T]) error {
	if len(nodes) == 0 {
		return nil
	}

	n := d.Search(target)
	if n == nil {
		return ErrTargetNoExist
	}

	last := n.Next
	for i := range nodes {
		n.Next = nodes[i]
		nodes[i].Prev = n
		n = nodes[i]
		d.size++
	}

	nodes[len(nodes)-1].Next = last
	if last != nil {
		last.Prev = nodes[len(nodes)-1]
	}

	return nil
}

func (d *DoubleHeadTail[T]) InsertFirst(nodes ...*Node[T]) {
	for i := range nodes {
		if d.size == 0 {
			d.head = nodes[i]
			d.tail = nodes[i]
			nodes[i].Next = nil
			nodes[i].Prev = nil
			d.size++
			continue
		}

		oldHead := d.head
		d.head = nodes[i]
		d.head.Prev = nil
		d.head.Next = oldHead
		oldHead.Prev = nodes[i]
		d.size++
	}
}

func (d *DoubleHeadTail[T]) InsertLast(nodes ...*Node[T]) {
	for i := range nodes {
		if d.size == 0 {
			d.head = nodes[i]
			d.tail = nodes[i]
			nodes[i].Next = nil
			nodes[i].Prev = nil
			d.size++
			continue
		}

		oldTail := d.tail
		oldTail.Next = nodes[i]
		d.tail = nodes[i]
		d.tail.Next = nil
		d.tail.Prev = oldTail
		d.size++
	}
}

func (d *DoubleHeadTail[T]) Search(target T) *Node[T] {
	for cur := d.head; cur != nil; cur = cur.Next {
		if cur.Data == target {
			return cur
		}
	}
	return nil
}

func (d *DoubleHeadTail[T]) Size() int {
	return d.size
}

func (d *DoubleHeadTail[T]) ToArray() []T {
	var res []T
	for cur := d.head; cur != nil; cur = cur.Next {
		res = append(res, cur.Data)
	}
	return res
}
