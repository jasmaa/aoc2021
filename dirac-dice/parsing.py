import re
from typing import List, Tuple


class ParsingException(Exception):
    pass


class InvalidInputException(Exception):
    pass


input_re = re.compile(
    '^\s*Player 1 starting position: (\d+)\nPlayer 2 starting position: (\d+)\s*$'
)


def parse_starting_positions(payload: str) -> Tuple[int, int]:
    m = input_re.match(payload)
    if m == None:
        raise ParsingException
    player1_pos = int(m.group(1))
    player2_pos = int(m.group(2))
    if player1_pos not in range(1, 11) or player2_pos not in range(1, 11):
        raise InvalidInputException
    return player1_pos, player2_pos
