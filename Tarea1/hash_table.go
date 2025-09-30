package main

// HashTable represents a key-value mapping data structure
// that provides average O(1) time complexity for basic operations
type HashTable interface {
	// Put stores a key-value pair in the hash table
	// If the key already exists, it updates the value
	Put(key interface{}, value interface{})

	// Get retrieves the value associated with the given key
	// Returns nil if the key doesn't exist
	Get(key interface{}) interface{}

	// Delete removes the key-value pair from the hash table
	// Returns true if the key was found and removed, false otherwise
	Delete(key interface{}) bool

	// Contains returns true if the key exists in the hash table
	Contains(key interface{}) bool

	// IsEmpty returns true if the hash table has no key-value pairs
	IsEmpty() bool

	// Size returns the number of key-value pairs in the hash table
	Size() int

	// Clear removes all key-value pairs from the hash table
	Clear()

	// Keys returns a slice of all keys in the hash table
	Keys() []interface{}

	// Values returns a slice of all values in the hash table
	Values() []interface{}
}

// KeyValue represents a key-value pair
type KeyValue struct {
	Key   interface{}
	Value interface{}
}

// ArrayHashTable implements HashTable using a slice of slices (buckets)
// This is a simple implementation using separate chaining for collision resolution
type ArrayHashTable struct {
	// TODO: Add fields to store the hash table data
	// Hint: You might want to use:
	// - A slice of slices to implement buckets (separate chaining)
	// - A size field to track the number of elements
	// - A capacity field to track the number of buckets
}

// NewArrayHashTable creates a new empty hash table with default capacity
func NewArrayHashTable() *ArrayHashTable {
	// TODO: Initialize and return a new ArrayHashTable
	// Consider starting with a reasonable default capacity (e.g., 16)
	return nil
}

// NewArrayHashTableWithCapacity creates a new empty hash table with specified capacity
func NewArrayHashTableWithCapacity(capacity int) *ArrayHashTable {
	// TODO: Initialize and return a new ArrayHashTable with given capacity
	return nil
}

// hash function computes the hash code for a given key
// This is a simple hash function - you might want to improve it
func (ht *ArrayHashTable) hash(key interface{}) int {
	// TODO: Implement a hash function
	// Hint: You can use type assertions to handle different key types
	// For strings: sum of character codes
	// For integers: use the integer value directly
	// For other types: convert to string and hash
	return 0
}

// TODO: Implement all HashTable interface methods for ArrayHashTable
// Remember to handle:
// - Hash collisions using separate chaining
// - Resizing when the load factor gets too high
// - Different key types in your hash function
// - Edge cases like nil keys/values
