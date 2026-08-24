# YR03 - Resumable large-table copier

Source: [Yandex-tagged entry on the sobes.tech Go bank](https://sobes.tech/en/bank/go?page=232).

There are two PostgreSQL systems: an OLTP production database and an analytical statistics database. Copy a roughly 10-TB `profiles` table from production to statistics. IDs can contain gaps.

```sql
CREATE TABLE profiles (
    id   SERIAL,
    data JSONB
);
```

```go
type Row []any

type Database interface {
	io.Closer

	GetMaxID(ctx context.Context) (uint64, error)
	LoadRows(ctx context.Context, minID, maxID uint64) ([]Row, error)
	SaveRows(ctx context.Context, rows []Row) error
}

func Connect(ctx context.Context, dbName string) (Database, error)

func CopyTable(fromName string, toName string, full bool) error
```

The provided database implementation may reconnect, and `SaveRows` is idempotent.

## Reconstructed basic requirements

- When `full` is true, copy the complete table.
- When `full` is false, resume from the last successfully copied position after a previous failure.
- Copy in bounded ID ranges rather than loading the whole table.
- Correctly handle gaps in IDs.
- Propagate load/save/connect errors and close both database handles.
- The reported basic level asks for a sequential implementation and recovery behaviour.
- The candidate may extend the interface or use `database/sql` if additional checkpoint operations are necessary.

## Concurrency practice extension

The following extension is useful preparation but was not explicitly present in the recovered basic statement:

- Load several non-overlapping ID ranges concurrently with a fixed worker limit.
- Bound loaded-but-not-saved memory.
- Save safely while maintaining a durable contiguous checkpoint.
- On the first error, cancel and join all workers.
- Never advance the checkpoint past a range that has not been saved successfully.

## Clarifications to ask

1. Where is the durable resume checkpoint stored?
2. Is destination `GetMaxID` sufficient, or can the destination contain holes?
3. Does `full=true` require clearing existing destination rows first?
4. What ID-range size should be used?
5. Can `LoadRows` and `SaveRows` run concurrently on their respective database values?
6. Is the source table changing while it is copied, and what consistency snapshot is required?
