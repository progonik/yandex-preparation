# H17 — Multi-stage resilient pipeline

Build:

    func RunPipeline(
        ctx context.Context,
        input <-chan Input,
        decodeWorkers int,
        enrichWorkers int,
        writeWorkers int,
        decode func(Input) (Decoded, error),
        enrich func(context.Context, Decoded) (Enriched, error),
        write func(context.Context, Enriched) error,
    ) error

Requirements:

- Three bounded parallel stages with bounded channels.
- Preserve input order at the write stage despite parallel earlier stages.
- Decode errors are terminal.
- Enrich retries twice with cancellable exponential backoff, then becomes terminal.
- Write errors are terminal.
- Cancel the entire graph on terminal error and return the original cause.
- Stop reading caller-owned input after cancellation.
- Join all stage goroutines and close each channel exactly once.
- Do not allow reorder-buffer memory to grow without bound.

Draw the complete lifecycle, including coordinator goroutines, before coding.
