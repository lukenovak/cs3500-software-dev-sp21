-include .env

# go variables
GOBASE := $(shell pwd)
GOFILES := $(wildcard ./src/*.go)

BINARYNAME := "a2"


build:
	go build -o $(BINARYNAME) $(GOFILES)
	chmod +x $(BINARYNAME)

clean:
	rm $(BINARYNAME)

