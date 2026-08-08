# E15 — Stop after N successes

Implement:

    type Attempt func(context.Context) (string, error)

    func Collect(
        ctx context.Context,
        attempts []Attempt,
        required int,
    ) ([]string, error)

Run all attempts concurrently and return as soon as required successful values exist.

Requirements:

- Cancel and join remaining attempts after the target is reached.
- If reaching required successes becomes impossible, return a sentinel error.
- required == 0 returns an empty result without starting attempts.
- required < 0 or required > len(attempts) is invalid.
- Result ordering is completion order.
- External cancellation returns ctx.Err().

Determine a channel buffer size that prevents completed attempts from being stranded during cancellation.
