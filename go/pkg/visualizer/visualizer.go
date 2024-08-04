package visualizer

import (
	"fmt"
	"image/color"
	"log"

	"balls_puzzle_solver/pkg/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
)

const (
	ScreenWidth  = 560
	ScreenHeight = 300
	cellSize     = 40
	buttonWidth  = 100
	buttonHeight = 40
)

var (
	colorMap = map[string]color.RGBA{
		"brown":       {165, 42, 42, 255},
		"green":       {0, 128, 0, 255},
		"light green": {144, 238, 144, 255},
		"gray":        {128, 128, 128, 255},
		"light blue":  {173, 216, 230, 255},
		"red":         {255, 0, 0, 255},
		"yellow":      {255, 255, 0, 255},
		"pink":        {255, 192, 203, 255},
		"purple":      {128, 0, 128, 255},
		"blue":        {0, 0, 255, 255},
		"orange":      {255, 165, 0, 255},
		"dark green":  {0, 100, 0, 255},
		"":            {255, 255, 255, 255},
	}

	gameFont font.Face
)

type Visualizer struct {
	States       []*models.Board
	CurrentIndex int
}

func (g *Visualizer) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if y > ScreenHeight-buttonHeight {
			if x < ScreenWidth/2 {
				g.prevPuzzle()
			} else {
				g.nextPuzzle()
			}
		}
	}
	return nil
}

func (g *Visualizer) Draw(screen *ebiten.Image) {
	board := g.States[g.CurrentIndex]
	for x, tower := range board.Towers {
		for y, ball := range tower.Balls {
			var cellColor color.RGBA
			if ball == nil {
				cellColor = colorMap[""]
			} else {
				cellColor = colorMap[ball.Color]
			}

			vector.DrawFilledRect(
				screen,
				float32(x*cellSize),
				float32(y*cellSize),
				cellSize,
				cellSize,
				cellColor,
				false,
			)
		}
	}

	// Draw buttons
	vector.DrawFilledRect(
		screen,
		0,
		float32(ScreenHeight-buttonHeight),
		buttonWidth,
		buttonHeight,
		color.RGBA{R: 200, G: 200, B: 200, A: 255},
		false,
	)
	vector.DrawFilledRect(
		screen,
		float32(ScreenWidth-buttonWidth),
		float32(ScreenHeight-buttonHeight),
		buttonWidth,
		buttonHeight,
		color.RGBA{R: 200, G: 200, B: 200, A: 255},
		false,
	)

	text.Draw(screen, "Previous", gameFont, 10, ScreenHeight-10, color.Black)
	text.Draw(
		screen,
		"Next",
		gameFont,
		ScreenWidth-90,
		ScreenHeight-10,
		color.Black,
	)

	// Draw puzzle number
	text.Draw(
		screen,
		fmt.Sprintf("Step %d/%d", g.CurrentIndex+1, len(g.States)),
		gameFont,
		ScreenWidth/2-40,
		ScreenHeight-10,
		color.White,
	)
}

func (g *Visualizer) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *Visualizer) nextPuzzle() {
	g.CurrentIndex = (g.CurrentIndex + 1) % len(g.States)
}

func (g *Visualizer) prevPuzzle() {
	g.CurrentIndex = (g.CurrentIndex - 1 + len(g.States)) % len(g.States)
}

func init() {
	tt, err := opentype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	gameFont, err = opentype.NewFace(
		tt, &opentype.FaceOptions{
			Size:    20,
			DPI:     72,
			Hinting: font.HintingFull,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
