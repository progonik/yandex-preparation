# E10 — Once-only configuration

Implement a lazily initialized configuration holder:

    type Loader func(context.Context) (Config, error)

    type ConfigStore struct { /* your fields */ }

    func (s *ConfigStore) Get(
        ctx context.Context,
        load Loader,
    ) (Config, error)

The first caller performs loading; all concurrent and future callers observe the same result.

Requirements:

- Loader runs at most once, including when it returns an error.
- Concurrent callers wait for initialization.
- Every caller may cancel its own wait without cancelling the shared initialization.
- The zero value must be usable.

You may define Config. Explain why plain sync.Once is awkward when individual waiters need cancellable waits.
