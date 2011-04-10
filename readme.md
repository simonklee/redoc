# redoc

Provides [Redis documentation](http://redis.io/commands) from a simple command
line interface. With `redoc` command you don't have to have a redis-server
running in order access documentation through `redis-cli help`. Furthermore, the
extended documentation as used on redis.io is also available for each command.

## Get it

    goinstall -v -u github.com/simonz05/redoc

Now you have a `redoc` command at your fingertips. 

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
