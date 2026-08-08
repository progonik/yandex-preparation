# E17 — Concurrent sum

Implement:

    func Sum(ctx context.Context, values []int64, parts int) (int64, error)

Split the input into at most parts contiguous ranges, sum them concurrently, and combine partial sums.

Requirements:

- parts <= 0 is invalid.
- Do not start empty partitions.
- Cancellation returns ctx.Err() and joins all goroutines.
- Integer overflow does not need special handling.
- No shared accumulator may be updated unsafely.

Explain your partition formula for lengths not divisible by parts and compare a partial-results channel with one slice slot per goroutine.
