# M01 — Retrying worker pool

Given:

    type Job struct { ID int }
    type Process func(context.Context, Job) error

Implement:

    func Run(
        ctx context.Context,
        jobs []Job,
        workers int,
        maxAttempts int,
        process Process,
    ) error

Requirements:

- No more than workers Process calls may be active.
- Retry a failed job until it succeeds or reaches maxAttempts.
- Wait 10ms * 2^(attempt-1) before retries.
- Retry waits must be cancellable.
- Different jobs may progress independently.
- After attempts are exhausted, cancel all work, join workers, and return an error containing the job ID.
- No new attempts begin after terminal failure is observed.

Discuss whether retries should re-enter the shared queue or stay in the same worker and how either choice affects fairness.
