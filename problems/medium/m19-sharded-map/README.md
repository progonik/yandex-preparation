# M19 — Sharded map

Implement:

    type Map[V any] struct { /* your fields */ }

    func NewMap[V any](shards int) (*Map[V], error)
    func (m *Map[V]) Load(key string) (V, bool)
    func (m *Map[V]) Store(key string, value V)
    func (m *Map[V]) Delete(key string)
    func (m *Map[V]) LoadOrStore(key string, value V) (actual V, loaded bool)
    func (m *Map[V]) Snapshot() map[string]V

Requirements:

- Hash keys deterministically to fixed shards.
- Each shard has its own lock.
- LoadOrStore is atomic per key.
- Snapshot must represent one logical instant across all shards.
- Never copy a mutex after use.
- The map has no Close operation.

Specify a global lock acquisition order for Snapshot and explain why inconsistent ordering can deadlock.
