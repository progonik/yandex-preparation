# E04 — Bounded function runner

Implement:

    type Task func(context.Context) error

    func RunAll(ctx context.Context, tasks []Task, limit int) error

Execute every task with no more than limit active calls.

Requirements:

- All tasks run unless the external context is cancelled.
- Do not fail fast: collect task errors and return errors.Join(errs...).
- limit <= 0 is an error.
- Never call a task after cancellation has been observed.
- Wait for active tasks before returning.

Explain the difference between limiting goroutine creation and limiting active task execution.
