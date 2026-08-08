# E05 — Ordered parallel map

Implement a generic helper:

    func Map[T, R any](
        ctx context.Context,
        values []T,
        workers int,
        fn func(context.Context, T) (R, error),
    ) ([]R, error)

Requirements:

- Preserve input order.
- Fail fast on the first processing error.
- Cancel the derived context and join all workers on failure.
- Return nil results on error.
- workers <= 0 is invalid.
- Do not start more workers than input elements.

State which shared fields require synchronization and why writing different output indexes does not require a mutex.
