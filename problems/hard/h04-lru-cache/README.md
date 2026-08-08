# H04 — Concurrent LRU cache

Implement:

    type Cache[K comparable, V any] struct { /* your fields */ }

    func NewCache[K comparable, V any](capacity int) (*Cache[K, V], error)
    func (c *Cache[K, V]) Get(key K) (V, bool)
    func (c *Cache[K, V]) Put(key K, value V) (evictedKey K, evicted bool)
    func (c *Cache[K, V]) Delete(key K) bool
    func (c *Cache[K, V]) Len() int

Requirements:

- Get and updating Put mark the entry most recently used.
- New Put evicts the least recently used item when full.
- Every method is concurrency-safe and O(1) average time.
- The zero capacity case is rejected.
- Do not expose internal list nodes or values by reference.
- The structure must stay consistent if many goroutines update the same key.

Use a map plus doubly linked list. State the invariants connecting map entries, list nodes, head, tail, and length.
