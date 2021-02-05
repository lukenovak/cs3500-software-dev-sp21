package main

import (
	numJson "github.ccs.neu.edu/CS4500-S21/Ormegland/A2/src/numJson"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

const port = ":8000"

func main() {

	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if conn != nil {
		defer conn.Close()
	}

	tcpStream := json.NewDecoder(bufio.NewReader(conn))

	numJsons, err := numJson.ParseNumJsonFromStream(tcpStream)

	var output json.RawMessage
	output = numJson.GenerateOutput(numJsons, numJson.Sum)

	if output == nil {
		fmt.Println("Error: no input")
		os.Exit(1)
	}
	_, err = conn.Write(output)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(conn, "\n")

	os.Exit(0)
}
