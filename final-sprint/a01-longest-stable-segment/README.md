# A01 - Longest stable segment

Target: 40 minutes.

Implement:

```go
func LongestStable(values []int, k int) int
```

A contiguous segment is stable when it contains at most `k` distinct values. Return the maximum length of a stable segment.

## Examples

```text
values = [1, 2, 1, 3, 4, 2, 3], k = 2
answer = 3
explanation: [1, 2, 1]

values = [5, 5, 5], k = 1
answer = 3

values = [1, 2, 3], k = 0
answer = 0

values = [], k = 4
answer = 0
```

## Constraints

- `0 <= len(values) <= 200_000`
- `0 <= k <= 200_000`
- Values may be negative.

Target complexity: `O(n)` expected time and `O(min(n, number of distinct values))` additional space.

Explain why each array position is processed only a constant number of times.
