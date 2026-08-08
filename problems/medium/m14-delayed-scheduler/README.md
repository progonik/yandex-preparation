# M14 — Delayed scheduler

Implement:

    type Scheduler struct { /* your fields */ }

    func NewScheduler(
        ctx context.Context,
        workers int,
    ) (*Scheduler, error)

    func (s *Scheduler) Schedule(
        id string,
        at time.Time,
        task func(context.Context),
    ) error

    func (s *Scheduler) Cancel(id string) bool
    func (s *Scheduler) Close()

Requirements:

- Execute no earlier than at.
- At most workers scheduled tasks run concurrently.
- IDs are unique while pending or running.
- Cancel succeeds only for pending tasks.
- Adding an earlier task must wake the scheduler and reset its timer.
- Close rejects new work, cancels pending tasks, cancels active contexts, and waits.
- Do not create one sleeping goroutine per task.

Use a min-heap and explain the timer reset protocol needed to avoid stale ticks.
