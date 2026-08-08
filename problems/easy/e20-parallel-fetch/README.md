# E20 — Parallel page fetch

Given:

    type Fetcher interface {
        Fetch(context.Context, string) ([]byte, error)
    }

    type Page struct {
        URL  string
        Body []byte
    }

Implement:

    func FetchAll(
        ctx context.Context,
        fetcher Fetcher,
        urls []string,
        maxConcurrent int,
    ) ([]Page, error)

Requirements:

- Preserve input order.
- Each URL occurrence is a separate request.
- Fail fast, cancel remaining calls, and return nil results on error.
- Do not exceed maxConcurrent active Fetch calls.
- Join all goroutines before returning.
- External cancellation returns ctx.Err().
- maxConcurrent <= 0 is invalid; empty input succeeds.

After implementation, test ordering, concurrency bounds, first error, external cancellation, and goroutine shutdown.
