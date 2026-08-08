# H18 — Replicated write coordinator

Given:

    type Replica interface {
        Write(context.Context, string, []byte) error
        Rollback(context.Context, string) error
    }

Implement:

    func ReplicatedWrite(
        ctx context.Context,
        replicas []Replica,
        key string,
        value []byte,
        required int,
    ) error

Requirements:

- Start writes concurrently with a configurable maximum concurrency; add the parameter.
- Succeed once required writes succeed.
- Cancel writes still in progress after quorum and join them.
- If quorum becomes impossible, roll back every replica known to have succeeded.
- Rollbacks run concurrently with the same bound and use a fresh cleanup context with a fixed timeout.
- Return the write failure joined with rollback failures.
- Copy value if implementations may retain it.
- Validate required and the concurrency bound.

Explain the limits of rollback as a distributed transaction and identify ambiguous outcomes when cancellation races with a remote success.
