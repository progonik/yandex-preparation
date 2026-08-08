# M10 — Ordered streaming map

Implement:

    type Item[T any] struct {
        Seq   int
        Value T
        Err   error
    }

    func OrderedMap[T, R any](
        ctx context.Context,
        input <-chan T,
        workers int,
        fn func(context.Context, T) (R, error),
    ) <-chan Item[R]

Requirements:

- Assign sequence numbers in input receive order.
- Process concurrently but emit strictly by sequence number.
- An item error is emitted in its slot and does not stop the stream.
- Bound the number of accepted-but-not-emitted items.
- Cancellation closes output and joins internal goroutines.
- Do not read the entire input before emitting.

Design a reorder buffer and explain the worst-case memory usage when the earliest item is slow.
