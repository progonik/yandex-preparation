# Final interview sprint: August 18-20, 2026

The interview is on August 21. This sprint contains exactly three problems per day:

- one 60-minute Go concurrency problem, matching the live-coding section;
- two algorithm problems with a combined 90-minute limit.

Do the problems in order. They deliberately progress from the current level to a stretch problem.

## Schedule

| Date | Go live coding, 60 min | Algorithm 1, 40 min | Algorithm 2, 50 min |
|---|---|---|---|
| Aug 18 | [C01 - Graceful worker](c01-graceful-worker/README.md) | [A01 - Longest stable segment](a01-longest-stable-segment/README.md) | [A02 - Merge maintenance windows](a02-merge-windows/README.md) |
| Aug 19 | [C02 - Duplicate suppression](c02-duplicate-suppression/README.md) | [A03 - Minimum servers](a03-minimum-servers/README.md) | [A04 - Route with one wall removal](a04-one-wall-route/README.md) |
| Aug 20 | [C03 - Per-key executor](c03-keyed-executor/README.md) | [A05 - Minimum daily capacity](a05-minimum-capacity/README.md) | [A06 - Minimum dictionary segmentation](a06-dictionary-segmentation/README.md) |

## Rules for the Go problem

1. Use a plain editor and no compiler for the first 60 minutes.
2. Spend the first five minutes clarifying semantics and writing invariants.
3. Before writing code, state:
   - goroutine ownership;
   - channel ownership and closure;
   - cancellation behaviour;
   - how every goroutine terminates;
   - which shared state requires synchronization.
4. At 60 minutes, stop and mark any unfinished sections. Only then compile, add tests, and run `go test -race`.

## Rules for the algorithm section

1. Run a single 90-minute timer for both problems.
2. For each problem, first state the brute-force approach, then derive the intended complexity.
3. Write compilable Go, including all boundary cases.
4. Finish by stating time and space complexity.

## Review protocol

After each attempt, record only four things:

- the main invariant;
- the bug or uncertainty that cost the most time;
- one missing test;
- what you would explain differently to the interviewer.

Do not replace the Aug 20 simulation with additional reading. Completing and reviewing these nine problems is higher value than opening more topics.
