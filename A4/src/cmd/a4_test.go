package main

import (
	"bytes"
	"net"
	"testing"
)

const (
	idMessage = "mockId"
)

// tests that the startup process works fine
func TestStartup(t *testing.T) {

	go func() {
		listener, err := net.Listen("tcp", ":8080")
		if err != nil {
			t.Fatal(err)
		}
		conn, _ := listener.Accept()
		defer conn.Close()
		var name bytes.Buffer
		_, err = conn.Read(name.Bytes())
		if err != nil {
			t.Fatal(err)
		}
		_, err = conn.Write([]byte(idMessage))
		if err != nil {
			t.Fatal(err)
		}

	}()

	connectToServer("", defaultPort, defaultName)

}