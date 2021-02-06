package main

import (
	"net"
	"testing"
)

const (
	idMessage = "mockId"
)

// tests that the startup process works fine
func TestStartup(t *testing.T) {

	// start a mock server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		t.Fatal(err)
	}

	// this does not return anything, and we can't capture stdout. dunno how we verify the correct message is being
	// written to the terminal
	go connectToServer("", defaultPort, defaultName)

	conn, _ := listener.Accept()
	defer conn.Close()
	nameBuf := make([]byte, 4096)
	l, err := conn.Read(nameBuf)
	nameBuf = nameBuf[0:l]
	if err != nil {
		t.Fatal(err)
	}
	if string(nameBuf) != defaultName {
		t.Fail()
	}
	_, err = conn.Write([]byte(idMessage))
	if err != nil {
		t.Fatal(err)
	}
	return

}