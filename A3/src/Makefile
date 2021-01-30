-include .env

# go variables
GOBASE := $(shell pwd)
MAINFILE := $(wildcard ./cmd/a3.go)

BINARYNAME := "a3"


build:
	go get "github.ccs.neu.edu/CS4500-S21/Ormegland/A2/src/numJson"
	go build -o $(BINARYNAME) $(MAINFILE)
	chmod +x $(BINARYNAME)

clean:
	rm $(BINARYNAME)
