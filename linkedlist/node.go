package linkedlist

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	Data       T
	Next, Prev *Node[T]
}

func NewNode[T constraints.Ordered](val T) *Node[T] {
	return &Node[T]{
		Data: val,
	}
}
