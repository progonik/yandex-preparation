# M06 — Token-bucket dispatcher

Implement:

    type Task func(context.Context) error

    func Dispatch(
        ctx context.Context,
        tasks []Task,
        rate int,
        burst int,
        workers int,
    ) []error

Requirements:

- Start at most rate tasks per second with token-bucket burst capacity.
- At most workers tasks execute simultaneously.
- Initially the bucket contains burst tokens.
- Preserve one result slot per input task.
- Cancellation marks unstarted tasks with ctx.Err() and joins active calls.
- Task failures do not cancel siblings.
- Reject non-positive parameters.

Use time.Ticker without accumulating unbounded tokens. Explain why rate limiting task starts and limiting concurrency solve different problems.
