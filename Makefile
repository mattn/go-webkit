include $(GOROOT)/src/Make.inc

TARG     = webkit
CGOFILES = webkit.go

CGO_CFLAGS  = `pkg-config --cflags webkit-1.0`
CGO_LDFLAGS = -lpthread `pkg-config --libs webkit-1.0`
CGO_DEPS=

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
