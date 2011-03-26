# redoc

Provides [Redis documentation](http://redis.io/commands) from a simple command
line interface. With `redoc` command you don't have to have a redis-server
running in order access documentation through `redis-cli help`. Furthermore, the
extended documentation as used on redis.io is also available for each command.

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
* `redoc -lc` output all available Redis commands in a comprehensive list.
* `redoc -lg` output all available Redis groups in a comprehensive list.

## Update

First make sure the [redis-doc](https://github.com/antirez/redis-doc/) is present in
the folder:

`git submodule update --init && git submodule foreach pull`

Next we need to compile the update app. s/6/8/ if you are on 32-bit.

`6g update.go && 6l -o update update.6`

Now you have a `update` app which will regenerate the commands.go-file. You can
simply call it without args.

`./update`

Compile redoc with `make` and you are up to date.
