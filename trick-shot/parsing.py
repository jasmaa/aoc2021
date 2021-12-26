from typing import Tuple
import re

bounds_re = re.compile(
    '^target area: x=(-?\d+)\.\.(-?\d+), y=(-?\d+)\.\.(-?\d+)$'
)


def parse_bounds(payload: str) -> Tuple[int]:
    m = bounds_re.match(payload)
    return (int(m.group(1)), int(m.group(2)), int(m.group(3)), int(m.group(4)))
