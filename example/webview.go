package main

import "os"
import "gtk"
import "webkit"

func main() {
	gtk.Init(nil)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("webkit")
	window.Connect("destroy", gtk.MainQuit, nil)

	swin := gtk.ScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.GTK_POLICY_AUTOMATIC, gtk.GTK_POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.GTK_SHADOW_IN)

	webview := webkit.WebView()
	swin.Add(webview)

	window.Add(swin)
	window.SetSizeRequest(600, 600)
	window.ShowAll()

	proxy := os.Getenv("HTTP_PROXY")
	if len(proxy) > 0 {
		soup_uri := webkit.SoupUriNew(proxy)
		webkit.GetDefaultSession().Set("proxy-uri", soup_uri)
		webkit.SoupUriFree(soup_uri)
	}
	webview.LoadUri("http://mattn.kaoriya.net/")
	gtk.Main()
}
