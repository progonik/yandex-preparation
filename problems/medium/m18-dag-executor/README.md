# M18 — DAG executor

Given:

    type Node struct {
        ID   string
        Deps []string
        Run  func(context.Context) error
    }

Implement:

    func ExecuteDAG(
        ctx context.Context,
        nodes []Node,
        workers int,
    ) map[string]error

Requirements:

- Run a node only after all dependencies succeed.
- Independent ready nodes may run concurrently.
- At most workers nodes run.
- Reject duplicate IDs, unknown dependencies, and cycles.
- If a node fails, mark all transitively dependent nodes with a skipped sentinel error.
- Unrelated branches continue.
- Cancellation marks every unstarted node with ctx.Err().

Use indegrees and reverse adjacency. Do not hold the graph-state mutex while running node code.
