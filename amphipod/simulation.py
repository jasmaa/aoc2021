from functools import lru_cache
import copy


class UnsolvableException(Exception):
    pass


room_idxs = [2, 4, 6, 8]


type2cost = {
    'A': 1,
    'B': 10,
    'C': 100,
    'D': 1000,
}
type2idx = {
    'A': 2,
    'B': 4,
    'C': 6,
    'D': 8,
}
idx2type = {
    2: 'A',
    4: 'B',
    6: 'C',
    8: 'D',
}


def check_obstructed(state, i, j):
    if i < j:
        for k in range(i+1, j):
            if k not in room_idxs and state[k] != None:
                return True
    else:
        for k in range(i-1, j, -1):
            if k not in room_idxs and state[k] != None:
                return True
    return False


def state2tuple(state):
    for i in range(len(state)):
        if type(state[i]) == list:
            state[i] = tuple(state[i])
    return tuple(state)


def tuple2state(state):
    state = list(state)
    for i in range(len(state)):
        if type(state[i]) == tuple:
            state[i] = list(state[i])
    return state


def check_room_homogenous(room, t, n):
    for d in range(n+1):
        if room == [t] * d:
            return True
    return False


@lru_cache(maxsize=None)
def dfs(state, room_depth):
    state = tuple2state(state)

    # Check game over
    is_game_over = True
    for i in room_idxs:
        t = idx2type[i]
        if state[i] != [t]*room_depth:
            is_game_over = False
            break
    if is_game_over:
        return 0

    best_cost = float('inf')
    for i in range(len(state)):
        if i in room_idxs:
            # Check moves from room to hallway
            room_t = idx2type[i]
            if not check_room_homogenous(state[i], room_t, room_depth):
                for j in range(len(state)):
                    if state[j] == None:
                        if not check_obstructed(state, i, j):
                            moves = room_depth + 1 - len(state[i]) + abs(i - j)
                            cost = type2cost[state[i][-1]] * moves
                            new_state = copy.deepcopy(state)
                            new_state[j] = new_state[i].pop()
                            best_cost = min(
                                dfs(state2tuple(new_state), room_depth) + cost,
                                best_cost,
                            )
        elif state[i] != None:
            # Check moves from hallway to room
            room_t = state[i]
            j = type2idx[room_t]
            if check_room_homogenous(state[j], room_t, room_depth-1):
                if not check_obstructed(state, i, j):
                    moves = abs(i - j) + room_depth - len(state[j])
                    cost = type2cost[state[i]] * moves
                    new_state = copy.deepcopy(state)
                    new_state[j].append(new_state[i])
                    new_state[i] = None
                    best_cost = min(
                        dfs(state2tuple(new_state), room_depth) + cost, best_cost)

    return best_cost


def find_best_cost(state, room_depth):
    cost = dfs(state2tuple(state), room_depth)
    if cost == float('inf'):
        raise UnsolvableException
    return cost
