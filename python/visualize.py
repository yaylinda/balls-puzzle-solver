import pygame
import sys
from typing import List, Tuple

# Initialize Pygame
pygame.init()

# Constants
WIDTH, HEIGHT = 800, 600
BALL_RADIUS = 20
TOWER_WIDTH = 50
TOWER_HEIGHT = 4 * (BALL_RADIUS * 2 + 5)
TOWER_SPACING = 10

# Colors
BACKGROUND = (200, 200, 200)
TOWER_COLOR = (100, 100, 100)
COLOR_MAP = {
    "brown": (165, 42, 42),
    "green": (0, 255, 0),
    "light green": (144, 238, 144),
    "gray": (128, 128, 128),
    "light blue": (173, 216, 230),
    "red": (255, 0, 0),
    "yellow": (255, 255, 0),
    "pink": (255, 192, 203),
    "purple": (128, 0, 128),
    "blue": (0, 0, 255),
    "orange": (255, 165, 0),
    "dark green": (0, 100, 0),
    "": BACKGROUND  # Empty slots are the same as the background
}

# Set up the display
screen = pygame.display.set_mode((WIDTH, HEIGHT))
pygame.display.set_caption("Ball Sorting Puzzle Visualizer")


def draw_state(state: List[List[str]]):
    screen.fill(BACKGROUND)

    for i, tower in enumerate(state):
        x = i * (TOWER_WIDTH + TOWER_SPACING) + TOWER_SPACING
        y = HEIGHT - TOWER_HEIGHT - TOWER_SPACING

        # Draw tower
        pygame.draw.rect(screen, TOWER_COLOR, (x, y, TOWER_WIDTH, TOWER_HEIGHT))

        # Draw balls
        for j, ball in enumerate(tower):
            if ball:
                ball_y = y + TOWER_HEIGHT - (j + 1) * (BALL_RADIUS * 2 + 5)
                pygame.draw.circle(screen, COLOR_MAP[ball], (x + TOWER_WIDTH // 2, ball_y), BALL_RADIUS)

    pygame.display.flip()


def visualize_solution(initial_state: List[List[str]], moves: List[Tuple[int, int]]):
    state = initial_state.copy()
    draw_state(state)
    pygame.time.wait(1000)  # Wait 1 second before starting

    for move in moves:
        from_tower, to_tower = move

        # Find the ball to move
        from_index, ball = next((i, b) for i, b in enumerate(state[from_tower]) if b)
        state[from_tower][from_index] = ""

        # Find where to place the ball
        to_index = next(i for i, b in enumerate(state[to_tower][::-1]) if not b)
        state[to_tower][3 - to_index] = ball

        draw_state(state)
        pygame.time.wait(500)  # Wait 0.5 seconds between moves

        # Handle Pygame events
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                pygame.quit()
                sys.exit()

    # Keep the final state visible until the user closes the window
    while True:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                pygame.quit()
                sys.exit()
