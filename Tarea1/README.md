# Data Structures Implementation Template

This project contains templates for three fundamental data structures in Go:

1. **Stack** - Last-In-First-Out (LIFO) data structure
2. **Queue** - First-In-First-Out (FIFO) data structure  
3. **Hash Table** - Key-value mapping with average O(1) operations

## Project Structure

```
Tarea1/
├── go.mod              # Go module file
├── stack.go            # Stack interface and ArrayStack struct
├── stack_test.go       # Stack tests
├── queue.go            # Queue interface and ArrayQueue struct
├── queue_test.go       # Queue tests
├── hash_table.go       # HashTable interface and ArrayHashTable struct
├── hash_table_test.go  # HashTable tests
└── README.md           # This file
```

## Implementation Instructions

### 1. Stack Implementation

**File**: `stack.go`

The `ArrayStack` struct needs to be implemented with the following methods:
- `Push(value interface{})` - Add element to top
- `Pop() interface{}` - Remove and return top element
- `Peek() interface{}` - Return top element without removing
- `IsEmpty() bool` - Check if stack is empty
- `Size() int` - Return number of elements
- `Clear()` - Remove all elements

**Hint**: Use a slice to store elements. The "top" of the stack can be the last element in the slice.

### 2. Queue Implementation

**File**: `queue.go`

The `ArrayQueue` struct needs to be implemented with the following methods:
- `Enqueue(value interface{})` - Add element to rear
- `Dequeue() interface{}` - Remove and return front element
- `Front() interface{}` - Return front element without removing
- `IsEmpty() bool` - Check if queue is empty
- `Size() int` - Return number of elements
- `Clear()` - Remove all elements

**Hint**: Use a slice to store elements. You can use the first element as the front and append to the end for the rear. Consider the efficiency of your implementation.

### 3. Hash Table Implementation

**File**: `hash_table.go`

The `ArrayHashTable` struct needs to be implemented with the following methods:
- `Put(key interface{}, value interface{})` - Store key-value pair
- `Get(key interface{}) interface{}` - Retrieve value by key
- `Delete(key interface{}) bool` - Remove key-value pair
- `Contains(key interface{}) bool` - Check if key exists
- `IsEmpty() bool` - Check if hash table is empty
- `Size() int` - Return number of key-value pairs
- `Clear()` - Remove all key-value pairs
- `Keys() []interface{}` - Return all keys
- `Values() []interface{}` - Return all values

**Additional methods to implement**:
- `hash(key interface{}) int` - Hash function for keys

**Hint**: Use a slice of slices (buckets) to implement separate chaining for collision resolution. The `KeyValue` struct is provided to store key-value pairs in each bucket.

## Running Tests

To test your implementations, run:

```bash
# Test all data structures
go test

# Test specific data structure
go test -run TestArrayStack
go test -run TestArrayQueue  
go test -run TestArrayHashTable

# Run tests with verbose output
go test -v
```

## Implementation Tips

1. **Stack**: The slice approach is straightforward - use `append()` for push and slice the last element for pop.

2. **Queue**: Consider using two indices (front and rear) or use a circular buffer approach for better efficiency.

3. **Hash Table**: 
   - Implement a simple hash function that works with different data types
   - Use separate chaining (slice of slices) for collision resolution
   - Consider implementing resizing when the load factor gets too high
   - Handle edge cases like nil keys/values

4. **Testing**: All test cases are provided. Make sure your implementation passes all tests before considering it complete.

5. **Error Handling**: The interfaces specify return values for edge cases (like nil for empty operations). Make sure to handle these correctly.

## Expected Behavior

- **Stack**: LIFO behavior - last element pushed is first to be popped
- **Queue**: FIFO behavior - first element enqueued is first to be dequeued  
- **Hash Table**: Key-value mapping with average O(1) operations, handles collisions

Good luck with your implementation!
