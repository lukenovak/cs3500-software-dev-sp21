-include .env

# go variables
GOBASE := $(shell pwd)
MAINFILE := $(wildcard ./cmd/main.go)

BINARYNAME := "snarl-adversary"

build:
	echo $(PWD)
	go build -o $(BINARYNAME) $(MAINFILE)
	chmod +x $(BINARYNAME)

clean:
	rm $(BINARYNAME)
