"""
Part 2:

No idea. It's brute force time, baby.
"""


from typing import List, Tuple


class Simulation:
    def __init__(self, v_0: Tuple[int], bounds: Tuple[int]):
        self.v_0 = v_0
        self.bounds = bounds

    def run(self):
        p = (0, 0)
        v = self.v_0
        entries = [(0, 0)]
        _, max_x, min_y, _ = self.bounds
        while True:
            p = (p[0] + v[0], p[1] + v[1])
            v = (max(0, v[0]-1), v[1]-1)
            entries.append(p)
            if p[0] > max_x or p[1] < min_y:
                break
        return entries

    def within_bounds(self, entries: List[Tuple[int]]) -> bool:
        min_x, max_x, min_y, max_y = self.bounds
        for p in entries[::-1]:
            if p[0] >= min_x and p[0] <= max_x and p[1] >= min_y and p[1] <= max_y:
                return True
        return False


def find_valid_velocities(bounds: Tuple[int], threshold=1000):
    valid_pairs = []
    for vx in range(0, threshold):
        for vy in range(-threshold, threshold):
            s = Simulation((vx, vy), bounds)
            entries = s.run()
            if s.within_bounds(entries):
                valid_pairs.append((vx, vy))
    return valid_pairs
