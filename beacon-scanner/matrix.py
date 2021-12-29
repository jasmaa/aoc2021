from typing import List, Tuple


def generate_rotation_matrices():
    vecs = [[1, 0, 0], [0, 1, 0], [0, 0, 1]]

    def generate_mappings(nums: List[int]):
        if len(nums) == 1:
            return [nums]
        possible = []
        for i in range(len(nums)):
            n = nums.pop(i)
            for endings in generate_mappings(nums):
                possible.append([n] + endings)
            nums.insert(i, n)
        return possible

    mappings = generate_mappings([0, 1, 2])
    mats = []
    for m in mappings:
        mat = [vecs[i] for i in m]
        for i in [-1, 1]:
            for j in [-1, 1]:
                for k in [-1, 1]:
                    mats.append([
                        [i*c for c in mat[0]],
                        [j*c for c in mat[1]],
                        [k*c for c in mat[2]],
                    ])
    return mats


def transform_vector(m: List[List[int]], v: List[int]):
    out = []
    for j in range(len(m[0])):
        v1 = [m[i][j] for i in range(len(m))]
        c3 = sum(c1 * c2 for c1, c2 in zip(v1, v))
        out.append(c3)
    return out


def match_points(points1, points2) -> Tuple[List[Tuple[int, int]], List[List[int]], List[int]]:
    """Find best matches bashing rotation and translation
    """
    best_rotation = None
    best_offset = None
    best_matches = []
    for m in generate_rotation_matrices():
        rotated_points2 = [transform_vector(m, p) for p in points2]
        matches, offset = match_points_rotated(points1, rotated_points2)
        if len(matches) >= len(best_matches):
            best_matches = matches
            best_rotation = m
            best_offset = offset
    return best_matches, best_rotation, best_offset


def match_points_rotated(points1, points2) -> Tuple[List[Tuple[int, int]], List[int]]:
    """Find best matches for points with set-rotation
    """
    best_matches = []
    best_offset = None
    for i in range(len(points1)):
        for j in range(len(points2)):
            dv = op(points1[i], points2[j], lambda c1, c2: c1 - c2)
            translated_points2 = [
                op(v, dv, lambda c1, c2: c1 + c2) for v in points2
            ]
            matches = match_points_translated(points1, translated_points2)
            if len(matches) >= len(best_matches):
                best_matches = matches
                best_offset = dv
    return best_matches, best_offset


def match_points_translated(points1, points2) -> List[Tuple[int, int]]:
    """Find mappings between points with set-translation.
    """
    matches = []
    for i in range(len(points1)):
        for j in range(len(points2)):
            if points1[i] == points2[j]:
                matches.append((i, j))
                break
    return matches


def op(v1: List[int], v2: List[int], f) -> List[int]:
    return [f(c1, c2) for c1, c2 in zip(v1, v2)]


def map_beacons(scanners: List[List[List[int]]], match_threshold: int = 12) -> Tuple[List[List[int]], List[List[int]]]:
    """Generates map of beacons and scanners by naively bashing possible rotations and offsets.
    """
    s0 = scanners.pop(0)
    all_points = set([list2tuple(v) for v in s0])
    all_scanners = [[0, 0, 0]]
    while len(scanners) > 0:
        for i in range(len(scanners)):
            matches, rot, offset = match_points(
                [tuple2list(t) for t in list(all_points)], scanners[i]
            )
            if len(matches) >= match_threshold:
                transformed_points = [
                    op(v, offset, lambda c1, c2: c1 + c2) for v in [transform_vector(rot, v) for v in scanners[i]]
                ]
                all_scanners.append(offset)
                additional = set([list2tuple(v) for v in transformed_points])
                all_points = all_points.union(additional)
                scanners.pop(i)
                break
        print(len(scanners), len(all_points))
        print(all_scanners)
    return [tuple2list(t) for t in list(all_points)], all_scanners


def list2tuple(l: List[int]) -> Tuple[int, int, int]:
    return (l[0], l[1], l[2])


def tuple2list(t: Tuple[int, int, int]) -> List[int]:
    return [t[0], t[1], t[2]]


def find_max_manhattan_distance(scanners: List[List[int]]) -> int:
    """Finds max Manhattan distance between two scanners given a list of scanner positions.
    """
    def l1(p1, p2):
        return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1]) + abs(p1[2]-p2[2])
    best_d = 0
    for i in range(len(scanners)):
        for j in range(i+1, len(scanners)):
            best_d = max(l1(scanners[i], scanners[j]), best_d)
    return best_d
