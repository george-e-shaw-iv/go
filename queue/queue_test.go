package queue_test

import (
	"testing"

	"github.com/george-e-shaw-iv/go/queue"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	var q queue.Queue[int]

	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)

	assert.Equal(t, 3, q.Len())
	assert.Equal(t, 0, q.Peek())

	assert.Equal(t, 0, q.Dequeue())
	assert.Equal(t, 1, q.Dequeue())

	assert.Equal(t, 1, q.Len())
	assert.Equal(t, 2, q.Peek())

	q.Enqueue(3)

	assert.Equal(t, 2, q.Len())
	assert.Equal(t, 2, q.Peek())

	assert.Equal(t, 2, q.Dequeue())
	assert.Equal(t, 3, q.Dequeue())

	assert.Equal(t, 0, q.Len())
}
