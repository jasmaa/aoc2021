import copy
from typing import List

class Simulation:
    def __init__(self, grid: List[List[str]]):
        self.grid = grid

    def step(self) -> bool:
        n_rows = len(self.grid)
        n_cols = len(self.grid[0])
        is_moved = False
        # Move east herd
        new_grid = copy.deepcopy(self.grid)
        for r in range(n_rows):
            for c in range(n_cols):
                if self.grid[r][c] == '>':
                    new_c = (c + 1) % n_cols
                    if self.grid[r][new_c] == '.':
                        new_grid[r][c] = '.'
                        new_grid[r][new_c] = '>'
                        is_moved = True
        self.grid = new_grid
        # Move south herd
        new_grid = copy.deepcopy(self.grid)
        for r in range(n_rows):
            for c in range(n_cols):
                if self.grid[r][c] == 'v':
                    new_r = (r + 1) % n_rows
                    if self.grid[new_r][c] == '.':
                        new_grid[r][c] = '.'
                        new_grid[new_r][c] = 'v'
                        is_moved = True
        self.grid = new_grid
        return is_moved

    def run(self) -> int:
        """Run simulation. Returns first step where no cucumbers moved
        """
        is_moved = True
        n_step = 0
        while is_moved:
            is_moved = self.step()
            n_step += 1
        return n_step
