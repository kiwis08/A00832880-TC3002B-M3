package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	fmt.Println("Data Structures Demo")

	fmt.Println("1. Stack Demo:")
	demoStack()

	fmt.Println("\n2. Queue Demo:")
	demoQueue()

	fmt.Println("\n3. Hash Map Demo:")
	demoHashMap()
}

func demoStack() {
	stack := &Stack{}

	stack.Push("First")
	stack.Push("Second")
	stack.Push("Third")

	fmt.Printf("Stack size: %d\n", stack.Size())
	fmt.Printf("Top element: %v\n", stack.Peek())

	fmt.Println("Popping elements:")
	for !stack.IsEmpty() {
		element := stack.Pop()
		fmt.Printf("Popped: %v\n", element)
	}

	fmt.Printf("Stack is empty: %v\n", stack.IsEmpty())
}

func demoQueue() {
	queue := NewArrayQueue()

	queue.Enqueue("First")
	queue.Enqueue("Second")
	queue.Enqueue("Third")

	fmt.Printf("Queue size: %d\n", queue.Size())
	fmt.Printf("Front element: %v\n", queue.Front())

	fmt.Println("Dequeuing elements:")
	for !queue.IsEmpty() {
		element := queue.Dequeue()
		fmt.Printf("Dequeued: %v\n", element)
	}

	fmt.Printf("Queue is empty: %v\n", queue.IsEmpty())
}

func demoHashMap() {
	hasher := func(s string) uint64 {
		h := fnv.New64a()
		_, _ = h.Write([]byte(s))
		return h.Sum64()
	}
	hashMap := NewHashMap[string, int](4, hasher)

	hashMap.Put("apple", 5)
	hashMap.Put("banana", 3)
	hashMap.Put("cherry", 8)
	hashMap.Put("date", 2)

	fmt.Printf("HashMap size: %d\n", hashMap.Len())

	if value, exists := hashMap.Get("banana"); exists {
		fmt.Printf("Value for 'banana': %d\n", value)
	}

	fmt.Printf("Contains 'apple': %v\n", hashMap.Contains("apple"))
	fmt.Printf("Contains 'grape': %v\n", hashMap.Contains("grape"))

	fmt.Printf("Keys: %v\n", hashMap.Keys())
	fmt.Printf("Values: %v\n", hashMap.Values())

	if deletedValue, deleted := hashMap.Delete("cherry"); deleted {
		fmt.Printf("Deleted 'cherry' with value: %d\n", deletedValue)
	}

	fmt.Printf("HashMap size after deletion: %d\n", hashMap.Len())
}
