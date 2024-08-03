from __future__ import annotations

from collections import deque
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
    full_towers = 0
    empty_towers = 0

    for tower in state:
        if all(ball == "" for ball in tower):
            empty_towers += 1
        elif all(ball != "" for ball in tower) and len(set(tower)) == 1:
            full_towers += 1

    return full_towers == 12 and empty_towers == 2


def solve(initial_state: List[List[str]]) -> List[Tuple[int, int]] | None:
    queue = deque([(initial_state, [])])  # (state, moves)
    visited = set()
    shortest_solution = None

    while queue:
        state, moves = queue.popleft()

        if is_solved(state):
            print(f"found a solution with {len(moves)} moves")
            if shortest_solution is None or len(moves) < len(shortest_solution):
                shortest_solution = moves
            continue  # Continue searching for potentially shorter solutions

        state_tuple = tuple(tuple(tower) for tower in state)
        if state_tuple in visited:
            continue

        visited.add(state_tuple)

        for from_tower in range(len(state)):
            for to_tower in range(len(state)):
                if from_tower != to_tower and is_valid_move(state[from_tower], state[to_tower]):
                    new_state = make_move(state, from_tower, to_tower)
                    queue.append((new_state, moves + [(from_tower, to_tower)]))

    return shortest_solution  # Returns None if no solution found


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
        print(f"optimal solution: {len(solution)} moves")
        # visualize_solution(balls, solution)
    else:
        print("No solution found")
