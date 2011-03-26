# redoc

Provides [Redis documentation](http://redis.io/documentation) as a command line tool. 

## Requires 

* [Go](http://golang.org/doc/install.html) to build the `redoc` binary.

## Get it

1. `git clone git://github.com/simonz05/redoc.git`
2. `cd redoc/`
3. `make`

Now you have a `./redoc` command at your fingertips. This can simply be copied
to a folder in your $PATH.

## Use it

* `redoc` without args will display all Redis commands, ordered by `group`.
* `redoc [group|command]` will either display a given command or all commands
  in a group.
* `redoc -d` displays a long description for Redis commands.
* `redoc -lc` will output all available commands in a comprehensive list.
* `redoc -lg` will output all available groups in a comprehensive list.

## Update

1. `git submodule update --init`
2. `6g update.go && 6l -o update update.6`

Now you have a `$ ./update` command which will regenerate the `api.go`-file.

3. `./update && make`

After running make `./redoc` is potentially updated with new documentation.
