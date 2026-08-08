# E03 — First successful result

Implement:

    type Query func(context.Context) (string, error)

    func First(ctx context.Context, queries []Query) (string, error)

Run all queries concurrently and return the first successful value.

Requirements:

- Cancel remaining queries after the first success.
- If every query fails, return any one of their errors.
- Empty input returns a documented sentinel error.
- External cancellation returns ctx.Err().
- Assume every Query respects context cancellation.
- Wait for every started goroutine before returning.

Discuss why an unbuffered result channel can leak goroutines and what buffer capacity is sufficient.
