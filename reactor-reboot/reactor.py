from typing import List, Tuple


def find_intersection_1d(r1, r2):
    """Find intersection of two 1D inclusive ranges
    """
    if r2[0] < r1[0]:
        temp = r2
        r2 = r1
        r1 = temp
    is_intersecting = r2[0] <= r1[1]
    return (r2[0], min(r1[1], r2[1])), is_intersecting


def find_intersection_3d(s1, s2):
    """Find intersection of two 3D cuboids
    """
    rx, is_intersecting_x = find_intersection_1d(s1[0], s2[0])
    ry, is_intersecting_y = find_intersection_1d(s1[1], s2[1])
    rz, is_intersecting_z = find_intersection_1d(s1[2], s2[2])
    is_intersecting = is_intersecting_x and is_intersecting_y and is_intersecting_z
    return (rx, ry, rz), is_intersecting


def get_cube_count(s):
    """Get number of cubes in cuboid
    """
    return (s[0][1]-s[0][0]+1)*(s[1][1]-s[1][0]+1)*(s[2][1]-s[2][0]+1)


class Reactor:
    def __init__(self, cuboids: List[Tuple[bool, Tuple[int, int], Tuple[int, int], Tuple[int, int]]]):
        self.cuboids = cuboids

    def initialize(self):
        count = 0
        grid = [[[False for _ in range(101)]
                 for _ in range(101)] for _ in range(101)]
        x_offset = 50
        y_offset = 50
        z_offset = 50
        for is_turn_on, (x1, x2), (y1, y2), (z1, z2) in self.cuboids:
            if x1 >= -50 and x2 <= 50 and y1 >= -50 and y2 <= 50 and z1 >= -50 and z2 <= 50:
                for x in range(x1+x_offset, x2+x_offset+1):
                    for y in range(y1+y_offset, y2+y_offset+1):
                        for z in range(z1+z_offset, z2+z_offset+1):
                            if is_turn_on and not grid[x][y][z]:
                                count += 1
                            elif not is_turn_on and grid[x][y][z]:
                                count -= 1
                            grid[x][y][z] = is_turn_on
        return count

    def reboot(self):
        # Iteratively place cuboids
        placed = []
        for op1, rx, ry, rz in self.cuboids:
            s1 = (rx, ry, rz)
            buffer = []
            # Only place if turning on
            if op1:
                buffer.append((op1, s1))
            for op2, s2 in placed:
                # Add additional cuboid to correct for double-count on intersection
                s3, is_intersecting = find_intersection_3d(s1, s2)
                if is_intersecting:
                    buffer.append((not op2, s3))
            placed = placed + buffer
        # Calculate total cube count
        v = 0
        for op, s in placed:
            if op:
                v += get_cube_count(s)
            else:
                v -= get_cube_count(s)
        return v
