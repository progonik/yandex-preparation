# A05 - Minimum daily capacity

Target: 40 minutes.

Implement:

```go
func MinimumCapacity(durations []int, days int) int64
```

Jobs must be processed in their original order. During a day, process one contiguous group of jobs whose total duration does not exceed the daily capacity. A job cannot be split across days.

Return the minimum integer capacity that allows all jobs to be completed within at most `days` days.

## Examples

```text
durations = [7, 2, 5, 10, 8], days = 2
answer = 18
explanation: [7,2,5] and [10,8]

durations = [7, 2, 5, 10, 8], days = 5
answer = 10

durations = [5], days = 1
answer = 5
```

## Constraints

- `1 <= len(durations) <= 200_000`
- Every duration is positive.
- `1 <= days <= len(durations)`.
- The total duration fits in `int64` but may not fit in `int`.

Target complexity: `O(n log(sum(durations)))` time and `O(1)` additional space.

Explain why feasibility is monotonic: if capacity `x` is sufficient, every capacity greater than `x` is also sufficient.
