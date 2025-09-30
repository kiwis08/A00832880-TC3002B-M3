package main

import (
	"testing"
)

func TestArrayStack(t *testing.T) {
	stack := NewArrayStack()

	// Test initial state
	t.Run("Initial state", func(t *testing.T) {
		if !stack.IsEmpty() {
			t.Error("New stack should be empty")
		}
		if stack.Size() != 0 {
			t.Error("New stack should have size 0")
		}
		if stack.Peek() != nil {
			t.Error("Peek on empty stack should return nil")
		}
		if stack.Pop() != nil {
			t.Error("Pop on empty stack should return nil")
		}
	})

	// Test push and peek
	t.Run("Push and Peek", func(t *testing.T) {
		stack.Push(1)
		if stack.IsEmpty() {
			t.Error("Stack should not be empty after push")
		}
		if stack.Size() != 1 {
			t.Error("Stack size should be 1 after one push")
		}
		if stack.Peek() != 1 {
			t.Error("Peek should return the last pushed element")
		}

		stack.Push(2)
		if stack.Size() != 2 {
			t.Error("Stack size should be 2 after two pushes")
		}
		if stack.Peek() != 2 {
			t.Error("Peek should return the most recently pushed element")
		}
	})

	// Test pop
	t.Run("Pop", func(t *testing.T) {
		// Stack should have [1, 2] with 2 on top
		popped := stack.Pop()
		if popped != 2 {
			t.Error("Pop should return the most recently pushed element")
		}
		if stack.Size() != 1 {
			t.Error("Stack size should be 1 after one pop")
		}
		if stack.Peek() != 1 {
			t.Error("Peek should return 1 after popping 2")
		}

		popped = stack.Pop()
		if popped != 1 {
			t.Error("Pop should return 1")
		}
		if !stack.IsEmpty() {
			t.Error("Stack should be empty after popping all elements")
		}
		if stack.Size() != 0 {
			t.Error("Stack size should be 0 after popping all elements")
		}
	})

	// Test multiple operations
	t.Run("Multiple operations", func(t *testing.T) {
		// Push multiple elements
		elements := []interface{}{"a", "b", "c", "d", "e"}
		for _, elem := range elements {
			stack.Push(elem)
		}

		// Verify LIFO order
		for i := len(elements) - 1; i >= 0; i-- {
			if stack.Peek() != elements[i] {
				t.Errorf("Expected %v, got %v", elements[i], stack.Peek())
			}
			popped := stack.Pop()
			if popped != elements[i] {
				t.Errorf("Expected %v, got %v", elements[i], popped)
			}
		}

		if !stack.IsEmpty() {
			t.Error("Stack should be empty after popping all elements")
		}
	})

	// Test clear
	t.Run("Clear", func(t *testing.T) {
		// Add some elements
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		stack.Clear()
		if !stack.IsEmpty() {
			t.Error("Stack should be empty after clear")
		}
		if stack.Size() != 0 {
			t.Error("Stack size should be 0 after clear")
		}
		if stack.Peek() != nil {
			t.Error("Peek should return nil after clear")
		}
	})

	// Test with different data types
	t.Run("Different data types", func(t *testing.T) {
		stack.Push("string")
		stack.Push(42)
		stack.Push(3.14)
		stack.Push(true)

		if stack.Pop() != true {
			t.Error("Should handle boolean values")
		}
		if stack.Pop() != 3.14 {
			t.Error("Should handle float values")
		}
		if stack.Pop() != 42 {
			t.Error("Should handle integer values")
		}
		if stack.Pop() != "string" {
			t.Error("Should handle string values")
		}
	})
}
