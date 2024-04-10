package linkedlist_test

import (
	"testing"

	"github.com/george-e-shaw-iv/go/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestSingly(t *testing.T) {
	// Test inserting and deleting only from the head on an initially empty list.
	t.Run("MultiInsertAndDeleteFromHeadOnEmptyList", func(t *testing.T) {
		ll := linkedlist.NewSingly[int]()

		ll.InsertFirst(linkedlist.NewNode(1))

		// Head and tail should be equal since only one element has been inserted.
		assert.Equal(t, 1, ll.GetFirst().Data)
		assert.Equal(t, 1, ll.GetLast().Data)

		// Insert another two elements at the front of the list.
		ll.InsertFirst(linkedlist.NewNode(2))
		ll.InsertFirst(linkedlist.NewNode(3))

		// Head should be the newly inserted element, tail should still be the first inserted.
		assert.Equal(t, 3, ll.GetFirst().Data)
		assert.Equal(t, 1, ll.GetLast().Data)

		// They should be linked which can be confirmed by looking at the array representation.
		assert.Equal(t, []int{3, 2, 1}, ll.ToArray())

		// Delete 3.
		ll.DeleteFirst()

		// 2 -> 1 should be the new representnation.
		assert.Equal(t, 2, ll.GetFirst().Data)
		assert.Equal(t, 1, ll.GetLast().Data)
		assert.Equal(t, []int{2, 1}, ll.ToArray())

		// Delete 2.
		ll.DeleteFirst()

		// Head and tail should be equal again now that only one element exists in the list.
		assert.Equal(t, 1, ll.GetFirst().Data)
		assert.Equal(t, 1, ll.GetLast().Data)
		assert.Equal(t, []int{1}, ll.ToArray())

		// Delete the last element in the list from the front.
		ll.DeleteFirst()

		// Nothing should be left in the list.
		assert.Nil(t, ll.GetFirst())
		assert.Nil(t, ll.GetLast())
		assert.Zero(t, len(ll.ToArray()))
	})
}
