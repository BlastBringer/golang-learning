package main  


import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"fmt"

)

type Shop struct{
	earnings int
	start bool
}







const (
	screenWidth = 400
	screenHeight = 400
)

type  Game struct{
	shop Shop 
	tick int 
	earningSpeed int
	counter int 
}
func (g *Game)Update() error{

	if g.shop.start{
		g.counter++
		if g.counter >= g.earningSpeed {
			g.shop.earnings++
			g.counter = 0
		}
	}
	return nil
}

func (g *Game)Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Earnings : %d\n Tick : %d", g.shop.earnings, g.tick))
}



func(g *Game)Layout(outsideWidth, outsideHeight int) (int, int){
	return screenWidth, screenHeight
}





func main() {
	game := &Game{shop: Shop{earnings : 0, start : true} , earningSpeed: 60}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Cat shop proto")
	err := ebiten.RunGame(game); if err != nil{
		fmt.Println("Error while runing the game! :", err)
	}	

}