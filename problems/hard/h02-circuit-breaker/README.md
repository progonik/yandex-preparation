# H02 — Circuit breaker

Implement:

    type Breaker struct { /* your fields */ }

    func NewBreaker(
        failureThreshold int,
        openFor time.Duration,
        halfOpenLimit int,
    ) (*Breaker, error)

    func (b *Breaker) Do(
        ctx context.Context,
        call func(context.Context) error,
    ) error

States are closed, open, and half-open.

Requirements:

- In closed state, consecutive failures increment a counter; success resets it.
- Reaching the threshold opens the circuit.
- Calls while open fail immediately with a sentinel error.
- After openFor, enter half-open and allow at most halfOpenLimit probe calls.
- One successful probe closes the breaker; one failed probe reopens it.
- Results from calls started in an older state generation must not corrupt the current state.
- Do not hold a mutex while executing call.

Use an injectable clock for tests and identify every state transition's linearization point.
