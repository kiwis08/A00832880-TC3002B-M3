package main

import (
	"testing"
)

func TestArrayHashTable(t *testing.T) {
	ht := NewStringMap[string](8)

	// Test initial state
	t.Run("Initial state", func(t *testing.T) {
		if ht.Len() != 0 {
			t.Error("New hash table should have size 0")
		}
		if _, ok := ht.Get("nonexistent"); ok {
			t.Error("Get on non-existent key should return false")
		}
		if ht.Contains("nonexistent") {
			t.Error("Contains on non-existent key should return false")
		}
		if _, ok := ht.Delete("nonexistent"); ok {
			t.Error("Delete on non-existent key should return false")
		}
		if len(ht.Keys()) != 0 {
			t.Error("Keys on empty hash table should return empty slice")
		}
		if len(ht.Values()) != 0 {
			t.Error("Values on empty hash table should return empty slice")
		}
	})

	// Test basic put and get
	t.Run("Basic put and get", func(t *testing.T) {
		ht.Put("key1", "value1")
		if ht.Len() != 1 {
			t.Error("Hash table size should be 1 after one put")
		}
		if !ht.Contains("key1") {
			t.Error("Hash table should contain key1")
		}
		if val, ok := ht.Get("key1"); !ok || val != "value1" {
			t.Error("Get should return the correct value")
		}

		ht.Put("key2", "value2")
		if ht.Len() != 2 {
			t.Error("Hash table size should be 2 after two puts")
		}
		if val, ok := ht.Get("key2"); !ok || val != "value2" {
			t.Error("Get should return the correct value for key2")
		}
	})

	// Test update existing key
	t.Run("Update existing key", func(t *testing.T) {
		ht.Put("key1", "updated_value1")
		if ht.Len() != 2 {
			t.Error("Size should not change when updating existing key")
		}
		if val, ok := ht.Get("key1"); !ok || val != "updated_value1" {
			t.Error("Get should return updated value")
		}
		if val, ok := ht.Get("key2"); !ok || val != "value2" {
			t.Error("Other keys should not be affected by update")
		}
	})

	// Test delete
	t.Run("Delete", func(t *testing.T) {
		if _, ok := ht.Delete("key1"); !ok {
			t.Error("Delete should return true for existing key")
		}
		if ht.Len() != 1 {
			t.Error("Size should be 1 after deleting one key")
		}
		if ht.Contains("key1") {
			t.Error("Hash table should not contain deleted key")
		}
		if _, ok := ht.Get("key1"); ok {
			t.Error("Get should return false for deleted key")
		}
		if val, ok := ht.Get("key2"); !ok || val != "value2" {
			t.Error("Other keys should not be affected by delete")
		}

		if _, ok := ht.Delete("key2"); !ok {
			t.Error("Delete should return true for existing key")
		}
		if ht.Len() != 0 {
			t.Error("Size should be 0 after deleting all keys")
		}
	})

	// Test with different string keys
	t.Run("Different string keys", func(t *testing.T) {
		ht.Put("string_key", "string_value")
		ht.Put("42", "int_key_value")
		ht.Put("3.14", "float_key_value")
		ht.Put("true", "bool_key_value")

		if val, ok := ht.Get("string_key"); !ok || val != "string_value" {
			t.Error("Should handle string keys")
		}
		if val, ok := ht.Get("42"); !ok || val != "int_key_value" {
			t.Error("Should handle string keys")
		}
		if val, ok := ht.Get("3.14"); !ok || val != "float_key_value" {
			t.Error("Should handle string keys")
		}
		if val, ok := ht.Get("true"); !ok || val != "bool_key_value" {
			t.Error("Should handle string keys")
		}

		if ht.Len() != 4 {
			t.Error("Size should be 4 after adding 4 different keys")
		}
	})

	// Test keys and values
	t.Run("Keys and Values", func(t *testing.T) {
		ht.Clear()
		ht.Put("a", "1")
		ht.Put("b", "2")
		ht.Put("c", "3")

		keys := ht.Keys()
		values := ht.Values()

		if len(keys) != 3 {
			t.Error("Keys should return 3 elements")
		}
		if len(values) != 3 {
			t.Error("Values should return 3 elements")
		}

		// Check that all keys are present (order doesn't matter)
		keyMap := make(map[string]bool)
		for _, key := range keys {
			keyMap[key] = true
		}
		if !keyMap["a"] || !keyMap["b"] || !keyMap["c"] {
			t.Error("Keys should contain all added keys")
		}

		// Check that all values are present (order doesn't matter)
		valueMap := make(map[string]bool)
		for _, value := range values {
			valueMap[value] = true
		}
		if !valueMap["1"] || !valueMap["2"] || !valueMap["3"] {
			t.Error("Values should contain all added values")
		}
	})

	// Test clear
	t.Run("Clear", func(t *testing.T) {
		ht.Clear()
		if ht.Len() != 0 {
			t.Error("Size should be 0 after clear")
		}
		if len(ht.Keys()) != 0 {
			t.Error("Keys should be empty after clear")
		}
		if len(ht.Values()) != 0 {
			t.Error("Values should be empty after clear")
		}
	})

	// Test hash collisions (if your hash function can produce them)
	t.Run("Hash collisions", func(t *testing.T) {
		ht.Clear()
		// Add multiple keys that might hash to the same bucket
		// This test depends on your hash function implementation
		ht.Put("key1", "value1")
		ht.Put("key2", "value2")
		ht.Put("key3", "value3")

		if val, ok := ht.Get("key1"); !ok || val != "value1" {
			t.Error("Should handle hash collisions correctly")
		}
		if val, ok := ht.Get("key2"); !ok || val != "value2" {
			t.Error("Should handle hash collisions correctly")
		}
		if val, ok := ht.Get("key3"); !ok || val != "value3" {
			t.Error("Should handle hash collisions correctly")
		}
	})
}

func TestArrayHashTableWithCapacity(t *testing.T) {
	ht := NewStringMap[string](4)

	if ht.Len() != 0 {
		t.Error("New hash table should be empty")
	}

	// Test that it works with the specified capacity
	ht.Put("test", "value")
	if val, ok := ht.Get("test"); !ok || val != "value" {
		t.Error("Hash table with custom capacity should work correctly")
	}
}
