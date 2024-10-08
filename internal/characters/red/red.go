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
	FacingRight      bool                     // 是否面向右边
}

func NewRedCharacter(startX, startY float64) *RedCharacter {
	return &RedCharacter{
		X:             startX,
		Y:             startY,
		Frames:        loader.LoadCharacter("red"),
		State:         "idle", // 初始状态
		CurrentFrame:  0,
		LastFrameTime: time.Now(),
		FacingRight:   false, // 初始面向左边
	}
}

func (c *RedCharacter) Draw(screen *ebiten.Image) {
	// 根据当前状态和当前帧绘制相应的动画帧
	var frameKey string
	fmt.Println(c.State, c.CurrentFrame)
	switch c.State {
	case "walk":
		frameKey = "red-" + fmt.Sprint(c.CurrentFrame%25) // 走路总共25帧
	case "jump":
		frameKey = "red-" + fmt.Sprint(25+c.CurrentFrame%18) // 跳跃总共18帧
	case "slide":
		frameKey = "red-" + fmt.Sprint(43+c.CurrentFrame%4) // 滑行总共4帧
	case "light_attack":
		frameKey = "red-" + fmt.Sprint(56+c.CurrentFrame%24) // 轻攻击总共25帧
	case "heavy_attack":
		frameKey = "red-" + fmt.Sprint(72+c.CurrentFrame%42) // 重攻击总共42帧
	case "blink":
		frameKey = "red-" + fmt.Sprint(114+c.CurrentFrame%7) // 闪现总共7帧
	default:
		frameKey = "red-0" // 默认为 idle 动画的第一帧,待机
	}

	if frame, ok := c.Frames[frameKey]; ok {
		op := &ebiten.DrawImageOptions{}
		if c.FacingRight {
			op.GeoM.Scale(-1, 1) // 镜像翻转
			op.GeoM.Translate(float64(frame.Bounds().Dx()), 0)
		}
		op.GeoM.Translate(c.X, c.Y)
		screen.DrawImage(frame, op)
	}
}

func (c *RedCharacter) Update() {
	// 处理输入
	handleRedInput(c)

	// 更新动画帧
	now := time.Now()
	if now.Sub(c.LastFrameTime) > time.Millisecond*100 { // 每100毫秒更新一帧
		c.CurrentFrame++
		c.LastFrameTime = now
	}

	// 如果正在进行攻击动画，则不处理其他输入
	if c.AttackInProgress {
		if c.State == "light_attack" && c.CurrentFrame >= 72 { // 假设 light_attack 的最后一帧索引是 71
			c.AttackInProgress = false
			c.State = "idle"
			c.CurrentFrame = 0
		} else if c.State == "heavy_attack" && c.CurrentFrame >= 114 { // 假设 heavy_attack 的最后一帧索引是 113
			c.AttackInProgress = false
			c.State = "idle"
			c.CurrentFrame = 0
		}
		return
	}

	// 更新碰撞盒的位置
	c.CollisionBox = image.Rect(int(c.X), int(c.Y), int(c.X)+50, int(c.Y)+100)
}

func handleRedInput(c *RedCharacter) {
	// 如果正在进行攻击动画，则不更新攻击状态，但仍然允许移动
	if c.AttackInProgress {
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			c.X -= 2
			c.FacingRight = false
		}
		if ebiten.IsKeyPressed(ebiten.KeyD) {
			c.X += 2
			c.FacingRight = true
		}
		if ebiten.IsKeyPressed(ebiten.KeyW) {
			c.Y -= 2
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) {
			c.Y += 2
		}
		return
	}

	// 根据按键更新角色的位置和状态
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		c.X -= 2
		c.State = "walk"
		c.FacingRight = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		c.X += 2
		c.State = "walk"
		c.FacingRight = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		c.Y -= 2
		c.State = "jump"
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		c.Y += 2
		c.State = "walk"
	}
	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		c.State = "light_attack"
		c.AttackInProgress = true
		c.CurrentFrame = 47 // light_attack 的起始帧索引
	}
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		c.State = "heavy_attack"
		c.AttackInProgress = true
		c.CurrentFrame = 72 // heavy_attack 的起始帧索引
	}

	// 如果没有按下任何移动键，则状态恢复为 idle
	if !ebiten.IsKeyPressed(ebiten.KeyA) && !ebiten.IsKeyPressed(ebiten.KeyD) && !ebiten.IsKeyPressed(ebiten.KeyW) && !ebiten.IsKeyPressed(ebiten.KeyS) && !ebiten.IsKeyPressed(ebiten.KeyJ) && !ebiten.IsKeyPressed(ebiten.KeyK) {
		c.State = "idle"
	}
}
