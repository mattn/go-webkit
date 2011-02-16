package main

import "os"
import "gtk"
import "webkit"

const HTML_STRING = `
<doctype html>
<meta charset="utf-8"/>
<style>
div { font-size: 5em }
</style>
<script src="http://code.jquery.com/jquery-latest.js"></script>
<script>
$(function() {
    $('#hello1').slideDown('slow', function() {
    	$('#hello2').fadeIn()
	})
})
</script>
<div id="hello1" style="display: none">Hello</div>
<div id="hello2" style="display: none">世界</div>
</div>
`

func main() {
	gtk.Init(nil)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("webkit")
	window.Connect("destroy", gtk.MainQuit)

	vbox := gtk.VBox(false, 1)

	entry := gtk.Entry()
	entry.SetText("http://mattn.kaoriya.net/")
	vbox.PackStart(entry, false, false, 0)

	swin := gtk.ScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.GTK_POLICY_AUTOMATIC, gtk.GTK_POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.GTK_SHADOW_IN)

	webview := webkit.WebView()
	webview.Connect("load-committed", func() {
		entry.SetText(webview.GetUri())
	},
		nil)
	swin.Add(webview)

	vbox.Add(swin)

	entry.Connect("activate", func() {
		webview.LoadUri(entry.GetText())
	},
		nil)
	button := gtk.ButtonWithLabel("load String")
	button.Clicked(func() {
		webview.LoadString("hello Go GTK!", "text/plain", "utf-8", ".")
	},
		nil)
	vbox.PackStart(button, false, false, 0)

	button = gtk.ButtonWithLabel("load HTML String")
	button.Clicked(func() {
		webview.LoadHtmlString(HTML_STRING, ".")
	},
		nil)
	vbox.PackStart(button, false, false, 0)

	window.Add(vbox)
	window.SetSizeRequest(600, 600)
	window.ShowAll()

	proxy := os.Getenv("HTTP_PROXY")
	if len(proxy) > 0 {
		soup_uri := webkit.SoupUri(proxy)
		webkit.GetDefaultSession().Set("proxy-uri", soup_uri)
		soup_uri.Free()
	}
	gtk.Main()
}
