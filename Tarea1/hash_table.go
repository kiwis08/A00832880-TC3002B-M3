package main

import (
	"hash/fnv"
)

type HashMap[K comparable, V any] struct {
	buckets    []*entry[K, V]
	size       int
	loadFactor float64
	hasher     func(K) uint64
}

type entry[K comparable, V any] struct {
	key   K
	value V
	next  *entry[K, V]
}

// I'll use a load factor of 0.75 by default.
func NewHashMap[K comparable, V any](capacity int, hasher func(K) uint64) *HashMap[K, V] {
	if capacity < 1 {
		capacity = 8
	}
	if hasher == nil {
		panic("NewHashMap: hasher must not be nil")
	}
	return &HashMap[K, V]{
		buckets:    make([]*entry[K, V], capacity),
		loadFactor: 0.75,
		hasher:     hasher,
	}
}

func NewStringMap[V any](capacity int) *HashMap[string, V] {
	return NewHashMap[string, V](capacity, func(s string) uint64 {
		h := fnv.New64a()
		_, _ = h.Write([]byte(s))
		return h.Sum64()
	})
}

func (h *HashMap[K, V]) Len() int { return h.size }

func (h *HashMap[K, V]) Put(key K, value V) {
	if float64(h.size+1) > h.loadFactor*float64(len(h.buckets)) {
		h.resize(len(h.buckets) * 2)
	}
	idx := h.index(key)
	for e := h.buckets[idx]; e != nil; e = e.next {
		if e.key == key {
			e.value = value
			return
		}
	}
	h.buckets[idx] = &entry[K, V]{key: key, value: value, next: h.buckets[idx]}
	h.size++
}

func (h *HashMap[K, V]) Get(key K) (V, bool) {
	var zero V
	idx := h.index(key)
	for e := h.buckets[idx]; e != nil; e = e.next {
		if e.key == key {
			return e.value, true
		}
	}
	return zero, false
}

func (h *HashMap[K, V]) Delete(key K) (V, bool) {
	var zero V
	idx := h.index(key)
	var prev *entry[K, V]
	for e := h.buckets[idx]; e != nil; e = e.next {
		if e.key == key {
			if prev == nil {
				h.buckets[idx] = e.next
			} else {
				prev.next = e.next
			}
			h.size--
			// Optionally shrink when table is sparse.
			if len(h.buckets) > 8 && float64(h.size) < 0.25*float64(len(h.buckets)) {
				h.resize(len(h.buckets) / 2)
			}
			return e.value, true
		}
		prev = e
	}
	return zero, false
}

func (h *HashMap[K, V]) Contains(key K) bool {
	_, ok := h.Get(key)
	return ok
}

func (h *HashMap[K, V]) Keys() []K {
	keys := make([]K, 0, h.size)
	for _, b := range h.buckets {
		for e := b; e != nil; e = e.next {
			keys = append(keys, e.key)
		}
	}
	return keys
}

func (h *HashMap[K, V]) Values() []V {
	values := make([]V, 0, h.size)
	for _, b := range h.buckets {
		for e := b; e != nil; e = e.next {
			values = append(values, e.value)
		}
	}
	return values
}

func (h *HashMap[K, V]) Clear() {
	for i := range h.buckets {
		h.buckets[i] = nil
	}
	h.size = 0
}

func (h *HashMap[K, V]) index(key K) int {
	return int(h.hasher(key) % uint64(len(h.buckets)))
}

func (h *HashMap[K, V]) resize(newCap int) {
	if newCap < 8 {
		newCap = 8
	}
	newBuckets := make([]*entry[K, V], newCap)
	for _, b := range h.buckets {
		for e := b; e != nil; {
			next := e.next
			idx := int(h.hasher(e.key) % uint64(newCap))
			e.next = newBuckets[idx]
			newBuckets[idx] = e
			e = next
		}
	}
	h.buckets = newBuckets
}
