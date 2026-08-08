# H03 — Fair multi-tenant scheduler

Given:

    type Task struct {
        Tenant string
        Run    func(context.Context) error
    }

Implement:

    type Scheduler struct { /* your fields */ }

    func NewScheduler(ctx context.Context, workers int) (*Scheduler, error)
    func (s *Scheduler) Submit(ctx context.Context, task Task) error
    func (s *Scheduler) Close() []error

Requirements:

- Maintain FIFO order within each tenant.
- Schedule non-empty tenants round-robin so one noisy tenant cannot starve others.
- At most one task per tenant may run at once.
- At most workers tasks run globally.
- Submit is cancellable until accepted into a bounded queue.
- Close rejects new tasks, drains accepted work, and waits.
- Remove idle tenant state.

Design one coordinator goroutine that owns scheduling state. Explain how worker completions wake it and how bounded admission creates backpressure.
