package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const(
	screenHeight = 400
	screenWidth = 600
)
type Game struct{
	
}

func(g *Game) Update()error {
	return nil 
}

func (g *Game) Draw(screen *ebiten.Image){
	ebitenutil.DebugPrint(screen, "Hello world!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int)(int, int){
	return screenWidth, screenHeight
}

func main(){
	game := &Game{}
	ebiten.SetWindowTitle("Game")
	ebiten.SetWindowSize(screenWidth, screenHeight)
	err := ebiten.RunGame(game); if err != nil{
		panic(err)
	}
}