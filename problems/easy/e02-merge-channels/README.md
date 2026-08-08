# E02 — Merge channels

Implement:

    func Merge(ctx context.Context, inputs ...<-chan int) <-chan int

Return a channel containing values from every input channel.

Requirements:

- Ordering within each individual input must be preserved; global ordering is unspecified.
- Close the output only after every input closes.
- Cancellation must stop all forwarding goroutines even when the consumer no longer reads the output.
- Nil input channels may be ignored.
- The returned channel must eventually close after cancellation.

Explain why closing the output from a forwarding goroutine is unsafe and how a WaitGroup participates in channel ownership.
