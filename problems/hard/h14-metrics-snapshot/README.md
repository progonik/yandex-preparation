# H14 — Consistent metrics snapshot

Implement:

    type Metrics struct { /* your fields */ }

    func NewMetrics(shards int) (*Metrics, error)
    func (m *Metrics) Add(name string, delta int64)
    func (m *Metrics) Observe(name string, value float64)
    func (m *Metrics) Snapshot() Snapshot

Snapshot contains each counter and, for observations, count, sum, minimum, and maximum.

Requirements:

- Updates are sharded to reduce contention.
- Snapshot must be internally consistent at one logical instant: count and sum cannot come from different update sets.
- Snapshot must not expose mutable state.
- Add and Observe for different metric names should usually proceed independently.
- NaN observations are rejected; adjust the signature.
- The zero observation state has no min/max value.

Choose and document an all-shard lock order. Compare it with atomics and explain why several independent atomic loads may not form a consistent snapshot.
