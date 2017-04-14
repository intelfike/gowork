package main

import (
	"time"

	"github.com/intelfike/lib/rcwindow"
)

func main() {
	rc := rcwindow.NewWindow(1, 1, 100)
	rc.FillX(
		func(x float64) (y float64) {
			y = x
			return
		},
		func() {
			time.Sleep(time.Second / 4)
		})
	rc.Redraw()
	rc.Wait()
}
