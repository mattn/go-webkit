package main

import (
	"os"
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-webkit/webkit"
)

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

const MAP_EMBED = `
<style> *{ margin : 0; padding : 0; } </style>
<iframe width="100%" height="100%" frameborder="0" scrolling="no" marginheight="0" marginwidth="0" src="http://maps.google.co.jp/maps?f=q&amp;source=s_q&amp;hl=en&amp;geocode=&amp;q=osaka&amp;aq=&amp;sll=34.885931,-115.180664&amp;sspn=29.912003,39.506836&amp;brcurrent=3,0x6000e86b2acc70d7:0xa399ff48811f596d,0&amp;ie=UTF8&amp;hq=&amp;hnear=%E5%A4%A7%E9%98%AA%E5%BA%9C%E5%A4%A7%E9%98%AA%E5%B8%82&amp;ll=34.693738,135.502165&amp;spn=0.471406,0.617294&amp;z=11&amp;output=embed"></iframe>
`

func main() {
	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("webkit")
	window.Connect("destroy", gtk.MainQuit)

	vbox := gtk.NewVBox(false, 1)

	entry := gtk.NewEntry()
	entry.SetText("http://golang.org/")
	vbox.PackStart(entry, false, false, 0)

	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.SHADOW_IN)

	webview := webkit.NewWebView()
	webview.Connect("load-committed", func() {
		entry.SetText(webview.GetUri())
	})
	swin.Add(webview)

	vbox.Add(swin)

	entry.Connect("activate", func() {
		webview.LoadUri(entry.GetText())
	})
	button := gtk.NewButtonWithLabel("load String")
	button.Clicked(func() {
		webview.LoadString("hello Go GTK!", "text/plain", "utf-8", ".")
	})
	vbox.PackStart(button, false, false, 0)

	button = gtk.NewButtonWithLabel("load HTML String")
	button.Clicked(func() {
		webview.LoadHtmlString(HTML_STRING, ".")
	})
	vbox.PackStart(button, false, false, 0)

	button = gtk.NewButtonWithLabel("Google Maps")
	button.Clicked(func() {
		webview.LoadHtmlString(MAP_EMBED, ".")
	})
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
	entry.Emit("activate")
	gtk.Main()
}
