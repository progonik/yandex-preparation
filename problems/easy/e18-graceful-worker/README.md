# E18 — Graceful background worker

Implement:

    type Worker struct { /* your fields */ }

    func Start(handler func(int))
    func (w *Worker) Submit(ctx context.Context, value int) error
    func (w *Worker) Stop()

Start creates one background goroutine. Submit queues work. Stop rejects new work, drains already accepted work, and waits for completion.

Requirements:

- Start may be called only once.
- Stop is idempotent and safe concurrently with Submit.
- No accepted value may be lost.
- Submit after shutdown returns a sentinel error.
- handler calls never overlap.
- No send-on-closed-channel panic.

Design the lifecycle states before coding and identify the linearization point at which a submission becomes accepted.
