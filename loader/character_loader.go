// 模块名称:角色加载器
// 模块功能:加载角色的图片资源
package loader

import (
	"encoding/json"
	"image"
	"image/png"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type Frame struct {
	Filename         string `json:"filename"`         // 文件名
	Frame            Rect   `json:"frame"`            // 帧
	Rotated          bool   `json:"rotated"`          // 旋转
	Trimmed          bool   `json:"trimmed"`          // 裁剪
	SpriteSourceSize Rect   `json:"spriteSourceSize"` // 精灵源大小
	SourceSize       Size   `json:"sourceSize"`       // 源大小
}

type Rect struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

type Size struct {
	W int `json:"w"`
	H int `json:"h"`
}

type FrameData struct {
	Frames []Frame `json:"frames"`
}

func LoadCharacter(characterName string) map[string]*ebiten.Image {
	frameMap := make(map[string]*ebiten.Image)

	framesFile := "assets/characters/" + characterName + "/" + characterName + ".json"
	spriteFile := "assets/characters/" + characterName + "/" + characterName + ".png"

	// 读取 JSON 文件
	framesData, err := Assets.ReadFile(framesFile)
	if err != nil {
		log.Fatal(err)
	}

	var frameData FrameData
	err = json.Unmarshal(framesData, &frameData)
	if err != nil {
		log.Fatal(err)
	}

	// 读取 PNG 文件
	spriteData, err := Assets.ReadFile(spriteFile)
	if err != nil {
		log.Fatal(err)
	}

	spriteImg, err := png.Decode(strings.NewReader(string(spriteData)))
	if err != nil {
		log.Fatal(err)
	}

	// 创建 Ebiten 图像
	ebitenImage := ebiten.NewImageFromImage(spriteImg)

	// 提取动画帧
	for _, frame := range frameData.Frames {
		subImage := ebitenImage.SubImage(image.Rect(frame.Frame.X, frame.Frame.Y, frame.Frame.X+frame.Frame.W, frame.Frame.Y+frame.Frame.H)).(*ebiten.Image)
		filenameWithoutExt := strings.TrimSuffix(frame.Filename, ".png")
		frameMap[filenameWithoutExt] = subImage
	}

	return frameMap
}
