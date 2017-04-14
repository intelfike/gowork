package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/intelfike/tools/lf/fileexp"
)

func main() {
	defer fmt.Println("\x1b[37m")
	ch, _ := fileexp.ReadDirAll(".", 200)
	n := 0
	_ = n
	for file := range ch {
		// if file.Info.IsDir() {
		// 	fmt.Println("\x1b[33m", file.Abs())
		// 	continue
		// }
		// fmt.Println("\x1b[37m", file.Rel())
		d, _ := filepath.Split(file.Rel())
		d = strings.Trim(d, "/")

		if file.Info.IsDir() {
		} else {
			rch := file.ReadChan(10, 30)
			str := ""
			for line := range rch {
				if !strings.Contains(line, "func") {
					continue
				}
				str += line + "\n"
			}
			if len(str) == 0 {
				continue
			}
			fmt.Println(">>>", file.Rel())
			fmt.Println(str)
		}
	}
}
