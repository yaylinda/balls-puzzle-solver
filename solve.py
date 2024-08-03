from __future__ import annotations

from typing import List, Tuple
from copy import deepcopy

from visualize import visualize_solution


def find_top_ball(tower: List[str]) -> Tuple[int, str]:
    for i, ball in enumerate(tower):
        if ball:
            return i, ball
    return -1, ""


def is_valid_move(from_tower: List[str], to_tower: List[str]) -> bool:
    if "" not in to_tower:  # Destination tower is full
        return False

    from_index, from_ball = find_top_ball(from_tower)
    if from_index == -1:  # No ball to move
        return False

    to_index, to_ball = find_top_ball(to_tower)
    return to_index == -1 or to_ball == from_ball


def make_move(state: List[List[str]], from_tower: int, to_tower: int) -> List[List[str]]:
    new_state = deepcopy(state)
    from_index, ball = find_top_ball(new_state[from_tower])
    new_state[from_tower][from_index] = ""

    for i in range(len(new_state[to_tower]) - 1, -1, -1):
        if new_state[to_tower][i] == "":
            new_state[to_tower][i] = ball
            break

    return new_state


def is_solved(state: List[List[str]]) -> bool:
    return all(len(set(tower) - {""}) <= 1 and tower.count("") == 0 for tower in state if "" not in tower)


def solve(initial_state: List[List[str]]) -> List[Tuple[int, int]] | None:
    stack = [(initial_state, [])]  # (state, moves)
    visited = set()

    while stack:
        state, moves = stack.pop()

        if is_solved(state):
            return moves

        state_tuple = tuple(tuple(tower) for tower in state)
        if state_tuple in visited:
            continue

        visited.add(state_tuple)

        for from_tower in range(len(state)):
            for to_tower in range(len(state)):
                if from_tower != to_tower and is_valid_move(state[from_tower], state[to_tower]):
                    new_state = make_move(state, from_tower, to_tower)
                    stack.append((new_state, moves + [(from_tower, to_tower)]))

    return None  # No solution found


if __name__ == "__main__":
    balls = [
        ["brown", "green", "brown", "light green"],
        ["light green", "gray", "light blue", "green"],
        ["red", "green", "light green", "yellow"],
        ["pink", "red", "yellow", "purple"],
        ["light green", "pink", "red", "green"],
        ["gray", "blue", "pink", "gray"],
        ["orange", "blue", "yellow", "dark green"],
        ["purple", "orange", "orange", "dark green"],
        ["light blue", "blue", "orange", "red"],
        ["purple", "dark green", "dark green", "light blue"],
        ["yellow", "pink", "purple", "blue"],
        ["light blue", "brown", "brown", "gray"],
        ["", "", "", ""],
        ["", "", "", ""],
    ]

    solution = solve(balls)

    if solution:
        print("Solution found!")
        visualize_solution(balls, solution)
    else:
        print("No solution found")
