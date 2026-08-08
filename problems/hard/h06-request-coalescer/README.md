# H06 — Request coalescing batcher

Given:

    type BatchLoad[K comparable, V any] func(
        context.Context,
        []K,
    ) (map[K]V, error)

Implement:

    type Coalescer[K comparable, V any] struct { /* your fields */ }

    func NewCoalescer[K comparable, V any](
        ctx context.Context,
        maxBatch int,
        maxWait time.Duration,
        load BatchLoad[K, V],
    ) (*Coalescer[K, V], error)

    func (c *Coalescer[K, V]) Get(ctx context.Context, key K) (V, error)
    func (c *Coalescer[K, V]) Close()

Requirements:

- Coalesce waiting keys into a load call by size or time.
- Duplicate keys in the same pending batch share one loaded value.
- Each caller may cancel independently.
- Remove a key if every waiter cancels before dispatch.
- Limit to one BatchLoad call at a time.
- Missing keys receive a sentinel error.
- Close rejects new calls, cancels pending waiters, cancels active loading, and waits.

Separate coordinator-owned state from per-caller reply channels and reason about reply buffers.
