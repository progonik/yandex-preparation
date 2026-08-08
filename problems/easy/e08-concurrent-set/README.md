# E08 — Concurrent string set

Implement:

    type Set struct { /* your fields */ }

    func (s *Set) Add(value string) bool
    func (s *Set) Remove(value string) bool
    func (s *Set) Contains(value string) bool
    func (s *Set) Snapshot() []string

Add and Remove report whether they changed the set.

Requirements:

- All methods are concurrency-safe.
- Snapshot returns a sorted copy and must not expose internal state.
- The zero value must be usable.
- Do not hold a lock while sorting.

Explain whether RWMutex is necessarily faster than Mutex and identify the linearization point of each mutation.
