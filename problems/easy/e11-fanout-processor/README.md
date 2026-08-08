# E11 — Fan-out string processor

Implement:

    func Process(
        ctx context.Context,
        input <-chan string,
        workers int,
        fn func(string) string,
    ) <-chan string

Several workers must read from input, apply fn, and publish results.

Requirements:

- Output ordering is unspecified.
- Close output only after all workers finish.
- Stop promptly when context is cancelled.
- Workers must not remain blocked sending when the consumer stops after cancellation.
- workers <= 0 returns an already-closed channel.
- fn is concurrency-safe and never panics.

Explain the fan-out/fan-in topology and identify the single owner responsible for closing output.
