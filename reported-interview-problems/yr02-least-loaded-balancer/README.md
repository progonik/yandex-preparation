# YR02 — Least-loaded client-side balancer

Suggested time: 45–60 minutes.

## Context

A service has several independently addressed backend instances. Implement a client-side balancer that exposes the same interface as one backend and sends each request to an instance with the smallest number of active requests.

## Provided API

```go
type Request any
type Response any

type Backend interface {
	Invoke(ctx context.Context, req Request) (Response, error)
}

// The implementation of BackendImpl is provided.
type BackendImpl struct {
	// unexported fields
}

func NewBackend(addr string) *BackendImpl
func (b *BackendImpl) Invoke(
	ctx context.Context,
	req Request,
) (Response, error)

type Balancer struct {
	// TODO
}

func NewBalancer(addrs []string) *Balancer
func (b *Balancer) Invoke(
	ctx context.Context,
	req Request,
) (Response, error)
```

You may add unexported types and fields. Assume `addrs` is non-empty and does not change after construction.

## Task

Implement `NewBalancer` and `(*Balancer).Invoke`.

Requirements:

- `Balancer` must be safe for concurrent use by many goroutines.
- For each call, choose a backend with the smallest number of active invocations.
- An invocation becomes active when that backend is selected and stops being active when its `Invoke` call returns.
- Selection and registration of the new active invocation must behave as one atomic operation relative to other callers.
- Restore the backend's load count on both success and error.
- Do not serialize network calls to different backends.
- Do not hold a balancer lock while calling a backend.
- Pass the caller's context and request to the selected backend unchanged.
- Return the selected backend's response and error unchanged.

Any backend with the minimum load is a valid choice when several are tied.

### Example

There are three backends with current active counts:

```text
backend A: 3
backend B: 1
backend C: 2
```

The next invocation must select backend B. While that invocation is in progress, B's active count is 2.

## Follow-up

Explain how the design would change if:

- backends could be added and removed at runtime;
- the balancer had to avoid an unhealthy backend temporarily;
- ties had to be resolved fairly rather than arbitrarily.

You do not need to implement the follow-up unless asked.

## Clarifying questions

1. Must a panic from a backend restore its load count?
2. Does a request waiting inside the backend count as active?
3. Are retries required after a backend error?
4. If retries are required, can the same request safely be executed twice?

## What the interviewer will test

- Protection of shared state.
- The boundary of the critical section.
- Cleanup on every return path.
- Whether independent slow calls remain concurrent.

## Provenance

Reconstructed from a [Yandex-tagged candidate report](https://sobes.tech/en/bank/go/d3c5335d-aedc-4b3d-b4e6-a67d3902c241).
