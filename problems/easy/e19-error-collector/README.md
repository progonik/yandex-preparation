# E19 — Collect all errors

Implement:

    type Action func(context.Context) error

    func CollectErrors(
        ctx context.Context,
        actions []Action,
    ) error

Run every action concurrently and return errors.Join for all non-nil errors.

Requirements:

- Do not cancel siblings because one action fails.
- External cancellation is passed to every action.
- Wait for every action before returning.
- Include ctx.Err() in the returned joined error if cancellation occurs.
- Empty input returns nil.
- Assume actions respect context.

Implement once with a mutex-protected error slice and once with a results channel; compare complexity and failure modes.
