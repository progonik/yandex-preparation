# M20 — Concurrent account transfers

Implement:

    type Bank struct { /* your fields */ }

    func NewBank(initial map[string]int64) *Bank
    func (b *Bank) Transfer(from, to string, amount int64) error
    func (b *Bank) Balance(id string) (int64, bool)
    func (b *Bank) Total() int64

Requirements:

- Transfers are atomic and never expose partially updated balances.
- Reject missing accounts, non-positive amounts, and insufficient funds.
- Concurrent transfers between the same two accounts must not deadlock.
- Balance may run concurrently.
- Total returns a consistent snapshot.
- Do not serialize all transfers with one global mutex.

Use per-account locks with a deterministic acquisition order. Explain how Total locks every account without deadlocking transfers.
