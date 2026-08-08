# H16 — Acknowledged message broker

Implement an in-memory work broker:

    type Message[T any] struct {
        ID      uint64
        Payload T
        Attempt int
    }

    type Broker[T any] struct { /* your fields */ }

    func NewBroker[T any](
        ackTimeout time.Duration,
        maxAttempts int,
    ) (*Broker[T], error)

    func (b *Broker[T]) Publish(ctx context.Context, value T) error
    func (b *Broker[T]) Receive(ctx context.Context) (Message[T], error)
    func (b *Broker[T]) Ack(id uint64) bool
    func (b *Broker[T]) DeadLetters() <-chan Message[T]
    func (b *Broker[T]) Close()

Requirements:

- Receive moves a message from ready to in-flight.
- Ack removes it permanently.
- Missing ack by ackTimeout requeues it and increments Attempt.
- After maxAttempts, send it to the dead-letter channel.
- Publish and ready queues are bounded; choose capacities in the constructor.
- Close wakes waiters, stops timers, and closes dead-letter output.
- Use one expiration loop, not one goroutine per delivery.

Define ownership of ready, in-flight, and dead-letter state and resolve Ack racing with expiration.
