# M09 — Backpressured pipeline

Build a three-stage pipeline:

    func Pipeline(
        ctx context.Context,
        input <-chan string,
        parseWorkers int,
        saveWorkers int,
        parse func(string) (Record, error),
        save func(context.Context, Record) error,
    ) error

Requirements:

- Parse in parallel, then save parsed records in parallel.
- Bound every internal queue; choose and document capacities.
- On first parse or save error, cancel the pipeline.
- Stop reading new input after cancellation.
- Join every stage and return the triggering error.
- The input channel is owned by the caller and must not be closed.
- No stage may block forever sending downstream.

Draw the channel ownership and shutdown flow before coding. Explain how backpressure propagates to the input reader.
