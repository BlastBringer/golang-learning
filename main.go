package main

import (
	"image"
	"os"

	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"fmt"
	"image/color"
	_ "image/png"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	screenH = 300
	screenW = 400
	SampleText = "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"
)


func increment(g *Game){
	g.fontSize++
	if g.fontSize >= 60{
		g.fontSize = 60
	}
	g.rebuildFont()
}

func decrement(g *Game){
	g.fontSize-- 
	if g.fontSize < 8{
		g.fontSize = 8
	}
	g.rebuildFont()
}


func NewGame() *Game{
	data, _ := os.ReadFile("ShinyCrystal-Yq3z4.ttf")
	g := &Game{
		fontSize: 16,
		ttfData : data,
	}
	g.rebuildFont()
	return g
}

type Game struct{
	debugUI debugui.DebugUI 
	fontSize int 
	fontFace font.Face 
	ttfData []byte
}

func (g *Game) rebuildFont(){
	tt , _ := opentype.Parse(g.ttfData)
	face , _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size : float64(g.fontSize),
		DPI : 72, 
		Hinting : font.HintingFull,

	})
	g.fontFace = face
}

func(g *Game) Update () error{
	_, err := g.debugUI.Update(func(ctx *debugui.Context) error{
		ctx.Window("fontsize", image.Rect(10,10,200,100), func(layout debugui.ContainerLayout){
			ctx.Text(fmt.Sprintf("Value : %d", g.fontSize))
			ctx.Button("increment").On(func(){increment(g)})
			ctx.Button("decrement").On(func(){decrement(g)})
		})

		return nil 
	})

	return err
}



func (g *Game)Draw(screen *ebiten.Image){
	screen.Fill(color.Black)
	
	
	g.debugUI.Draw(screen)
	text.Draw(screen, SampleText, g.fontFace, 50 , 100, color.White)

	

}

func (g *Game)Layout(outsideWidth , outsideHeight int) (int, int){
	return screenW, screenH
}

func main(){
	game := NewGame()	
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Font size changer!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}