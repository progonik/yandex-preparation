# Yandex Go concurrency preparation

This repository contains 60 standalone production-style Go concurrency exercises:

- [20 easy problems](problems/easy/README.md)
- [20 medium problems](problems/medium/README.md)
- [20 hard problems](problems/hard/README.md)
- [10-day schedule](SCHEDULE.md)

The set targets the Go coding section: goroutines, channels, package sync, context cancellation, bounded concurrency, shutdown, races, and goroutine leaks. It intentionally contains problem statements but no solutions.

## Practice protocol

For every problem:

1. Start a timer and read only the statement.
2. Spend the first 3–5 minutes clarifying semantics and edge cases.
3. State an architecture before coding.
4. Write the first version without compiler assistance.
5. Explain channel ownership, cancellation, synchronization, and complexity.
6. Only then compile and run tests.
7. Always finish with:

       go test ./path/to/problem -race

Create solution.go and solution_test.go inside the problem directory. Use package problem unless the statement says otherwise.

## Timeboxes

| Level | First attempt | Review and tests |
|---|---:|---:|
| Easy | 30 minutes | 15 minutes |
| Medium | 45 minutes | 20 minutes |
| Hard | 60 minutes | 30 minutes |

The schedule is intentionally intensive: approximately 4.5–6 hours per day. If time is limited, finish four tasks and analyze the remaining two verbally.

## Interview checklist

Before submitting any solution, verify:

- Is the maximum concurrency actually bounded?
- Who creates and who closes every channel?
- Can any send or receive block forever after cancellation?
- Are all goroutines joined before the public function returns?
- Is every shared variable synchronized?
- Does the code preserve required ordering?
- Which error wins when several operations fail?
- What happens for empty input, invalid limits, and an already-cancelled context?
- What are the time and space complexities?
