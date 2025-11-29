package main 

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	_ "image/png"
 "image/color"
)

const (
	gridSize = 40
	screenH = 600
	screenW = 800
)


type Game struct{
	camX float64
	camY float64
}

func(g *Game) Update () error{
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {g.camX += 4}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {g.camX -= 4}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {g.camY += 4}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {g.camY -= 4}
	return nil 
}
func (g *Game)Draw(screen *ebiten.Image){
	screen.Fill(color.RGBA{20,20,20,255})
	offX := int(g.camX) % gridSize 
	offY := int(g.camY) % gridSize 

	for x := -offX; x < screenW; x += gridSize{
		for y := -offY ; y < screenH; y += gridSize{
			rect := ebiten.NewImage(gridSize-1, gridSize-1)
			rect.Fill(color.RGBA{60,60,60,255})
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(rect, op)
		}
	}


}

func (g *Game)Layout(outsideWidth , outsideHeight int) (int, int){
	return screenW, screenH
}

func main(){

	ebiten.SetWindowTitle("flappy fish prototype!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}