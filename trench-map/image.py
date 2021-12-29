from typing import List, Tuple
import re


class ParsingException(Exception):
    pass


input_re = re.compile('^\s*([\.|#]{512})\n\n([\.|#|\n]+)\s*$')


def parse(payload: str) -> Tuple[List[str], List[List[str]]]:
    m = input_re.match(payload)
    if m == None:
        raise ParsingException
    raw_algo = m.group(1)
    raw_im = m.group(2)
    return [c for c in raw_algo], [[c for c in r] for r in raw_im.split('\n')]


def pad(im: List[List[str]], n: int, bg_c: str) -> List[List[str]]:
    w = len(im[0])
    new_im = []
    for _ in range(n):
        new_im.append([bg_c for _ in range(w+2*n)])
    for row in im:
        new_im.append([bg_c for _ in range(n)] +
                      row + [bg_c for _ in range(n)])
    for _ in range(n):
        new_im.append([bg_c for _ in range(w+2*n)])
    return new_im


def enhance(algo: List[str], im: List[List[str]], bg_c: str) -> Tuple[List[List[str]], str]:
    bg_c = algo[0] if bg_c == '.' else algo[0b111111111]
    im = pad(im, 2, bg_c)
    out_im = [[c for c in row] for row in im]
    h = len(im)
    w = len(im[0])
    for r in range(h):
        for c in range(w):
            n = 0
            for dr in [-1, 0, 1]:
                for dc in [-1, 0, 1]:
                    if r+dr < 0 or r+dr >= h or c+dc < 0 or c+dc >= w:
                        bit = 1 if bg_c == '#' else 0
                        n = n << 1 | bit
                    else:
                        bit = 1 if im[r+dr][c+dc] == '#' else 0
                        n = n << 1 | bit
            out_im[r][c] = algo[n]
    return out_im, bg_c


def n_enhance(algo: List[str], im: List[List[str]], n: int) -> List[List[str]]:
    bg_c = algo[0]
    for _ in range(n):
        im, bg_c = enhance(algo, im, bg_c)
    return im


def count_lit(im: List[List[str]]) -> int:
    n = 0
    for row in im:
        for c in row:
            if c == '#':
                n += 1
    return n
