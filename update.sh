#!/bin/bash

git submodule update --init 
git submodule foreach 'git pull origin master'
cd update/
go build .
./update
mv commands.go ../
cd ..
./fmt.sh .
