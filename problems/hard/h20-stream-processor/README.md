# H20 — In-memory stream processor

Given:

    type Event struct {
        Key   string
        Value int64
        At    time.Time
    }

    type Result struct {
        Key         string
        WindowStart time.Time
        Count       int
        Sum         int64
    }

Implement:

    func Aggregate(
        ctx context.Context,
        input <-chan Event,
        window time.Duration,
        allowedLateness time.Duration,
        workers int,
    ) (<-chan Result, <-chan Event)

Requirements:

- Compute tumbling-window count and sum per key.
- Partition keys across workers; all events for a key go to the same worker.
- Input may be out of order.
- Track a global maximum event time and close windows behind its watermark.
- Events older than the watermark go to the late-events output.
- Emit each window exactly once.
- Bound queues and propagate backpressure.
- On input close, flush every remaining window and close both outputs.
- Cancellation must unblock every stage and join goroutines.
- Results across keys may be unordered.

State the window boundary formula, watermark rule, channel ownership, and shutdown ordering before implementation.
