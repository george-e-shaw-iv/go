package linkedlist_test

import (
	"testing"

	"github.com/george-e-shaw-iv/go/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	tests := []struct {
		Name           string
		Implementation linkedlist.LinkedList[int]
	}{
		{
			Name:           "Classic",
			Implementation: linkedlist.NewClassic[int](),
		},
		{
			Name:           "SingleLinkedTrackHeadAndTail",
			Implementation: linkedlist.NewSingleHeadTail[int](),
		},
		{
			Name:           "DoubleLinkedTrackHeadAndTail",
			Implementation: linkedlist.NewDoubleHeadTail[int](),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			// Test inserting and deleting only from the head on an initially empty list.
			t.Run("MultiInsertAndDeleteFromHeadOnEmptyList", func(t *testing.T) {
				t.Cleanup(test.Implementation.Clear)

				test.Implementation.InsertFirst(linkedlist.NewNode(1))

				// Head and tail should be equal since only one element has been inserted.
				assert.Equal(t, 1, test.Implementation.GetFirst().Data)
				assert.Equal(t, 1, test.Implementation.GetLast().Data)

				// Insert another two elements at the front of the list.
				test.Implementation.InsertFirst(linkedlist.NewNode(2))
				test.Implementation.InsertFirst(linkedlist.NewNode(3))

				// Head should be the newly inserted element, tail should still be the first inserted.
				assert.Equal(t, 3, test.Implementation.GetFirst().Data)
				assert.Equal(t, 1, test.Implementation.GetLast().Data)

				// They should be linked which can be confirmed by looking at the array representation.
				assert.Equal(t, []int{3, 2, 1}, test.Implementation.ToArray())

				// Delete 3.
				test.Implementation.DeleteFirst()

				// 2 -> 1 should be the new representnation.
				assert.Equal(t, 2, test.Implementation.GetFirst().Data)
				assert.Equal(t, 1, test.Implementation.GetLast().Data)
				assert.Equal(t, []int{2, 1}, test.Implementation.ToArray())

				// Delete 2.
				test.Implementation.DeleteFirst()

				// Head and tail should be equal again now that only one element exists in the list.
				assert.Equal(t, 1, test.Implementation.GetFirst().Data)
				assert.Equal(t, 1, test.Implementation.GetLast().Data)
				assert.Equal(t, []int{1}, test.Implementation.ToArray())

				// Delete the last element in the list from the front.
				test.Implementation.DeleteFirst()

				// Nothing should be left in the list.
				assert.Nil(t, test.Implementation.GetFirst())
				assert.Nil(t, test.Implementation.GetLast())
				assert.Zero(t, len(test.Implementation.ToArray()))
			})

			// Test inserting and deleting only from the tail on an initially empty list.
			t.Run("MultiInsertAndDeleteFromTailOnEmptyList", func(t *testing.T) {
				t.Cleanup(test.Implementation.Clear)

				test.Implementation.InsertLast(linkedlist.NewNode(1))

				// Head and tail should be equal since only one element has been inserted.
				assert.Equal(t, 1, test.Implementation.GetFirst().Data)
				assert.Equal(t, 1, test.Implementation.GetLast().Data)

				// Insert another two elements at the back of the list.
				test.Implementation.InsertLast(linkedlist.NewNode(2))
				test.Implementation.InsertLast(linkedlist.NewNode(3))

				// Tail should be the newly inserted element, head should still be the first inserted.
				assert.Equal(t, 1, test.Implementation.GetFirst().Data)
				assert.Equal(t, 3, test.Implementation.GetLast().Data)

				// They should be linked which can be confirmed by looking at the array representation.
				assert.Equal(t, []int{1, 2, 3}, test.Implementation.ToArray())

				// Delete 3.
				test.Implementation.DeleteLast()

				// 1 -> 2 should be the new representnation.
				assert.Equal(t, 1, test.Implementation.GetFirst().Data)
				assert.Equal(t, 2, test.Implementation.GetLast().Data)
				assert.Equal(t, []int{1, 2}, test.Implementation.ToArray())

				// Delete 2.
				test.Implementation.DeleteLast()

				// Head and tail should be equal again now that only one element exists in the list.
				assert.Equal(t, 1, test.Implementation.GetFirst().Data)
				assert.Equal(t, 1, test.Implementation.GetLast().Data)
				assert.Equal(t, []int{1}, test.Implementation.ToArray())

				// Delete the last element in the list.
				test.Implementation.DeleteLast()

				// Nothing should be left in the list.
				assert.Nil(t, test.Implementation.GetFirst())
				assert.Nil(t, test.Implementation.GetLast())
				assert.Zero(t, len(test.Implementation.ToArray()))
			})

			t.Run("InsertAfter", func(t *testing.T) {
				t.Cleanup(test.Implementation.Clear)

				test.Implementation.InsertLast(linkedlist.NewNode(1), linkedlist.NewNode(2), linkedlist.NewNode(4))
				assert.Equal(t, []int{1, 2, 4}, test.Implementation.ToArray())

				// Try to insert after a target that doesn't exist.
				assert.ErrorIs(t, test.Implementation.InsertAfter(6, linkedlist.NewNode(7)), linkedlist.ErrTargetNoExist)

				// Insert a node an ensure it ends up in the right place.
				assert.NoError(t, test.Implementation.InsertAfter(2, linkedlist.NewNode(3)))
				assert.Equal(t, []int{1, 2, 3, 4}, test.Implementation.ToArray())

				// Insert multiple nodes at the end of the list via InsertAfter.
				assert.NoError(t, test.Implementation.InsertAfter(4, linkedlist.NewNode(5), linkedlist.NewNode(6)))
				assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, test.Implementation.ToArray())
			})

			t.Run("DeleteSpecificElements", func(t *testing.T) {
				t.Cleanup(test.Implementation.Clear)

				test.Implementation.InsertLast(linkedlist.NewNode(1), linkedlist.NewNode(2), linkedlist.NewNode(3), linkedlist.NewNode(4))
				assert.Equal(t, []int{1, 2, 3, 4}, test.Implementation.ToArray())

				// Delete 2.
				test.Implementation.Delete(2)
				assert.Equal(t, []int{1, 3, 4}, test.Implementation.ToArray())

				// Delete an element that doesn't exist.
				test.Implementation.Delete(5)
				assert.Equal(t, []int{1, 3, 4}, test.Implementation.ToArray())

				// Delete an element from the front of the list.
				test.Implementation.Delete(1)
				assert.Equal(t, []int{3, 4}, test.Implementation.ToArray())

				// Delete an element from the back of the list.
				test.Implementation.Delete(4)
				assert.Equal(t, []int{3}, test.Implementation.ToArray())

				// Delete the last element from the list.
				test.Implementation.Delete(3)
				assert.Zero(t, len(test.Implementation.ToArray()))
			})
		})
	}
}
