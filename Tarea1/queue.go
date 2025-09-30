package main

// Queue represents a First-In-First-Out (FIFO) data structure
// where elements are added to the rear and removed from the front
type Queue interface {
	// Enqueue adds an element to the rear of the queue
	Enqueue(value interface{})

	// Dequeue removes and returns the front element from the queue
	// Returns nil if the queue is empty
	Dequeue() interface{}

	// Front returns the front element without removing it
	// Returns nil if the queue is empty
	Front() interface{}

	// IsEmpty returns true if the queue has no elements
	IsEmpty() bool

	// Size returns the number of elements in the queue
	Size() int

	// Clear removes all elements from the queue
	Clear()
}

// ArrayQueue implements Queue using a slice
type ArrayQueue struct {
	// TODO: Add fields to store the queue data
	// Hint: You might want to use a slice to store elements
	// Consider how to efficiently implement FIFO with a slice
}

// NewArrayQueue creates a new empty queue
func NewArrayQueue() *ArrayQueue {
	// TODO: Initialize and return a new ArrayQueue
	return nil
}

// TODO: Implement all Queue interface methods for ArrayQueue
// Remember to handle edge cases like empty queue operations
// Consider the efficiency of your implementation
