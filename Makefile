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

format:
	gofmt -s=true -tabs=false -tabwidth=4 -w .
