# H19 — Concurrent dependency scheduler

Implement a long-lived scheduler:

    type Job struct {
        ID       string
        Depends  []string
        Priority int
        Run      func(context.Context) error
    }

    type Scheduler struct { /* your fields */ }

    func NewScheduler(ctx context.Context, workers int) (*Scheduler, error)
    func (s *Scheduler) Submit(job Job) error
    func (s *Scheduler) Cancel(id string) bool
    func (s *Scheduler) Wait(ctx context.Context, id string) error
    func (s *Scheduler) Close()

Requirements:

- Dependencies may be submitted after dependants.
- Ready jobs use greatest priority, then submission order.
- Unknown dependencies keep a job blocked until supplied.
- Reject any submission that creates a cycle.
- Failed or cancelled jobs cause transitive dependency failure.
- Wait supports many independently cancellable callers.
- Close cancels active jobs, resolves all waiters, and joins workers.
- Remove completed graph state only when no unresolved future reference requires it; define a retention policy.

Use one coordinator to own graph state and avoid invoking user code while holding scheduler locks.
