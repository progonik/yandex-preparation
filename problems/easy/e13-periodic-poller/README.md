# E13 — Periodic poller

Implement:

    type Poll func(context.Context) error

    func StartPoller(
        ctx context.Context,
        interval time.Duration,
        poll Poll,
    ) <-chan error

Run poll immediately and then once per interval.

Requirements:

- Send every poll error to the returned channel without stopping.
- The output channel closes after cancellation and the active poll finishes.
- Do not overlap poll calls.
- Cancellation must unblock an attempted error send.
- Reject interval <= 0 by returning a channel containing one error and then closing it.

Explain why time.Ticker must be stopped and whether ticks missed during a slow poll should be queued.
