# E16 — Channel queue

Implement:

    type Queue[T any] struct { /* your fields */ }

    func NewQueue[T any](capacity int) (*Queue[T], error)
    func (q *Queue[T]) Put(ctx context.Context, value T) error
    func (q *Queue[T]) Get(ctx context.Context) (T, error)
    func (q *Queue[T]) Close()

Requirements:

- FIFO order.
- Put blocks when full; Get blocks when empty.
- Both operations are cancellable.
- Close is idempotent.
- After Close, Put returns a sentinel error.
- Existing buffered items may be drained; then Get returns the closed error.
- No send-on-closed-channel panic is allowed under concurrent Put and Close.

Discuss why exposing a raw channel makes safe concurrent Close difficult.
