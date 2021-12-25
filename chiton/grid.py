import heapq
from pprint import pprint


class Grid:
    def __init__(self, grid):
        self.grid = grid

    @staticmethod
    def parse(payload):
        payload = payload.strip()
        lines = payload.split("\r\n")
        grid = []
        for line in lines:
            nums = [int(c) for c in line]
            grid.append(nums)
        return Grid(grid)

    def lowest_risk(self):
        n_rows = len(self.grid)
        n_cols = len(self.grid[0])
        pq = []
        for r in range(n_rows):
            for c in range(n_cols):
                heapq.heappush(pq, (float('inf'), (r, c)))
        heapq.heappush(pq, (0, (0, 0)))
        visited = [[False for _ in row] for row in self.grid]
        dist = [[float('inf') for _ in row] for row in self.grid]
        dist[0][0] = 0

        while len(pq) > 0:
            _, u = heapq.heappop(pq)
            visited[u[0]][u[1]] = True
            for (dr, dc) in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
                v = (u[0] + dr, u[1] + dc)
                if v[0] < 0 or v[0] >= n_rows or v[1] < 0 or v[1] >= n_cols:
                    continue
                if visited[v[0]][v[1]]:
                    continue
                alt = dist[u[0]][u[1]] + \
                    self.grid[v[0]][v[1]]
                if alt < dist[v[0]][v[1]]:
                    dist[v[0]][v[1]] = alt
                    heapq.heappush(pq, (alt, (v[0], v[1])))

        return dist[-1][-1]

    def expand_grid(self):
        n_rows = len(self.grid)
        n_cols = len(self.grid[0])
        expanded_grid = [[None for _ in range(5*n_cols)] for _ in range(5*n_rows)]
        for tr in range(5):
            for tc in range(5):
                for r in range(n_rows):
                    for c in range(n_cols):
                        expanded_grid[tr*n_rows+r][tc*n_cols+c] = ((tr+tc+self.grid[r][c]) - 1) % 9 + 1
        self.grid = expanded_grid

    def print_grid(self):
        n_rows = len(self.grid)
        n_cols = len(self.grid[0])
        for r in range(n_rows):
            for c in range(n_cols):
                print(f"{self.grid[r][c]}", end='')
            print()
