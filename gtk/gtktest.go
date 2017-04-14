package main

import "github.com/mattn/go-gtk/gtk"

func main() {
	gtk.Init(nil)
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	win.SetTitle("空のウィンドウ")
	win.SetSizeRequest(400, 300)
	win.Connect("destroy", gtk.MainQuit)

	lab := gtk.NewLabel("このツールはgo-gtkのテストです。")
	win.Add(lab)

	win.ShowAll()
	gtk.Main()
}
