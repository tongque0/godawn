package characters

import (
	"fmt"
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tongque0/godawn/loader"
)

// 角色:红(RedCharacter)
// 第一个角色，小红帽，起名红
type RedCharacter struct {
	Frames           map[string]*ebiten.Image // 角色的动画帧
	StartX, StartY   float64                  // 角色的初始位置
	X, Y             float64                  // 角色的当前位置
	Speed            float64                  // 角色的移动速度
	Health           int                      // 角色的当前生命值
	MaxHealth        int                      // 角色的最大生命值
	State            string                   // 角色的当前状态
	CollisionBox     image.Rectangle          // 角色的碰撞盒
	SoundEffects     map[string]string        // 角色的音效
	Direction        string                   // 角色面向的方向
	AttackPower      int                      // 角色的攻击力
	DefensePower     int                      // 角色的防御力
	Skills           map[string]interface{}   // 角色的技能列表
	CurrentFrame     int                      // 当前动画帧的索引
	LastFrameTime    time.Time                // 上一帧的时间
	AttackInProgress bool                     // 是否正在进行攻击动画
}

func NewRedCharacter(startX, startY float64) *RedCharacter {
	return &RedCharacter{
		X:             startX,
		Y:             startY,
		Frames:        loader.LoadCharacter("red"),
		State:         "idle", // 初始状态
		CurrentFrame:  0,
		LastFrameTime: time.Now(),
	}
}

func (c *RedCharacter) Draw(screen *ebiten.Image) {
	// 根据当前状态和当前帧绘制相应的动画帧
	var frameKey string
	switch c.State {
	case "walk":
		frameKey = "red-" + fmt.Sprint(c.CurrentFrame%25)
	case "bow":
		frameKey = "red-" + fmt.Sprint(25+c.CurrentFrame%9)
	case "jump":
		frameKey = "red-" + fmt.Sprint(34+c.CurrentFrame%18)
	case "slide":
		frameKey = "red-" + fmt.Sprint(52+c.CurrentFrame%4)
	case "light_attack_1":
		frameKey = "red-" + fmt.Sprint(56+c.CurrentFrame%7)
	case "light_attack_2":
		frameKey = "red-" + fmt.Sprint(63+c.CurrentFrame%8)
	case "light_attack_3":
		frameKey = "red-" + fmt.Sprint(71+c.CurrentFrame%10)
	case "heavy_attack":
		frameKey = "red-" + fmt.Sprint(80+c.CurrentFrame%42)
	case "blink":
		frameKey = "red-" + fmt.Sprint(120+c.CurrentFrame%7)
	default:
		frameKey = "red-" + fmt.Sprint(c.CurrentFrame%25)
	}

	if frame, ok := c.Frames[frameKey]; ok {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(c.X, c.Y)
		screen.DrawImage(frame, op)
	}
}

func (c *RedCharacter) Update() {
	// 如果正在进行攻击动画，则不处理其他输入
	if c.AttackInProgress {
		now := time.Now()
		if now.Sub(c.LastFrameTime) > time.Millisecond*100 { // 每100毫秒更新一帧
			c.CurrentFrame++
			c.LastFrameTime = now
		}

		// 检查攻击动画是否结束
		if c.CurrentFrame >= 63 { // 假设 light_attack_1 的最后一帧索引是 62
			c.AttackInProgress = false
			c.State = "idle"
			c.CurrentFrame = 0
		}
		return
	}

	// 根据按键更新角色的位置和状态
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		c.X -= 2
		c.State = "walk"
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		c.X += 2
		c.State = "walk"
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		c.Y -= 2
		c.State = "jump"
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		c.Y += 2
		c.State = "walk"
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		c.State = "light_attack_1"
		c.AttackInProgress = true
		c.CurrentFrame = 56 // light_attack_1 的起始帧索引
	}

	// 如果没有按下任何移动键，则状态恢复为 idle
	if !ebiten.IsKeyPressed(ebiten.KeyA) && !ebiten.IsKeyPressed(ebiten.KeyD) && !ebiten.IsKeyPressed(ebiten.KeyW) && !ebiten.IsKeyPressed(ebiten.KeyS) && !ebiten.IsKeyPressed(ebiten.KeySpace) {
		c.State = "idle"
	}

	// 更新动画帧
	now := time.Now()
	if now.Sub(c.LastFrameTime) > time.Millisecond*100 { // 每100毫秒更新一帧
		c.CurrentFrame++
		c.LastFrameTime = now
	}

	// 更新碰撞盒的位置
	c.CollisionBox = image.Rect(int(c.X), int(c.Y), int(c.X)+50, int(c.Y)+100)
}
