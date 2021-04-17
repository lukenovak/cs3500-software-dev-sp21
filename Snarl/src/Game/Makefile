-include .env

# go variables
GOBASE := $(shell pwd)
MAINFILE := $(wildcard ./local/cmd/main.go)

BINARYNAME := "snarl"

build:
	echo $(PWD)
	go build -o $(BINARYNAME) $(MAINFILE)
	chmod +x $(BINARYNAME)

clean:
	rm $(BINARYNAME)
