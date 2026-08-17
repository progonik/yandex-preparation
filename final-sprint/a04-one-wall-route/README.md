# A04 - Shortest route with one wall removal

Target: 50 minutes.

Implement:

```go
func ShortestRoute(grid []string) int
```

The rectangular grid contains:

- `S`: the starting cell;
- `E`: the destination;
- `.`: an open cell;
- `#`: a wall.

You may move one cell up, down, left, or right. You may remove at most one wall by moving into its cell. Return the minimum number of moves from `S` to `E`, or `-1` if no route exists.

`S` and `E` each occur exactly once.

## Examples

```text
grid = [
    "S#.",
    ".#E",
    "...",
]
answer = 3
explanation: remove the wall immediately to the right of S

grid = ["S##E"]
answer = -1

grid = ["SE"]
answer = 1
```

## Constraints

- `1 <= rows, columns <= 1_000`
- `rows * columns <= 200_000`
- The grid is rectangular.

Target complexity: `O(rows * columns)` time and space.

Your visited state must distinguish reaching the same cell with the wall removal still available from reaching it after the removal has been used.
