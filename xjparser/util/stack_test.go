package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := NewStack[string]()
	assert.Equal(t, 0, stack.Len())

	stack.Push("value1")
	stack.Push("value2")
	stack.Push("value3")
	assert.Equal(t, 3, stack.Len())

	str, err := stack.Peek()
	assert.Nil(t, err)
	assert.Equal(t, "value3", str)
	assert.Equal(t, 3, stack.Len())
	str, err = stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, "value3", str)
	assert.Equal(t, 2, stack.Len())

	str, err = stack.Peek()
	assert.Nil(t, err)
	assert.Equal(t, "value2", str)
	assert.Equal(t, 2, stack.Len())
	str, err = stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, "value2", str)
	assert.Equal(t, 1, stack.Len())

	str, err = stack.Peek()
	assert.Nil(t, err)
	assert.Equal(t, "value1", str)
	assert.Equal(t, 1, stack.Len())
	str, err = stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, "value1", str)
	assert.Equal(t, 0, stack.Len())

	_, err = stack.Peek()
	assert.NotNil(t, err)
	_, err = stack.Pop()
	assert.NotNil(t, err)
}

func TestEmptyStack(t *testing.T) {
	stack := NewStack[string]()
	_, err := stack.Peek()
	assert.NotNil(t, err)
	_, err = stack.Pop()
	assert.NotNil(t, err)
}
