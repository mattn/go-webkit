include $(GOROOT)/src/Make.inc

TARG     = github.com/mattn/go-webkit/webkit
CGOFILES = webkit.go

include $(GOROOT)/src/Make.pkg
