-include .env

# go variables
GOBASE := $(shell pwd)
OLDGOPATH := ${GOPATH}
MAINFILE := $(wildcard ./cmd/a2.go)

BINARYNAME := "a2"


build:
	go build -o $(BINARYNAME) $(MAINFILE)
	chmod +x $(BINARYNAME)

clean:
	rm $(BINARYNAME)

