# H01 — Bounded concurrent crawler

Given:

    type Fetcher interface {
        Fetch(context.Context, string) (links []string, err error)
    }

Implement:

    func Crawl(
        ctx context.Context,
        fetcher Fetcher,
        seeds []string,
        maxDepth int,
        workers int,
    ) (visited []string, err error)

Requirements:

- Fetch each normalized URL at most once.
- Seed depth is zero; do not fetch links deeper than maxDepth.
- At most workers Fetch calls are active.
- Newly discovered work may create more work while the crawl is running.
- Return URLs sorted for deterministic output.
- Fail fast, cancel, and join on the first fetch error.
- Completion detection must work when the dynamic work queue becomes empty.
- Avoid deadlock when workers both produce and consume jobs.

Explain how you count outstanding work and why closing a jobs channel from an arbitrary worker is unsafe.
