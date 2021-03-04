-include .env

# go variables
GOBASE := $(shell pwd)
ROOM_MAINFILE := $(wildcard ./cmd/roomTest/main.go)
LEVEL_MAINFILE := $(wildcard ./cmd/levelTest/main.go)

BINARYNAME := "testLevel"

roomTest:
	echo $(PWD)
	go build -o $(BINARYNAME) $(ROOM_MAINFILE)
	chmod +x $(BINARYNAME)

levelTest:
	echo $(PWD)
	go build -o $(BINARYNAME) $(LEVEL_MAINFILE)
	chmod +x $(BINARYNAME)

clean:
	rm $(BINARYNAME)
