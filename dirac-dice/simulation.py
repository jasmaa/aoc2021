from typing import List


class Player:
    def __init__(self, position: int):
        self.position = position
        self.score = 0

    def move(self, n: int):
        self.position = (self.position - 1 + n) % 10 + 1

    def clone(self) -> "Player":
        p = Player(self.position)
        p.score = self.score
        return p


class Simulation:
    def __init__(self, players: List[Player], winning_score: int):
        self.players = players
        self.turn = 0
        self.n_rolls = 0
        self.winning_score = winning_score

    def step(self, roll: int):
        p = self.players[self.turn]
        p.move(roll)
        self.n_rolls += 1
        if self.n_rolls % 3 == 0:
            p.score += p.position
            self.turn = (self.turn + 1) % len(self.players)

    def is_done(self):
        for p in self.players:
            if p.score >= self.winning_score:
                return True
        return False

    def deterministic_run(self):
        while not self.is_done():
            roll = self.n_rolls + 1
            self.step(roll)

    def clone(self) -> "Simulation":
        players_clone = [p.clone() for p in self.players]
        s_clone = Simulation(players_clone, self.winning_score)
        s_clone.turn = self.turn
        s_clone.n_rolls = self.n_rolls
        return s_clone


def find_loser_metric(players: List[Player], winning_score: int = 1000) -> int:
    s = Simulation(players, winning_score)
    s.deterministic_run()
    return min(p.score for p in players) * s.n_rolls
