# H07 — Lease manager

Implement an in-memory exclusive lease service:

    type Lease struct {
        Key   string
        Token uint64
        Until time.Time
    }

    type Manager struct { /* your fields */ }

    func (m *Manager) Acquire(
        ctx context.Context,
        key string,
        ttl time.Duration,
    ) (Lease, error)

    func (m *Manager) Renew(lease Lease, ttl time.Duration) (Lease, error)
    func (m *Manager) Release(lease Lease) bool
    func (m *Manager) Close()

Requirements:

- Only one unexpired lease exists per key.
- Acquire waits until release/expiration or context cancellation.
- Every acquisition receives a monotonically increasing fencing token.
- Stale Lease values cannot renew or release newer ownership.
- Expiration wakes waiters without one goroutine per lease.
- Close wakes all waiters and rejects operations.

Use a timer and min-heap or one condition loop. Explain why a random owner ID alone is weaker than a fencing token.
