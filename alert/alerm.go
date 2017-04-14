package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	f := false
	for {
		time.Sleep(time.Second * 5)
		now := time.Now()
		minute := now.Minute()
		switch minute {
		case 0:
			if f {
				continue
			}
			fmt.Print("\a")
			time.Sleep(time.Second >> 5)
			fallthrough
		case 30:
			if f {
				continue
			}
			log.Print("alert!\a")
			f = true
		default:
			f = false
		}
	}
}
