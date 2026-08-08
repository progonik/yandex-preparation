# H09 — MapReduce engine

Given:

    type Pair struct {
        Key   string
        Value int
    }

    type Mapper func(context.Context, string) ([]Pair, error)
    type Reducer func(context.Context, string, []int) (int, error)

Implement:

    func MapReduce(
        ctx context.Context,
        inputs []string,
        mapWorkers int,
        reduceWorkers int,
        mapper Mapper,
        reducer Reducer,
    ) (map[string]int, error)

Requirements:

- Map inputs in parallel.
- Shuffle values by key without concurrent map writes.
- Begin reduction only after mapping and shuffle finish.
- Reduce different keys concurrently.
- Fail fast across either phase.
- Bound internal queues and join every goroutine.
- Return no partial output.
- Produce deterministic reducer input order based on mapper input order.

Describe phase ownership and why a reducer cannot safely begin before all values for its key are known under this API.
