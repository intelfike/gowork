package main

import (
	"fmt"
	"go/token"
	"go/types"
	"log"
	"os"
)

func main() {
	tv, err := types.Eval(token.NewFileSet(), types.NewPackage("main", "main"), token.NoPos, os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tv.Value)
}
