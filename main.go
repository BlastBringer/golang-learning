package main 

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	_ "image/png"
)

var fishimg *ebiten.Image

func init(){
	var err error 
	fishimg, _, err = ebitenutil.NewImageFromFile("fish.png")
	if err != nil{
		 log.Fatal(err)
	}
}

type Game struct{

}

func(g *Game) Update () error{
	return nil
}
func (g *Game)Draw(screen *ebiten.Image){
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.75, 0.75)

	ebitenutil.DebugPrint(screen, "hello there ebiten!")
	screen.DrawImage(fishimg, op)
}

func (g *Game)Layout(outsideWidth , outsideHeight int) (int, int){
	return 640, 480
}

func main(){

	ebiten.SetWindowTitle("flappy fish prototype!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}