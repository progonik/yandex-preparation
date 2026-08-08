# E01 — Parallel square

Implement:

    func Squares(ctx context.Context, values []int, workers int) ([]int, error)

Compute the square of every input using a fixed worker pool.

Requirements:

- At most workers computations may be active.
- Results must preserve input order.
- workers <= 0 is an error.
- Empty input returns an empty non-nil slice.
- Stop promptly and return ctx.Err() if the context is cancelled.
- Do not leave goroutines running after return.

Explain why separate workers may safely write to separate slice indexes and identify the owner of the jobs channel.
