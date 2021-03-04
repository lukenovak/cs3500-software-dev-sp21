-include .env

# go variables
GOBASE := $(shell pwd)
LEVEL_MAINFILE := $(wildcard ./cmd/main.go)

BINARYNAME := "testLevel"

build:
	echo $(PWD)
	go build -o $(BINARYNAME) $(LEVEL_MAINFILE)
	chmod +x $(BINARYNAME)

clean:
	rm $(BINARYNAME)
