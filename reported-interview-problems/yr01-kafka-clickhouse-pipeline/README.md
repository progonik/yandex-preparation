# YR01 — Acknowledged batch pipeline

Suggested time: 60 minutes.

## Context

A service reads records from a message broker and writes them to an analytical database. The broker returns relatively small batches. The database works more efficiently with larger batches, but accepts no more than `MaxItems` records in one request.

Together with every source batch, the broker returns an opaque `cookie`. Committing that cookie confirms that the corresponding source records have been handled successfully. After a restart, reading continues from the last committed position.

## Provided API

```go
const MaxItems = 9999

type Producer interface {
	Next() (items []any, cookie int, err error)
	Commit(cookie int) error
}

type Consumer interface {
	Process(items []any) error
}

func Pipe(p Producer, c Consumer) error
```

For this practice version:

- `Next` returns `io.EOF` when a finite input is exhausted. The returned `items` and `cookie` are invalid whenever `err != nil`.
- A successful `Next` call returns at most `MaxItems` items.
- The producer does not modify a returned `items` slice after `Next` returns.
- At most one call to each individual method may be active at a time. A `Next` call and a `Commit` call may run concurrently with each other.
- `Process` and `Commit` may be slow, but eventually return.

## Part 1 — Correct batching

Implement `Pipe`.

Requirements:

- Read source batches until `Next` returns `io.EOF`.
- Combine small source batches before passing records to `Process`.
- Never pass more than `MaxItems` records to one `Process` call.
- Preserve record order.
- Commit every cookie exactly once and in the order returned by `Next`.
- Commit a cookie only after all records returned with that cookie have been processed successfully.
- If the final aggregate is non-empty, process it before returning.
- Return errors from `Next`, `Process`, and `Commit`.
- Do not commit data whose processing failed.

### Example

Suppose the producer returns:

```text
Next -> [A, B, C], cookie 100
Next -> [D, E],    cookie 101
Next -> io.EOF
```

One valid call sequence is:

```text
Process [A, B, C, D, E]
Commit 100
Commit 101
```

## Part 2 — Pipeline slow operations

Assume `Next`, `Process`, and `Commit` are independent slow network operations. Change the implementation so work in different phases can overlap. For example, the producer should be able to read later records while an earlier group is being processed.

Additional requirements:

- Keep all queues bounded; memory usage must not grow with the length of the stream.
- Preserve every batching and cookie rule from Part 1.
- Do not call `Process` concurrently with itself or `Commit` concurrently with itself.
- When any stage fails, stop accepting new work and return the error that triggered shutdown.
- Cancellation must be observed by goroutines blocked on internal channel operations.
- Wait for every goroutine started by `Pipe` before returning.
- Do not leak goroutines on success or failure.

## Clarifying questions

Ask these before coding if the interviewer has not specified them:

1. May one source batch be split across multiple `Process` calls?
2. What should happen when `Next` returns an empty successful batch with a cookie?
3. Can two stages fail at nearly the same time, and which error wins?
4. Is `Process` idempotent if the service crashes after processing but before committing?
5. What shutdown guarantee is possible if a method blocks forever and the API has no context?

## What the interviewer will test

- Exact batching and commit semantics.
- Clear ownership of slices and channels.
- Safe error propagation and shutdown.
- Correct channel closure and goroutine joining.

## Provenance

Reconstructed from a [Yandex-tagged candidate report](https://sobes.tech/en/bank/go/f5db5534-99f2-4d3e-b26f-8481e1414783) and an [alternate report of the same task](https://sobes.tech/en/bank/go/b1078edb-700a-4fc6-a489-5050e83fc792). The explicit `io.EOF` and concurrency-safety assumptions were added to make the practice version deterministic.
