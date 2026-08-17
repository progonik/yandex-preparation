# C01 - Graceful background worker

Target: 60 minutes without a compiler.

Implement:

```go
var ErrStopped = errors.New("worker is stopped")

type Worker struct {
    // your fields
}

func Start(handler func(int), queueCapacity int) *Worker
func (w *Worker) Submit(ctx context.Context, value int) error
func (w *Worker) Stop()
```

Assume `queueCapacity >= 0` and `handler != nil`. The handler does not panic.

`Start` creates exactly one background goroutine. That goroutine calls `handler` for accepted values.

## Requirements

- Handler calls never overlap.
- `Submit` returns `nil` only if the value has been accepted.
- Every accepted value is passed to `handler` exactly once.
- While the worker is running, `Submit` may wait for queue space and must return `ctx.Err()` if its context wins before acceptance.
- `Stop` stops accepting new values, drains every accepted value, and waits for the background goroutine to exit.
- After shutdown begins, new calls to `Submit` return `ErrStopped`.
- When `Submit` and `Stop` overlap, either outcome is valid:
  - submission succeeds and the value must be handled; or
  - submission returns `ErrStopped` and the value must not be handled.
- `Stop` is idempotent. Concurrent calls to `Stop` must all wait for the same shutdown to complete.
- The implementation must never send to a closed channel.
- No goroutine may remain after `Stop` returns.

## Example

```go
var mu sync.Mutex
var handled []int

w := Start(func(value int) {
    mu.Lock()
    handled = append(handled, value)
    mu.Unlock()
}, 2)

_ = w.Submit(context.Background(), 10)
_ = w.Submit(context.Background(), 20)
w.Stop()

err := w.Submit(context.Background(), 30)
// errors.Is(err, ErrStopped) == true
// handled contains 10 and 20, each exactly once.
```

## Questions to answer before coding

1. What are the worker's lifecycle states?
2. At what exact point does a submission become accepted?
3. How can `Stop` prove that no sender can use the jobs channel before closing it?
4. What makes concurrent `Stop` calls wait rather than return early?
