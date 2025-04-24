package util

import "errors"

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{make([]T, 0)}
}

func (stack *Stack[T]) Push(element T) {
	stack.elements = append(stack.elements, element)
}

func (stack *Stack[T]) Pop() (T, error) {
	var zero T
	currentSize := len(stack.elements)
	if currentSize == 0 {
		return zero, errors.New("Stack is empty. Cannot pop.")
	}
	element := stack.elements[currentSize-1]
	stack.elements = stack.elements[0 : currentSize-1]
	return element, nil
}

func (stack *Stack[T]) Peek() (T, error) {
	var zero T
	currentSize := len(stack.elements)
	if currentSize == 0 {
		return zero, errors.New("Stack is empty. Cannot peek.")
	}
	return stack.elements[currentSize-1], nil
}

func (stack *Stack[T]) Len() int {
	return len(stack.elements)
}
