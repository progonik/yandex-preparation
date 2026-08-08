# H10 — Supervised worker pool

Given:

    type Worker func(context.Context) error

Implement:

    func Supervise(
        ctx context.Context,
        count int,
        maxRestarts int,
        backoff func(attempt int) time.Duration,
        worker Worker,
    ) error

Requirements:

- Keep count worker instances active.
- Restart an instance when it returns a non-nil error.
- A normal nil return is terminal and shuts down the group.
- Each worker slot has its own restart count.
- Restart waits are cancellable.
- Exceeding maxRestarts in any slot cancels and joins the group.
- Capture panics as errors without crashing the process.
- Result delivery must not block after cancellation.

Explain the difference between restarting a goroutine and reusing the same goroutine loop, and define panic recovery boundaries.
