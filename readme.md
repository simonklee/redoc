# redoc

redoc - an interface to the [Redis documentation](http://redis.io/commands).
You can access extended documentation for each Redis command or get a list of
commands and groups available. With redoc you don't need to run redis-cli or
access redis.io to lookup documentation.

## Get redoc 

This command will checkout and install redoc.

    $ git clone git://github.com/simonz05/redoc.git && cd redoc && make install

## Use redoc

Without args redoc will display all Redis commands.

Arguments:

* `command` display a given command. `redoc set` would display the command SET.
* `[@]group` display all commands in a group. The optional @ is so to resolve name
  conflicts. `redoc @set` would display all set commands, and not the SET command.
* `-lc` list available Redis commands.
* `-lg` list available Redis groups.

Options will modify how the output is displayed: 

* `-d` long description for each Redis command.
* `-s` version of Redis since the commands was supported.
* `-c=false` don't display output in colors.

## Update redoc

First make sure the [redis-doc](https://github.com/antirez/redis-doc/) is present in
the folder:

`git submodule update --init && git submodule foreach git pull`

Next we can generate a new commands.go file by running

`make update`

Compile redoc with `make install` and you are up to date.
