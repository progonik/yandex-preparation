# C02 - Duplicate suppression

Target: 60 minutes without a compiler.

Implement a group that coalesces concurrent loads for the same key:

```go
type Loader func(context.Context, string) (string, error)

type Group struct {
    // your fields
}

func NewGroup(loadCtx context.Context, loader Loader) *Group

func (g *Group) Do(
    waitCtx context.Context,
    key string,
) (value string, shared bool, err error)
```

`loadCtx` controls the lifetime of shared loader calls. `waitCtx` controls only how long one caller waits for a result. Cancelling one caller's `waitCtx` must not cancel the shared load.

Assume `loadCtx` and `loader` are non-nil and that `loader` eventually returns after `loadCtx` is cancelled.

## Requirements

- At most one `loader(loadCtx, key)` call may be active for a key.
- Calls for different keys may load concurrently.
- A caller that finds an active load joins it instead of calling `loader` again.
- The caller that starts a load waits through the same completion-notification mechanism as followers.
- A caller whose `waitCtx` finishes first returns zero value, `false`, and `waitCtx.Err()`.
- Such cancellation does not affect the loader or other callers.
- `shared` is `true` for callers receiving a completed result if at least one duplicate caller joined that active call before it completed. A duplicate still counts if it later cancels its own wait.
- The loader's value and error are published to all callers that remain waiting.
- Once a load completes, remove its entry. A later `Do` for the same key must start a new loader call.
- Never call `loader` while holding the group map mutex.
- Do not use `golang.org/x/sync/singleflight`.

## Example

If five goroutines call `Do(ctx, "user:42")` before the loader finishes:

- the loader is called once for `"user:42"`;
- all non-cancelled callers receive the same value and error;
- those callers receive `shared == true`.

After that load completes, another call for `"user:42"` starts a fresh load and receives `shared == false` if nobody joins it.

## Questions to answer before coding

1. Which fields of an active call are written before completion is announced?
2. Why is closing a completion channel useful as a broadcast?
3. Who is allowed to remove an entry from the map?
4. Can a completed old load accidentally delete a newer entry for the same key?
