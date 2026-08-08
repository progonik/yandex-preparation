# H08 — Partitioned ordered processor

Given:

    type Event struct {
        Partition string
        Sequence  int64
        Payload   []byte
    }

Implement:

    func Process(
        ctx context.Context,
        input <-chan Event,
        workers int,
        handle func(context.Context, Event) error,
    ) error

Requirements:

- Events within a partition execute one at a time in ascending sequence order.
- Input may arrive out of order.
- Different partitions may execute concurrently.
- At most workers handlers run globally.
- Duplicate sequence numbers are ignored.
- A gap blocks later events in that partition.
- On input close, return an error describing unresolved gaps.
- First handler error cancels and joins everything.
- Clean up completed partition state.

Define how the first expected sequence is established and use a heap or ordered buffer per partition.
