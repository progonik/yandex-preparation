# M08 — Resizable worker pool

Implement:

    type Pool struct { /* your fields */ }

    func NewPool(workers int, handle func(int)) (*Pool, error)
    func (p *Pool) Submit(ctx context.Context, value int) error
    func (p *Pool) Resize(workers int) error
    func (p *Pool) Close()

Requirements:

- Resize may increase or decrease the number of workers at runtime.
- Downsizing lets selected workers finish their current item before exiting.
- Accepted items are never lost.
- Close rejects submissions, drains the queue, stops all workers, and waits.
- Resize racing with Close returns a shutdown error.
- handle calls may overlap up to the current target worker count.

Do not close a shared per-worker stop channel to remove only one worker. Design explicit worker identities or cancellation handles.
