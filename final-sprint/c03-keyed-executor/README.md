# C03 - Per-key serial executor

Target: 60 minutes without a compiler. This is the stretch live-coding problem.

Implement:

```go
type Task struct {
    Key string
    Run func(context.Context) error
}

func Execute(
    ctx context.Context,
    tasks []Task,
    workers int,
) []error
```

Assume `workers >= 1` and every `Task.Run` is non-nil.

## Requirements

- Return a result slice with exactly `len(tasks)` slots.
- The error for `tasks[i]` must be stored in `results[i]`.
- Tasks with the same key execute sequentially in input order.
- Tasks with different keys may execute concurrently.
- At most `workers` calls to `Task.Run` may be active globally.
- A task error does not prevent later tasks for the same key from running.
- Once cancellation is observed, no unstarted task may begin.
- Every task that was not started because of cancellation receives `ctx.Err()` in its result slot.
- A task already running when cancellation happens is allowed to finish; store the error it returns.
- Join every goroutine before `Execute` returns.
- Do not create one permanent goroutine per distinct key.

## Example

For this input order:

```text
0: key=A, task=A1
1: key=B, task=B1
2: key=A, task=A2
3: key=C, task=C1
4: key=B, task=B2
```

Valid execution properties are:

- `A1` finishes before `A2` starts;
- `B1` finishes before `B2` starts;
- `A1`, `B1`, and `C1` may overlap if the worker limit allows it;
- result slots remain in positions `0..4`, regardless of completion order.

## Design discussion required

Before coding, describe how a key moves through these conceptual states:

```text
idle -> ready -> running -> ready or idle
```

State the invariants that ensure:

1. A key is never running twice.
2. A ready key is not scheduled twice.
3. Completion can be detected without leaving goroutines blocked.
