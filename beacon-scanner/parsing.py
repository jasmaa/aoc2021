import re
from typing import List

header_re = re.compile('^--- scanner (\d+) ---$')
point_re = re.compile('^(-?\d+),(-?\d+),(-?\d+)$')


def parse_scanners(payload: str) -> List[List[int]]:
    scanners = []
    curr_scanner = None
    lines = payload.split('\n')
    is_parsing_scanner = False
    for l in lines:
        if not is_parsing_scanner and header_re.match(l) != None:
            curr_scanner = []
            is_parsing_scanner = True
        elif is_parsing_scanner:
            m = point_re.match(l)
            if m != None:
                p = [int(m.group(1)), int(m.group(2)), int(m.group(3))]
                curr_scanner.append(p)
            elif len(l) == 0:
                scanners.append(curr_scanner)
                is_parsing_scanner = False
    scanners.append(curr_scanner)
    return scanners
