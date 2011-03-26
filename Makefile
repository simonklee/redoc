include $(GOROOT)/src/Make.inc

TARG=redoc
GOFILES=\
	redoc.go\
	api.go\

include $(GOROOT)/src/Make.cmd
