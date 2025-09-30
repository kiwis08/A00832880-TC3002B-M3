package main

// Stack represents a Last-In-First-Out (LIFO) data structure
// where elements are added and removed from the top
type Stack interface {
	// Push adds an element to the top of the stack
	Push(value interface{})

	// Pop removes and returns the top element from the stack
	// Returns nil if the stack is empty
	Pop() interface{}

	// Peek returns the top element without removing it
	// Returns nil if the stack is empty
	Peek() interface{}

	// IsEmpty returns true if the stack has no elements
	IsEmpty() bool

	// Size returns the number of elements in the stack
	Size() int

	// Clear removes all elements from the stack
	Clear()
}

// ArrayStack implements Stack using a slice
type ArrayStack struct {
	// TODO: Add fields to store the stack data
	// Hint: You might want to use a slice to store elements
}

// NewArrayStack creates a new empty stack
func NewArrayStack() *ArrayStack {
	// TODO: Initialize and return a new ArrayStack
	return nil
}

// TODO: Implement all Stack interface methods for ArrayStack
// Remember to handle edge cases like empty stack operations
