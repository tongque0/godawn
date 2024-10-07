package main

import (
	"embed"

	"github.com/tongque0/godawn/cmd/game"
	"github.com/tongque0/godawn/loader"
)

// 内嵌文件系统
//
//go:embed assets/**/*
var embedAssets embed.FS

func main() {
	//加载内嵌资源
	loader.LoadAssets(embedAssets)

	//启动游戏
	game.Run()
}
