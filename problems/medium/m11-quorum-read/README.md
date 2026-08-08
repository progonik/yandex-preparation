# M11 — Quorum read

Given:

    type Replica interface {
        Read(context.Context, string) (version int64, value string, err error)
    }

Implement:

    func QuorumRead(
        ctx context.Context,
        replicas []Replica,
        key string,
        quorum int,
    ) (string, error)

Requirements:

- Query all replicas concurrently.
- Once quorum successful responses arrive, return the value with the highest version among those responses.
- Cancel and join remaining reads.
- If quorum becomes impossible because of failures, return errors.Join of observed errors.
- Validate 1 <= quorum <= len(replicas).
- External cancellation returns ctx.Err().
- Replica implementations respect cancellation.

Explain why returning after quorum without joining cancelled calls violates the no-leak contract.
