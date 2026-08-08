# M05 — Size-or-time batcher

Implement:

    type Batcher[T any] struct { /* your fields */ }

    func NewBatcher[T any](
        maxSize int,
        maxWait time.Duration,
        handle func([]T),
    ) (*Batcher[T], error)

    func (b *Batcher[T]) Add(ctx context.Context, value T) error
    func (b *Batcher[T]) Close()

Requirements:

- Flush when maxSize items accumulate or maxWait passes after the first item in a batch.
- Never call handle concurrently.
- Preserve item order.
- Close rejects new items, flushes the final non-empty batch, and waits.
- Add is cancellable until its item has been accepted.
- Close is idempotent and safe concurrently with Add.
- Do not retain the caller's backing array when invoking handle.

Pay special attention to stopping and draining timers correctly when a size-triggered flush wins.
