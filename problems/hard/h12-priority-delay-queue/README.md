# H12 — Priority delay queue

Implement:

    type Item[T any] struct {
        ID       string
        ReadyAt  time.Time
        Priority int
        Value    T
    }

    type Queue[T any] struct { /* your fields */ }

    func NewQueue[T any]() *Queue[T]
    func (q *Queue[T]) Push(item Item[T]) error
    func (q *Queue[T]) Pop(ctx context.Context) (Item[T], error)
    func (q *Queue[T]) Remove(id string) bool
    func (q *Queue[T]) Close()

Requirements:

- Pop returns only items whose ReadyAt has arrived.
- Among ready items, return greatest Priority; break ties by insertion order.
- A newly pushed earlier item must wake waiting Pop calls.
- IDs are unique while queued.
- Remove is O(log n).
- Close wakes all waiters; queued items are discarded.
- Multiple Push, Pop, and Remove calls may run concurrently.
- Do not poll with short sleeps.

You may need two heaps or a carefully designed ordering strategy. Explain timer ownership and reset races.
