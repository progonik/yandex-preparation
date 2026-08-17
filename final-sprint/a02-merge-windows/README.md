# A02 - Merge maintenance windows

Target: 50 minutes.

Implement:

```go
type Interval struct {
    Start int
    End   int
}

func MergeWindows(windows []Interval) []Interval
```

Every interval is valid and satisfies `Start < End`. Return the union of all windows, sorted by `Start`, with no overlaps.

For this problem, touching windows must also be merged: `[1, 3)` and `[3, 5)` become `[1, 5)`.

The function must not modify the input slice or its elements.

## Examples

```text
input  = [[1,3), [2,6), [8,10), [10,12)]
output = [[1,6), [8,12)]

input  = [[5,7), [1,2), [2,4)]
output = [[1,4), [5,7)]

input  = []
output = []
```

## Constraints

- `0 <= len(windows) <= 200_000`
- Endpoints fit in `int`.

Target complexity: `O(n log n)` time and `O(n)` additional space, including the returned result.

Explain why comparing only with the final merged interval is sufficient.
