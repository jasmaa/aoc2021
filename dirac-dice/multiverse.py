from typing import List, Tuple
from functools import lru_cache


def calculate_roll_distribution() -> List[int]:
    dist = [0 for _ in range(10)]
    for i in range(1, 4):
        for j in range(1, 4):
            for k in range(1, 4):
                dist[i+j+k] += 1
    return dist


def calculate_all_score_copies() -> List[List[int]]:
    roll_dist = calculate_roll_distribution()
    all_score_copies = []
    for pos in range(11):
        score_copies = [0 for _ in range(11)]
        for i, d in enumerate(roll_dist):
            score_copies[(pos + i - 1) % 10 + 1] += d
        all_score_copies.append(score_copies)
    return all_score_copies


all_score_copies = calculate_all_score_copies()


@lru_cache(maxsize=None)
def find_multiverse_wins(player1_pos, player1_score, player2_pos, player2_score, is_player1_turn) -> Tuple[int, int]:
    if player1_score >= 21:
        return (1, 0)
    elif player2_score >= 21:
        return (0, 1)

    score_copies = all_score_copies[player1_pos if is_player1_turn else player2_pos]
    player1_wins = 0
    player2_wins = 0
    for score in range(1, 11):
        if is_player1_turn:
            w1, w2 = find_multiverse_wins(
                score, player1_score+score, player2_pos, player2_score, False
            )
        else:
            w1, w2 = find_multiverse_wins(
                player1_pos, player1_score, score, player2_score+score, True
            )
        player1_wins += score_copies[score]*w1
        player2_wins += score_copies[score]*w2
    return player1_wins, player2_wins