# M03 — Duplicate suppression

Implement:

    type Loader func(context.Context, string) (string, error)

    type Group struct { /* your fields */ }

    func (g *Group) Do(
        ctx context.Context,
        key string,
        load Loader,
    ) (value string, shared bool, err error)

Concurrent calls for the same key share one active load.

Requirements:

- Calls for different keys may load concurrently.
- Exactly one Loader call runs per key at a time.
- Waiting callers may cancel independently without cancelling the shared load.
- The leader also receives its result through the common completion mechanism.
- Remove completed calls so future requests load again.
- shared reports whether a result was observed by more than one caller.

Do not use x/sync/singleflight. Carefully reason about map locking, completion notification, and when an entry can be removed.
