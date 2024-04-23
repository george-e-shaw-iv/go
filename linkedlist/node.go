package linkedlist

type Node[T comparable] struct {
	Data       T
	Next, Prev *Node[T]
}

func NewNode[T comparable](val T) *Node[T] {
	return &Node[T]{
		Data: val,
	}
}
