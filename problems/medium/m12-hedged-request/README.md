# M12 — Hedged request

Given:

    type Call func(context.Context) (string, error)

Implement:

    func Hedge(
        ctx context.Context,
        primary Call,
        backup Call,
        delay time.Duration,
    ) (string, error)

Start the primary immediately. Start the backup only if the primary has not completed after delay.

Requirements:

- Return the first successful response.
- If the primary fails before delay, start the backup immediately.
- If both fail, return errors.Join.
- Cancel and join the losing call after success.
- If the primary succeeds before delay, never call backup.
- External cancellation returns ctx.Err().
- Avoid leaking a timer.

Enumerate the event races between primary completion, timer firing, backup completion, and parent cancellation.
