#!/bin/bash

mkdir ./src/$1
echo -e "package main\n\nimport (\n\"fmt\"\n)\n\nfunc main() {\n}" >> ./src/$1/main.go