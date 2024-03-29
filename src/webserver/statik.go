// statik 打包静态文件初始化

package webserver

import (
	"net/http"

	// generated by statik: statik -f -src=./statics
	_ "github.com/axiaoxin-com/pink-lady/statik"
	"github.com/rakyll/statik/fs"
)

// StatikFS statik obj
var StatikFS http.FileSystem

// InitStatikFS 初始化全局变量 StatikFS
func init() {
	var err error
	StatikFS, err = fs.New()
	if err != nil {
		panic(err)
	}
}
