arr1 = [10, 12, 15, -9, 15, 10, 14, -5, 14, -7, -12, -10, -1, -11]
arr2 = [15,  8,  2,  6, 13,  4,  1,  9,  5, 13,  9,    6,  2,   2]
arr3 = [1,   1,  1, 26,  1,  1,  1, 26,  1, 26, 26,   26, 26,  26]
ops = [False, False, False, True, False, False,
       False, True, False, True, True, True, True, True]


def monad1(model: str):
    input = [int(c) for c in model]
    z = 0
    for c, x, y, d in zip(input, arr1, arr2, arr3):
        curr = z % 26 + x
        z //= d
        if curr != c:
            z *= 26
            z += c+y
    return z


def monad2(model: str):
    input = [int(c) for c in model]
    s = []
    for should_pop, c, x, y in zip(ops, input, arr1, arr2):
        v = s[-1] if len(s) > 0 else 0
        if should_pop and len(s) > 0:
            s.pop()
        if v + x != c:
            s.append(c+y)
    # Calculate z
    acc = 0
    m = 1
    while len(s) > 0:
        d = s.pop()
        acc += d*m
        m *= 26
    return acc
