from typing import List, Tuple, Union


class ParsingException(Exception):
    pass


class Cons:
    def __init__(self, left: Union["Cons", int], right: Union["Cons", int], parent: Union["Cons", None]):
        self.left = left
        self.right = right
        self.parent = parent


def is_digit(c: str) -> bool:
    return c in '0123456789'


def parse_digit(payload: str, idx: int = 0) -> Tuple[int, int]:
    if is_digit(payload[idx]):
        v = 0
        while is_digit(payload[idx]):
            v = v*10 + int(payload[idx])
            idx += 1
    return v, idx


def parse(payload: str) -> Union[Cons, None]:
    c, idx = parse_cons(payload, 0)
    if idx != len(payload):
        raise ParsingException
    return c


def parse_cons(payload: str, idx: int) -> Tuple[Union[Cons, None], int]:
    if len(payload) == 0:
        return None

    if payload[idx] != '[':
        raise ParsingException
    idx += 1

    if is_digit(payload[idx]):
        l_res, idx = parse_digit(payload, idx)
    else:
        l_res, idx = parse_cons(payload, idx)

    if payload[idx] != ',':
        raise ParsingException
    idx += 1

    if is_digit(payload[idx]):
        r_res, idx = parse_digit(payload, idx)
    else:
        r_res, idx = parse_cons(payload, idx)

    if payload[idx] != ']':
        raise ParsingException
    idx += 1

    c = Cons(l_res, r_res, None)
    if isinstance(l_res, Cons):
        l_res.parent = c
    if isinstance(r_res, Cons):
        r_res.parent = c

    return c, idx


def cons_to_string(c: Union[Cons, None]) -> str:
    if c == None:
        return '[]'
    else:
        l_res = c.left if type(c.left) == int else cons_to_string(c.left)
        r_res = c.right if type(c.right) == int else cons_to_string(c.right)
        return f'[{l_res},{r_res}]'


def explode(root: Cons) -> bool:
    def find_first_level_4(c: Union[Cons, int], level: int) -> Union[Cons, None]:
        if type(c) == int:
            return None
        elif level < 4:
            l_res = find_first_level_4(c.left, level+1)
            r_res = find_first_level_4(c.right, level+1)
            return l_res if l_res != None else r_res
        else:
            return c

    c = find_first_level_4(root, 0)
    if c == None:
        return False

    # Explode left
    curr_c = c
    while curr_c.parent != None and curr_c.parent.right != curr_c:
        curr_c = curr_c.parent
    curr_c = curr_c.parent
    if curr_c != None:
        if type(curr_c.left) == int:
            curr_c.left += c.left
        else:
            curr_c = curr_c.left
            while type(curr_c.right) != int:
                curr_c = curr_c.right
            curr_c.right += c.left

    # Explode right
    curr_c = c
    while curr_c.parent != None and curr_c.parent.left != curr_c:
        curr_c = curr_c.parent
    curr_c = curr_c.parent
    if curr_c != None:
        if type(curr_c.right) == int:
            curr_c.right += c.right
        else:
            curr_c = curr_c.right
            while type(curr_c.left) != int:
                curr_c = curr_c.left
            curr_c.left += c.right

    # Replace pair with 0
    if c == c.parent.left:
        c.parent.left = 0
    else:
        c.parent.right = 0

    return True


def split(root: Cons) -> bool:
    def splittable(v) -> bool:
        return v >= 10

    def find_first_splittable_cons(c: Union[Cons, int, None]) -> Union[Cons, None]:
        if type(c) == int:
            return None
        else:
            l_res = c if type(c.left) == int and splittable(c.left) \
                else find_first_splittable_cons(c.left)
            r_res = c if type(c.right) == int and splittable(c.right) \
                else find_first_splittable_cons(c.right)
            return l_res if l_res != None else r_res

    c = find_first_splittable_cons(root)
    if c == None:
        return False

    if type(c.left) == int and splittable(c.left):
        new_c = Cons(c.left//2, c.left-c.left//2, c)
        c.left = new_c
    else:
        new_c = Cons(c.right//2, c.right-c.right//2, c)
        c.right = new_c

    return True


def add(c1: Cons, c2: Cons) -> Cons:
    c1 = parse(cons_to_string(c1))
    c2 = parse(cons_to_string(c2))
    root = Cons(c1, c2, None)
    c1.parent = root
    c2.parent = root
    took_action = True
    while took_action:
        took_action = explode(root)
        if not took_action:
            took_action = split(root)
    return root


def add_all(cons_l: List[Cons]) -> Cons:
    summed_cons = cons_l[0]
    for i in range(1, len(cons_l)):
        summed_cons = add(summed_cons, cons_l[i])
    return summed_cons


def magnitude(cons: Cons) -> int:
    v1 = cons.left if type(cons.left) == int else magnitude(cons.left)
    v2 = cons.right if type(cons.right) == int else magnitude(cons.right)
    return 3*v1+2*v2


def find_largest_twosum_magnitude(cons_l: List[Cons]) -> int:
    best_magnitude = 0
    for i in range(len(cons_l)):
        for j in range(len(cons_l)):
            if i != j:
                c = add(cons_l[i], cons_l[j])
                best_magnitude = max(magnitude(c), best_magnitude)
    return best_magnitude
