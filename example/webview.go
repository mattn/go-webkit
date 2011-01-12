package main

import "gtk"
import "webkit"

func main() {
	gtk.Init(nil)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("webkit")
	window.Connect("destroy", gtk.MainQuit, nil)

	webview := webkit.WebView()
	window.Add(webview)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	webview.LoadUri("http://mattn.kaoriya.net/")
	gtk.Main()
}
