package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tongque0/godawn/internal/characters"
)

type Game struct {
	player1 *characters.Player
	player2 *characters.Player
}

func (g *Game) Update() error {
	if g.player1 != nil {
		g.player1.Update()
	}
	if g.player2 != nil {
		g.player2.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.player1 != nil {
		g.player1.Draw(screen)
	}
	if g.player2 != nil {
		g.player2.Draw(screen)
	}
	// ebitenutil.DebugPrint(screen, "Go之破晓")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// 保持 16:9 的宽高比
	const aspectRatio = 16.0 / 9.0
	width := outsideWidth
	height := int(float64(outsideWidth) / aspectRatio)
	if height > outsideHeight {
		height = outsideHeight
		width = int(float64(outsideHeight) * aspectRatio)
	}
	return width, height
}

func Run() {
	game := &Game{
		player1: characters.NewPlayer("red", 800, 600),
	}
	ebiten.SetWindowSize(1280, 720) // 初始窗口大小
	// ebiten.SetWindowTitle("Go之破晓")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
