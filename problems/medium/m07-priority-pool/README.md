# M07 — Priority worker pool

Implement:

    type Task struct {
        ID       int
        Priority int
        Run      func(context.Context) error
    }

    type Pool struct { /* your fields */ }

    func NewPool(ctx context.Context, workers int) (*Pool, error)
    func (p *Pool) Submit(task Task) error
    func (p *Pool) Close() map[int]error

Requirements:

- A higher Priority value runs before lower queued priorities.
- Equal-priority tasks run in submission order.
- Running tasks are never preempted.
- Close rejects new tasks, drains queued tasks, waits, and returns errors by task ID.
- Submit and Close may race safely.
- Do not hold the scheduling mutex while running user code.

Use container/heap. Explain why workers reading directly from a FIFO channel cannot implement strict queued priority.
