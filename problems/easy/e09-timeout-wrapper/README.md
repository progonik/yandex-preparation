# E09 — Timeout wrapper

Implement:

    type Operation func(context.Context) error

    func WithTimeout(
        ctx context.Context,
        timeout time.Duration,
        op Operation,
    ) error

Requirements:

- Execute op with a child context limited by timeout.
- Return the operation error if it finishes first.
- Return the child context error when timeout or parent cancellation wins.
- Wait for op before returning; assume it respects cancellation.
- timeout <= 0 should result in immediate cancellation.

Describe the race where completion and timeout happen together and define an acceptable precedence rule.
