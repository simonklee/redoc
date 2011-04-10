include $(GOROOT)/src/Make.inc

TARG=redoc
GOFILES=\
	redoc.go\
	commands.go\

include $(GOROOT)/src/Make.cmd

update: update.go 
	$(GC) $<
	$(LD) -o $@ $@.$O
	./update
