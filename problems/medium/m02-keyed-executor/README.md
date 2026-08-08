# M02 — Per-key serial executor

Implement:

    type Task struct {
        Key string
        Run func(context.Context) error
    }

    func Execute(
        ctx context.Context,
        tasks []Task,
        workers int,
    ) []error

Requirements:

- Tasks with the same Key execute sequentially in input order.
- Tasks with different keys may execute concurrently.
- At most workers tasks execute globally.
- A task error does not stop later tasks for that key.
- Return one error slot per input task.
- Cancellation prevents unstarted tasks and places ctx.Err() in their slots.
- Join every started goroutine.

Avoid creating one permanent goroutine per distinct key. Explain how keys transition between idle, ready, and running states.
