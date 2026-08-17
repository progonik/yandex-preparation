# A03 - Minimum servers

Target: 40 minutes.

Implement:

```go
type Request struct {
    Start int
    End   int
}

func MinimumServers(requests []Request) int
```

Each request occupies one server during the half-open interval `[Start, End)`. A server that becomes free at time `t` may immediately accept another request starting at `t`.

Return the minimum number of servers required to handle every request without delay.

## Examples

```text
requests = [[0,30), [5,10), [10,20)]
answer = 2

requests = [[1,2), [2,3), [3,4)]
answer = 1

requests = [[1,10), [2,9), [3,8), [4,7)]
answer = 4

requests = []
answer = 0
```

## Constraints

- `0 <= len(requests) <= 200_000`
- Every request satisfies `Start < End`.
- Times may be negative.
- Do not modify the input.

Target complexity: `O(n log n)` time and `O(n)` additional space.

Be precise about what happens when one request ends at the exact time another begins.
