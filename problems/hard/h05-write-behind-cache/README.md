# H05 — Write-behind cache

Given:

    type Store[K comparable, V any] interface {
        Save(context.Context, K, V) error
    }

Implement a cache whose Set returns after updating memory while persistence happens asynchronously.

    func NewCache[K comparable, V any](
        ctx context.Context,
        store Store[K, V],
        workers int,
        queueSize int,
    ) (*Cache[K, V], error)

    func (c *Cache[K, V]) Set(ctx context.Context, key K, value V) error
    func (c *Cache[K, V]) Get(key K) (V, bool)
    func (c *Cache[K, V]) Close() error

Requirements:

- Writes for the same key persist in Set order.
- Writes for different keys may persist concurrently.
- Set is accepted only if both memory and eventual queueing are guaranteed.
- Backpressure blocks Set when the bounded queue is full.
- Persistence failures are collected and returned from Close.
- Close drains accepted writes and is concurrency-safe.

Explain how to prevent an older slow write from overwriting a newer value in the store.
