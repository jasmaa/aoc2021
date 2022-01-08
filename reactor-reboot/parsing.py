import re
from typing import List, Tuple


class ParsingException(Exception):
    pass


input_re = re.compile(
    '^(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)$'
)


def parse_cuboids(lines: List[str]) -> List[Tuple[bool, Tuple[int, int], Tuple[int, int], Tuple[int, int]]]:
    cuboids = []
    for l in lines:
        m = input_re.match(l)
        if m == None:
            raise ParsingException
        is_turn_on = m.group(1) == 'on'
        x1 = int(m.group(2))
        x2 = int(m.group(3))
        y1 = int(m.group(4))
        y2 = int(m.group(5))
        z1 = int(m.group(6))
        z2 = int(m.group(7))
        cuboids.append((is_turn_on, (x1, x2), (y1, y2), (z1, z2)))
    return cuboids
