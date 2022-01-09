import re
from typing import List, Union


class ParsingException(Exception):
    pass


input_re = re.compile(
    '^#############\n#\.\.\.\.\.\.\.\.\.\.\.#\n###(A|B|C|D)#(A|B|C|D)#(A|B|C|D)#(A|B|C|D)###\n  #(A|B|C|D)#(A|B|C|D)#(A|B|C|D)#(A|B|C|D)#\n  #########$'
)


def parse_state(payload: str, room_depth: int) -> List[Union[str, List[str]]]:
    m = input_re.match(payload)
    if m == None:
        raise ParsingException

    r1 = [m.group(5), m.group(1)]
    r2 = [m.group(6), m.group(2)]
    r3 = [m.group(7), m.group(3)]
    r4 = [m.group(8), m.group(4)]

    # Add additional layers if room depth is 4
    if room_depth == 4:
        r1 = [r1[0], 'D', 'D', r1[1]]
        r2 = [r2[0], 'B', 'C', r2[1]]
        r3 = [r3[0], 'A', 'B', r3[1]]
        r4 = [r4[0], 'C', 'A', r4[1]]

    return [None, None, r1, None, r2, None, r3, None, r4, None, None]
