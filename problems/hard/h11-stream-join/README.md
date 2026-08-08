# H11 — Event-time stream join

Given:

    type Left struct {
        Key string
        At  time.Time
        A   string
    }

    type Right struct {
        Key string
        At  time.Time
        B   string
    }

    type Joined struct {
        Left
        Right
    }

Implement:

    func Join(
        ctx context.Context,
        left <-chan Left,
        right <-chan Right,
        window time.Duration,
    ) <-chan Joined

Requirements:

- Join events with the same key when their timestamps differ by at most window.
- An event may match several events from the opposite side.
- Inputs can arrive out of timestamp order.
- Assume lateness is bounded by window relative to the greatest timestamp seen on that input.
- Evict state once future matches are impossible.
- Close output after both inputs close and all possible matches are emitted.
- Cancellation stops everything, including blocked output sends.
- Emit matches deterministically by arrival order where possible.

Define the watermark and eviction rule before implementation.
