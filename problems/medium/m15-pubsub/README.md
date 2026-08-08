# M15 — Pub/sub broker

Implement:

    type Broker[T any] struct { /* your fields */ }

    func NewBroker[T any]() *Broker[T]
    func (b *Broker[T]) Subscribe(buffer int) (id uint64, messages <-chan T, err error)
    func (b *Broker[T]) Unsubscribe(id uint64)
    func (b *Broker[T]) Publish(value T) error
    func (b *Broker[T]) Close()

Requirements:

- Publish delivers to every currently subscribed client.
- A slow subscriber must not block other subscribers; drop messages for a full subscriber buffer.
- Unsubscribe closes that subscriber's channel.
- Close is idempotent, closes all subscriptions, and rejects future calls.
- Publish, Unsubscribe, and Close may run concurrently without panic.
- Do not hold a global lock during a potentially blocking send.

Resolve the race between copying subscriber references and concurrently closing their channels.
