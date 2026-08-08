# M17 — Service group shutdown

Given:

    type Service interface {
        Run(context.Context) error
    }

Implement:

    func RunGroup(
        ctx context.Context,
        services []Service,
        shutdownTimeout time.Duration,
    ) error

Requirements:

- Start every service concurrently.
- The first service error cancels the group.
- Normal nil return by any service also stops the group and is represented by a sentinel reason.
- Wait up to shutdownTimeout for all remaining services.
- Return the triggering error joined with shutdown errors.
- Do not leak result-sending goroutines even if timeout expires.

The Service contract alone may make a strict no-leak guarantee impossible. State the necessary contract and design the result channel so senders can always finish.
