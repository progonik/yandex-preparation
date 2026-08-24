# YR01 - Kafka-to-ClickHouse acknowledged pipeline

Source: [Yandex-tagged sobes.tech report](https://sobes.tech/en/bank/go/f5db5534-99f2-4d3e-b26f-8481e1414783). An [alternate entry](https://sobes.tech/en/bank/go/b1078edb-700a-4fc6-a489-5050e83fc792) contains the same core task. This reconstruction also matches the interview recollection from August 24, 2026.

The source produces small groups of records. The destination is more efficient when given larger groups. A motivating example is reading from Kafka and writing to ClickHouse.

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

## Reconstructed requirements

- The source is conditionally infinite and returns new data on successive `Next` calls within one `Pipe` invocation.
- A `Next` result never contains more than `MaxItems` records.
- `Consumer.Process` must never receive more than `MaxItems` records.
- Combine small source results into larger destination batches without exceeding the limit.
- Every `Next` result has an opaque cookie representing source progress.
- Commit a cookie only after all items associated with it have been processed successfully.
- Call `Commit` for every returned cookie, preserving the order in which `Next` produced them.
- After a restart, the source resumes from its last committed position.
- Return errors from `Next`, `Process`, or `Commit`.

## Concurrency extension reported with the task

`Next`, `Process`, and `Commit` are slow network operations. Pipeline them so reading and batch formation can continue while an earlier batch is processed or committed.

The intended overlap is conceptually:

```text
read and aggregate batch N+2
process batch N+1
commit cookies for batch N
```

The implementation must use bounded queues, propagate the first failure, avoid blocked downstream sends, and join every goroutine it starts.

## Cookie example

```text
Next -> [A, B, C], cookie 100
Next -> [D, E],    cookie 101

Process [A, B, C, D, E]
Commit 100
Commit 101
```

Committing before successful processing can lose records after a restart. Processing successfully and crashing before commit can repeat records, so the natural guarantee is at-least-once unless the destination supplies additional deduplication or transactional semantics.

## Clarifications the recovered statement does not answer

1. Does `io.EOF`, an empty item slice, or another signal terminate the stream?
2. Are `Producer.Next` and `Producer.Commit` safe to call concurrently on the same value?
3. May `Next` reuse or mutate the returned slice after its next call?
4. Is `Consumer.Process` idempotent?
5. If the current aggregate cannot fit the next source batch, should the current aggregate be flushed without splitting the source batch?
6. With no context parameters, what termination guarantee is expected when a network method blocks forever?
7. When two stages fail concurrently, which error should be returned?
