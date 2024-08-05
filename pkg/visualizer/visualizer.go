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
	States              []*models.Board
	CurrentIndex        int
	UnitSize            int
	ButtonWidth         int
	ScreenHeight        int
	ScreenWidth         int
	SecondRowTowerIndex int
	SecondRowOffset     float32
}

func (g *Visualizer) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if y > g.ScreenHeight-g.UnitSize {
			if x < g.ScreenWidth/2 {
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

			xPos := x
			yPos := float32(y)
			if x >= g.SecondRowTowerIndex {
				xPos = x % g.SecondRowTowerIndex
				yPos += g.SecondRowOffset
			}

			vector.DrawFilledRect(
				screen,
				float32(xPos*g.UnitSize),
				yPos*float32(g.UnitSize),
				float32(g.UnitSize),
				float32(g.UnitSize),
				cellColor,
				false,
			)
		}
	}

	textPaddingOffset := float32(g.UnitSize) / 4

	// Draw buttons
	vector.DrawFilledRect(
		screen,
		0,
		float32(g.ScreenHeight-g.UnitSize),
		float32(g.ButtonWidth),
		float32(g.UnitSize),
		color.RGBA{R: 200, G: 200, B: 200, A: 255},
		false,
	)
	text.Draw(
		screen,
		"Previous",
		gameFont,
		int(textPaddingOffset),
		g.ScreenHeight-int(textPaddingOffset),
		color.Black,
	)

	vector.DrawFilledRect(
		screen,
		float32(g.ScreenWidth-g.ButtonWidth),
		float32(g.ScreenHeight-g.UnitSize),
		float32(g.ButtonWidth),
		float32(g.UnitSize),
		color.RGBA{R: 200, G: 200, B: 200, A: 255},
		false,
	)
	text.Draw(
		screen,
		"Next",
		gameFont,
		g.ScreenWidth-g.ButtonWidth+int(textPaddingOffset),
		g.ScreenHeight-int(textPaddingOffset),
		color.Black,
	)

	// Draw step number
	text.Draw(
		screen,
		fmt.Sprintf("Step %d / %d", g.CurrentIndex+1, len(g.States)),
		gameFont,
		0,
		g.ScreenHeight-(2*g.UnitSize),
		color.White,
	)
}

func (g *Visualizer) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
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
