# redoc

redoc - an interface to the [Redis documentation](http://redis.io/commands).
You can access extended documentation for each Redis command or get a list of
commands and groups available. With redoc you don't need to run redis-cli or
access redis.io to lookup documentation.

## Get it

    $ git clone git://github.com/simonz05/redoc.git && cd redoc && make install

This will install `redoc` in your $GOROOT directory.

## Usage

`redoc`:

* Without args redoc redoc will display all Redis commands.
* `[command|@group]` display either a given command or all commands in a group.
  `$ redoc set` would display the command SET.
* `-d` long description for Redis commands.
* `-s` since which version of Redis command is supported.
* `-lc` list available Redis commands.
* `-lg` list available Redis groups.

## Update

First make sure the [redis-doc](https://github.com/antirez/redis-doc/) is present in
the folder:

`git submodule update --init && git submodule foreach pull`

Next we can generate a new commands.go file by running

`make update`

Compile redoc with `make install` and you are up to date.
