import re

input_re = re.compile('^[v|>|\.]+(\n[v|>|\.]+)*$')


class ParsingException(Exception):
    pass


def parse_grid(payload: str):
    m = input_re.match(payload)
    if m == None:
        raise ParsingException
    grid = [[c for c in row] for row in payload.strip().split('\n')]
    # Verify grid dimensions
    n_cols = len(grid[0])
    for row in grid:
        if len(row) != n_cols:
            raise ParsingException
    return grid
