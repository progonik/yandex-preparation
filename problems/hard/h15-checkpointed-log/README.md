# H15 — Checkpointed log processor

Given:

    type Record struct {
        Offset int64
        Data   []byte
    }

    type Checkpointer interface {
        Commit(context.Context, int64) error
    }

Implement:

    func ProcessLog(
        ctx context.Context,
        input <-chan Record,
        workers int,
        handle func(context.Context, Record) error,
        checkpoint Checkpointer,
    ) error

Requirements:

- Input offsets are strictly increasing.
- Handle records concurrently.
- Commit the greatest contiguous successfully completed offset.
- Never commit past a failed or unfinished record.
- At most one Commit call may be active.
- Coalesce commits when completions arrive quickly.
- First handler or commit error cancels and joins the pipeline.
- Bound the number of in-flight records to control the completion buffer.

Explain how out-of-order completion advances a contiguous frontier and what delivery semantics are possible after a crash.
