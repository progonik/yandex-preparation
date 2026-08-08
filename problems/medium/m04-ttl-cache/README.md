# M04 — TTL cache

Implement:

    type Cache[K comparable, V any] struct { /* your fields */ }

    func NewCache[K comparable, V any](
        ttl time.Duration,
    ) (*Cache[K, V], error)

    func (c *Cache[K, V]) Get(key K) (V, bool)
    func (c *Cache[K, V]) Set(key K, value V)
    func (c *Cache[K, V]) Delete(key K)
    func (c *Cache[K, V]) Close()

Requirements:

- Entries expire ttl after their latest Set.
- Get never returns expired data.
- A background goroutine removes expired entries.
- Close is idempotent and waits for cleanup termination.
- Methods are safe during concurrent Close.
- Set after Close returns a sentinel error; adjust its signature accordingly.
- The implementation must not start one timer per entry.

Use an injectable clock or expiration interval for deterministic tests. Explain lazy versus eager expiration.
