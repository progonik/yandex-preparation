# M13 — Concurrent file aggregation

Given:

    type Reader interface {
        ReadFile(context.Context, string) ([]byte, error)
    }

Implement:

    func CountLines(
        ctx context.Context,
        reader Reader,
        paths []string,
        workers int,
    ) (map[string]int, error)

Requirements:

- Read at most workers files concurrently.
- Count lines in each file; define the empty-file and trailing-newline semantics.
- Duplicate paths are read once and reported once.
- Fail fast and return no partial map.
- Cancel and join on error.
- Never mutate a map concurrently without synchronization.
- Reject invalid worker counts.

Present two designs: workers send immutable results to one collector, or workers mutate a mutex-protected map. Implement one and justify it.
