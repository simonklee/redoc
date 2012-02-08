# redoc

redoc - an interface to the [Redis documentation](http://redis.io/commands).
You can access extended documentation for each Redis command or get a list of
commands and groups available. With redoc you don't need to run redis-cli or
access redis.io to lookup documentation.

## Get redoc 

This command will checkout and install redoc.

    $ go get github.com/simonz05/redoc

## Use redoc

Without args redoc will display all Redis commands.

Arguments.

    command     display a given command
    [@]group    display all commands in a group. 
                @ is so to resolve naming conflicts.
    -lc         list available Redis commands
    -lg         list available Redis groups

Options.

    -d          long description for each Redis command
    -s          version of Redis since the commands was 
                supported
    -c=false    don't display output in colors


Examples.

    $ redoc set

        SET key value
        summary: Set the string value of a key
        group: string

    $ redoc -d set 

        SET key value
        summary: Set the string value of a key
        group: string

        @complexity

        O(1)

        Set 'key' to hold the string 'value'. If 'key' already holds
        a value, it is
        overwritten, regardless of its type.

        @return

        @status-reply: always 'OK' since 'SET' can't fail.

        @examples

        @cli
        SET mykey "Hello"
        GET mykey

## Update redoc

We generate a new commands.go file by running using the
update `update/update.go`. Simply run.

    $ ./update.sh

This will build and install an updated version of redoc
using the latest version of redis-doc.
