package characters

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	characters "github.com/tongque0/godawn/internal/characters/red"
)

// Player 结构体适配各种角色
type Player struct {
	Character Character
}

// Character 接口定义了角色的基本行为和属性
type Character interface {
	Draw(screen *ebiten.Image)
	Update()
}

func NewPlayer(characterName string, startX, startY float64) *Player {
	character := createCharacter(characterName, startX, startY)
	if character == nil {
		log.Fatalf("无法创建角色: %s", characterName)
	}
	return &Player{
		Character: character,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Character.Draw(screen)
}

func (p *Player) Update() {
	p.Character.Update()
}

// createCharacter 根据角色名创建具体的角色实例
func createCharacter(characterName string, startX, startY float64) Character {
	switch characterName {
	case "red":
		return characters.NewRedCharacter(startX, startY)
	// 可以添加更多的角色类型
	default:
		return nil
	}
}
