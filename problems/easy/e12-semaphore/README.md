# E12 — Semaphore

Implement a context-aware counting semaphore:

    type Semaphore struct { /* your fields */ }

    func NewSemaphore(capacity int) (*Semaphore, error)
    func (s *Semaphore) Acquire(ctx context.Context) error
    func (s *Semaphore) Release()

Requirements:

- At most capacity acquisitions may be held.
- Acquire blocks until a permit is available or ctx is cancelled.
- Release without a matching successful Acquire must panic.
- The constructor rejects non-positive capacity.
- Semaphore methods may be called concurrently.

Use channels or sync.Cond. Explain the cancellation tradeoff of your choice and test that cancelled acquisition never consumes a permit.
