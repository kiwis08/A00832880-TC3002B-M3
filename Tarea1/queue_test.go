package main

import (
	"testing"
)

func TestArrayQueue(t *testing.T) {
	queue := NewArrayQueue()

	// Test initial state
	t.Run("Initial state", func(t *testing.T) {
		if !queue.IsEmpty() {
			t.Error("New queue should be empty")
		}
		if queue.Size() != 0 {
			t.Error("New queue should have size 0")
		}
		if queue.Front() != nil {
			t.Error("Front on empty queue should return nil")
		}
		if queue.Dequeue() != nil {
			t.Error("Dequeue on empty queue should return nil")
		}
	})

	// Test enqueue and front
	t.Run("Enqueue and Front", func(t *testing.T) {
		queue.Enqueue(1)
		if queue.IsEmpty() {
			t.Error("Queue should not be empty after enqueue")
		}
		if queue.Size() != 1 {
			t.Error("Queue size should be 1 after one enqueue")
		}
		if queue.Front() != 1 {
			t.Error("Front should return the first enqueued element")
		}

		queue.Enqueue(2)
		if queue.Size() != 2 {
			t.Error("Queue size should be 2 after two enqueues")
		}
		if queue.Front() != 1 {
			t.Error("Front should return the first enqueued element (FIFO)")
		}
	})

	// Test dequeue
	t.Run("Dequeue", func(t *testing.T) {
		// Queue should have [1, 2] with 1 at front
		dequeued := queue.Dequeue()
		if dequeued != 1 {
			t.Error("Dequeue should return the first enqueued element")
		}
		if queue.Size() != 1 {
			t.Error("Queue size should be 1 after one dequeue")
		}
		if queue.Front() != 2 {
			t.Error("Front should return 2 after dequeuing 1")
		}

		dequeued = queue.Dequeue()
		if dequeued != 2 {
			t.Error("Dequeue should return 2")
		}
		if !queue.IsEmpty() {
			t.Error("Queue should be empty after dequeuing all elements")
		}
		if queue.Size() != 0 {
			t.Error("Queue size should be 0 after dequeuing all elements")
		}
	})

	// Test multiple operations
	t.Run("Multiple operations", func(t *testing.T) {
		// Enqueue multiple elements
		elements := []interface{}{"a", "b", "c", "d", "e"}
		for _, elem := range elements {
			queue.Enqueue(elem)
		}

		// Verify FIFO order
		for i := 0; i < len(elements); i++ {
			if queue.Front() != elements[i] {
				t.Errorf("Expected %v, got %v", elements[i], queue.Front())
			}
			dequeued := queue.Dequeue()
			if dequeued != elements[i] {
				t.Errorf("Expected %v, got %v", elements[i], dequeued)
			}
		}

		if !queue.IsEmpty() {
			t.Error("Queue should be empty after dequeuing all elements")
		}
	})

	// Test clear
	t.Run("Clear", func(t *testing.T) {
		// Add some elements
		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)

		queue.Clear()
		if !queue.IsEmpty() {
			t.Error("Queue should be empty after clear")
		}
		if queue.Size() != 0 {
			t.Error("Queue size should be 0 after clear")
		}
		if queue.Front() != nil {
			t.Error("Front should return nil after clear")
		}
	})

	// Test with different data types
	t.Run("Different data types", func(t *testing.T) {
		queue.Enqueue("string")
		queue.Enqueue(42)
		queue.Enqueue(3.14)
		queue.Enqueue(true)

		if queue.Dequeue() != "string" {
			t.Error("Should handle string values")
		}
		if queue.Dequeue() != 42 {
			t.Error("Should handle integer values")
		}
		if queue.Dequeue() != 3.14 {
			t.Error("Should handle float values")
		}
		if queue.Dequeue() != true {
			t.Error("Should handle boolean values")
		}
	})

	// Test alternating enqueue/dequeue
	t.Run("Alternating operations", func(t *testing.T) {
		queue.Enqueue(1)
		queue.Enqueue(2)

		if queue.Dequeue() != 1 {
			t.Error("First dequeue should return 1")
		}

		queue.Enqueue(3)
		queue.Enqueue(4)

		if queue.Dequeue() != 2 {
			t.Error("Second dequeue should return 2")
		}
		if queue.Dequeue() != 3 {
			t.Error("Third dequeue should return 3")
		}
		if queue.Dequeue() != 4 {
			t.Error("Fourth dequeue should return 4")
		}

		if !queue.IsEmpty() {
			t.Error("Queue should be empty after all operations")
		}
	})
}
