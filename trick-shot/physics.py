"""
Part 1:

All we care about is p_y:

p_y[n+1] = p_y[n] + v_y[n]
v_y[n+1] = v_y[n] - 1

Telescope p_y[n+1]:

p_y[n] = (...((((p_y[0] + v_y[0]) + v_y[0]-1) + v_y[0]-1-1) + v_y[0]-1-1-1) + ...)
p_y[n] = p_y[0] + n*v_y[0] - sum(i-1, i=0..n)

sum(i-1, i=0..n)
= 0 + sum(i, i=1..n) - sum(1, i=0..n)
= sum(i, i=1..n) - n

Lookup sum(i, i=1..n) formula because you forgot it:
sum(i, i=0..n) = n(n+1)/2

sum(i-1, i=0..n)
= n(n+1)/2 - n
= n(n-1)/2

So:

p_y[n] = p_y[0] + n*v_y[0] - n(n-1)/2

Set p_y[0] = 0 and simplify:

p_y[n] = 0 + n*v_y[0] - n(n-1)/2
p_y[n] = n*v_y[0] - n^2/2 + n/2
p_y[n] = -n^2/2 + (v_y[0]+1/2)*n
0 = -n^2/2 + (v_y[0]+1/2)*n - p_y[n]

Since we are aiming for a range of p_y, change p_y[n] to p_y (it's a target now):

0 = -n^2/2 + (v_y[0]-1/2)*n - p_y

Know: p_y
Want: v_y[0]
Constraint: n >= 0 and integer

Use quadratic formula:

( -(v_y[0]-1/2) +/- sqrt( (v_y[0]+1/2)^2 - 4*-1/2*-p_y ) ) / (2*-1/2)
= v_y[0] + 1/2 -/+ sqrt( (v_y[0]+1/2)^2 - 2*p_y )

We can sub in y bounds for p_y and calculate n (time). If there is a positive integer
between time ranges, then there exists a step that will be in the box.
"""

import math
from typing import Tuple


def find_highest_y(bounds: Tuple[int], miss_threshold=100):
    _, _, y_min, y_max = bounds

    def p_y(t, v_0):
        return t*v_0 - t*(t-1)/2

    def time_range(v_0):
        return sorted([v_0 + 0.5 + math.sqrt((v_0+0.5)**2 - 2*p_y) for p_y in [y_min, y_max]])

    def highest_y(v_0):
        t = v_0 + 0.5
        return p_y(t, v_0)

    v_0 = 1
    entries = [(0, True, 0)]
    consecutive_misses = 0
    while True:
        min_t, max_t = time_range(v_0)
        between = (
            int(min_t) if abs(int(min_t)-min_t) < 0.00001 else int(min_t) + 1
        )
        miss = between > max_t
        if miss:
            consecutive_misses += 1
        else:
            consecutive_misses = 0
        if consecutive_misses > miss_threshold:
            break
        # Assume slow-moving at apex so can just floor max continuous value and
        # pray that discrete steps will hit the floored value at some point.
        entries.append(
            (v_0, miss, int(highest_y(v_0)))
        )
        v_0 += 1

    # Find last hit
    for v_0, miss, best_height in entries[::-1]:
        if not miss:
            return best_height
    return 0
