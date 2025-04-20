package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct{}

var (
	fillColor color.Color = color.RGBA{0xff, 0, 0, 0xff}
)

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.Key(ebiten.KeySpace)) {
		log.Println("The 'Space' key was just pressed!")
		red := uint8(rand.Int31n(255))
		green := uint8(rand.Int31n(255))
		blue := uint8(rand.Int31n(255))

		fillColor = color.RGBA{red, green, blue, 0xff}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(fillColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Fill example")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
