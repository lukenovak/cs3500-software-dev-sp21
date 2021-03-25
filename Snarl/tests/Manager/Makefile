-include .env

# go variables
GOBASE := $(shell pwd)
MAINFILE := $(wildcard ./cmd/main.go)
TESTDIR := $(wildcard ./tests/)

BINARYNAME := "testManager"

build:
	echo $(PWD)
	go build -o $(BINARYNAME) $(MAINFILE)
	chmod +x $(BINARYNAME)

test:
	go build -o $(BINARYNAME) $(MAINFILE)
	chmod +x $(BINARYNAME)
	for test in $(TESTDIR);\
	do ./$(BINARYNAME) < $$test;\
	done

clean:
	rm $(BINARYNAME)
