# E07 — Safe counter

Implement a concurrent counter:

    type Counter struct { /* your fields */ }

    func (c *Counter) Add(delta int64)
    func (c *Counter) Value() int64
    func (c *Counter) Reset() int64

Reset atomically sets the value to zero and returns the previous value.

Requirements:

- Every method may be called concurrently.
- Copying a Counter after first use is unsupported and should be documented.
- Choose either sync.Mutex or sync/atomic and justify the choice.
- The zero value must be ready to use.

Write stress tests that can be run with the race detector.
