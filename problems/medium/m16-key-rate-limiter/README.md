# M16 — Per-key rate limiter

Implement:

    type Limiter struct { /* your fields */ }

    func NewLimiter(
        ratePerSecond int,
        burst int,
        idleTTL time.Duration,
    ) (*Limiter, error)

    func (l *Limiter) Allow(key string, now time.Time) bool
    func (l *Limiter) Close()

Requirements:

- Maintain an independent token bucket per key.
- Calls for different keys should not require one global lock for their full update.
- Buckets refill according to elapsed time and never exceed burst.
- Remove buckets unused for idleTTL in one cleanup goroutine.
- Close stops cleanup and is idempotent.
- Allow after Close returns false.
- Do not start a goroutine per key.

Define behavior for clock movement backwards and explain safe deletion while another goroutine holds a bucket reference.
