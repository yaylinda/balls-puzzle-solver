from typing import List, Tuple


class Ball:
    def __init__(self, id: str, color: str):
        self.id = id
        self.color = color


class BoardState:
    towers: List[List[Ball]]
    move: Tuple[int, int, str]

    def __init__(self, towers: List[List[Ball]], move: Tuple[int, int, str]):
        self.towers = towers
        self.move = move

    def __repr__(self):
        return f"BoardState(towers={self.towers}, move={self.move})"

    def get_towers(self) -> List[List[str]]:
        return self.towers

    def get_move(self) -> Tuple[int, int, str]:
        return self.move

    def set_towers(self, towers: List[List[str]]):
        self.towers = towers

    def set_move(self, move: Tuple[int, int, str]):
        self.move = move