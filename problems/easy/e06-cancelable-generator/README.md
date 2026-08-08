# E06 — Cancelable generator

Implement:

    func Generate(ctx context.Context, start, step int) <-chan int

Generate start, start+step, start+2*step, and so on until cancellation.

Requirements:

- The returned channel closes after cancellation.
- The producer must exit even if the consumer stops reading without draining.
- Integer overflow does not need special handling.
- Creating the generator with an already-cancelled context must return a channel that closes promptly.

Explain why both the send and cancellation receive must appear in the same select.
