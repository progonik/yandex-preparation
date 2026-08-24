# YR02 - Least-loaded client-side balancer

Source: [Yandex-tagged sobes.tech report](https://sobes.tech/en/bank/go/d3c5335d-aedc-4b3d-b4e6-a67d3902c241).

A microservice has many independently addressed instances. Individual instances may be overloaded or unavailable. Implement a client-side balancer that itself satisfies the backend interface and chooses the instance with the fewest active requests.

```go
type Request any
type Response any

type Backend interface {
	Invoke(ctx context.Context, req Request) (Response, error)
}

// Provided by the interviewer.
type BackendImpl struct{}
func NewBackend(addr string) *BackendImpl
func (b *BackendImpl) Invoke(context.Context, Request) (Response, error)

type Balancer struct {
	// TODO
}

func NewBalancer(addrs []string) *Balancer
func (b *Balancer) Invoke(ctx context.Context, req Request) (Response, error)
```

## Reconstructed requirements

- `Balancer` must be safe for concurrent calls to `Invoke`.
- Choose a backend with the smallest number of currently active invocations.
- Selection and incrementing the selected backend's load must form one atomic decision.
- Decrement its load when the call finishes, whether it succeeds or fails.
- Do not hold the balancer mutex during the network invocation.
- Forward the caller's context unchanged to the selected backend.
- Keep unrelated backend calls concurrent.

## Clarifications to ask

1. How should ties be broken: first backend, round-robin, or randomly?
2. What should happen when `addrs` is empty?
3. Are retries required when a backend fails? The recovered core statement does not explicitly require them.
4. Must panics from a backend restore the active-load counter?
5. Can the backend set change after construction?

The central invariant is that the load counter equals the number of invocations that selected that backend and have not yet returned.
