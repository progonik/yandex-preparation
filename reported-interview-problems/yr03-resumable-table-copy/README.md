# YR03 — Resumable large-table copy

Suggested time: 60 minutes.

## Context

There are two PostgreSQL systems: a production database and an analytical database. Copy the `profiles` table from production to analytics. The table is approximately 10 TB, so it cannot be loaded into memory at once. IDs are increasing but may contain gaps.

```sql
CREATE TABLE profiles (
    id   SERIAL,
    data JSONB
);
```

The copy may be interrupted by a process crash or a database error. A later run must be able to continue without starting from the beginning.

## Provided API

```go
type Row []any

type Database interface {
	io.Closer

	GetMaxID(ctx context.Context) (uint64, error)
	LoadRows(
		ctx context.Context,
		minID uint64,
		maxID uint64,
	) ([]Row, error)
	SaveRows(ctx context.Context, rows []Row) error
}

func Connect(ctx context.Context, dbName string) (Database, error)

func CopyTable(
	ctx context.Context,
	fromName string,
	toName string,
	full bool,
) error
```

The database implementation may reconnect internally. `SaveRows` is idempotent: saving the same rows more than once does not create duplicates.

Before implementation, agree with the interviewer on how the durable copy checkpoint is read and written. You may extend the provided API with a small checkpoint interface if needed.

## Part 1 — Sequential resumable copy

Implement `CopyTable`.

Requirements:

- When `full` is `true`, copy the complete source table.
- When `full` is `false`, continue from the last durably completed position.
- Read and write bounded ID ranges; memory use must not depend on total table size.
- Correctly handle ranges containing no rows and gaps between IDs.
- Never advance the durable checkpoint beyond rows that were saved successfully.
- Respect context cancellation between operations.
- Return connection, load, save, and checkpoint errors.
- Close every database handle that was opened, including on failure.

Choose a range size and be ready to explain the trade-off.

### Example

The source contains IDs:

```text
1, 2, 5, 8, 9
```

With an ID-range size of 3, valid ranges are:

```text
[1, 3] -> rows 1, 2
[4, 6] -> row 5
[7, 9] -> rows 8, 9
```

An empty range is not the end of the table. The source maximum ID determines when the copy is complete.

## Part 2 — Bounded parallel loading

After the sequential version works, allow up to `N` source ranges to be loaded concurrently.

Additional requirements:

- Never have more than `N` active load operations.
- Bound the number and total size of loaded ranges waiting to be saved.
- Saving and checkpoint advancement must remain correct when ranges finish out of order.
- A durable checkpoint may describe only a contiguous successfully saved prefix.
- On the first error or context cancellation, stop scheduling new ranges.
- Ensure goroutines blocked on internal communication can stop.
- Wait for every goroutine before returning.

This concurrency extension is an added practice follow-up; the recovered report explicitly described the basic sequential and resume parts.

## Clarifying questions

1. Where is the durable checkpoint stored, and what exactly does it represent?
2. Does `full=true` clear or replace existing destination data?
3. Can the source table change during the copy?
4. Is a consistent database snapshot required?
5. May source loads and destination saves execute concurrently?
6. How should a copy error be combined with a later `Close` error?

## What the interviewer will test

- Bounded processing of a very large data set.
- Correct progress tracking in the presence of ID gaps.
- Failure recovery and idempotency reasoning.
- For Part 2, ordering, backpressure, cancellation, and joining.

## Provenance

Reconstructed from a [Yandex-tagged Go interview report listing](https://sobes.tech/en/bank/go?page=232). The context-aware function signature and bounded-parallel follow-up were added for practice.
