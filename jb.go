package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/intelfike/jsonbase"
)

func main() {
	fdlist := []string{
		"./main.go",
		"./lib/module.go",
		"./lib/reg.go",
		"./empty/",
	}

	fb := jsonbase.New()
	fb.Indent = "  "
	for _, v := range fdlist {
		d, f := filepath.Split(v)       // ディレクトリ・ファイルの分割
		d = strings.Trim(d, "/")        // 余分なスラッシュを削除
		fb.ChildPath(d).Push().Value(f) // 追加
	}
	fmt.Println(fb)
}
