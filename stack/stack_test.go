package stack_test

import (
	"testing"

	"github.com/george-e-shaw-iv/go/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	tt := []struct {
		Name           string
		Implementation stack.Stack[int]
	}{
		{
			Name:           "Classic",
			Implementation: stack.NewClassic[int](),
		},
		{
			Name:           "QueueBased",
			Implementation: stack.NewQueueBased[int](),
		},
	}

	for _, test := range tt {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			test.Implementation.Push(0)
			test.Implementation.Push(1)
			test.Implementation.Push(2)

			assert.Equal(t, 3, test.Implementation.Len())
			assert.Equal(t, 2, test.Implementation.Top())

			test.Implementation.Pop()
			test.Implementation.Pop()

			assert.Equal(t, 1, test.Implementation.Len())
			assert.Equal(t, 0, test.Implementation.Top())

			test.Implementation.Push(3)

			assert.Equal(t, 2, test.Implementation.Len())
			assert.Equal(t, 3, test.Implementation.Top())

			test.Implementation.Pop()
			test.Implementation.Pop()

			assert.Equal(t, 0, test.Implementation.Len())
		})
	}
}
