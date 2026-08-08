# H13 — Adaptive concurrency limiter

Implement:

    type Limiter struct { /* your fields */ }

    func NewLimiter(min, max, initial int) (*Limiter, error)
    func (l *Limiter) Do(
        ctx context.Context,
        call func(context.Context) error,
    ) error
    func (l *Limiter) Limit() int
    func (l *Limiter) Close()

Requirements:

- At most the current limit calls execute.
- After 20 consecutive successes, increase limit by one up to max.
- On an overload sentinel error, halve the limit down to min and reset successes.
- Lowering the limit does not cancel active calls but blocks new calls until active count falls.
- Waiting acquisition is cancellable and FIFO.
- Close rejects waiters and new calls but waits for active calls.
- Do not replace a semaphore channel while goroutines are using it.

Design explicit waiter and active-count state. State the invariants when limit becomes lower than active.
