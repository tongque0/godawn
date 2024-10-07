// 模块名称: 资源加载器
// 模块功能: 加载静态资源，供其他模块使用
// 模块接口: LoadAssets(assets embed.FS)
package loader

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

var Assets embed.FS

func LoadAssets(assets embed.FS) {
	Assets = assets
	//遍历资源
	err := fs.WalkDir(Assets, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
