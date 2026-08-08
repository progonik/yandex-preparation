# E14 — Heartbeat worker

Implement:

    func Work(
        ctx context.Context,
        heartbeatInterval time.Duration,
        jobs <-chan int,
    (heartbeats <-chan struct{}, results <-chan int)

The worker squares jobs and periodically emits a best-effort heartbeat.

Requirements:

- A slow heartbeat consumer must never block job processing.
- Results are reliable: the worker waits to deliver them unless cancelled.
- Close both output channels when jobs closes or context is cancelled.
- Only the worker goroutine closes its channels.
- No goroutine may be started per heartbeat.

Explain why a non-blocking heartbeat send is appropriate but a non-blocking result send is not.
