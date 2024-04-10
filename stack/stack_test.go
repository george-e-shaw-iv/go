package stack_test

import (
	"testing"

	"github.com/george-e-shaw-iv/go/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	var st stack.Stack[int]

	st.Push(0)
	st.Push(1)
	st.Push(2)

	assert.Equal(t, 3, st.Len())
	assert.Equal(t, 2, st.Top())

	st.Pop()
	st.Pop()

	assert.Equal(t, 1, st.Len())
	assert.Equal(t, 0, st.Top())

	st.Push(3)

	assert.Equal(t, 2, st.Len())
	assert.Equal(t, 3, st.Top())

	st.Pop()
	st.Pop()

	assert.Equal(t, 0, st.Len())
}
