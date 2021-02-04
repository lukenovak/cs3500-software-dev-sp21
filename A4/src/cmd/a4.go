package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

const (
	// default arguments
	defaultIp = "127.0.0.1"
	defaultPort = 8080
	defaultName = "Glorifrir Flintshoulder"

	// utility
	tcp = "tcp"
)

func main() {
	conn, sessionId := connectToServer(parseArgs(os.Args[1:]))

	// TODO: handle commands
}

func parseArgs(args []string) (string, int, string) {
	switch len(args) {
	case 0:
		return defaultIp, defaultPort, defaultName
	case 1:
		return args[0], defaultPort, defaultName
	case 2:
		portNumber, err := strconv.Atoi(args[1])
		if err != nil {
			return "", 0, ""
		}
		return args[0], portNumber, defaultName
	case 3:
		portNumber, err := strconv.Atoi(args[1])
		if err != nil {
			return  "", 0, ""
		}
		return args[0], portNumber, args[2]
	default:
		panic(fmt.Errorf("incorrect number of arguments"))

	}
}

// connects to the server, sends the name over, and returns the connection and session ID
func connectToServer(ip string, port int, name string) (*net.Conn, string) {
	fullIp := fmt.Sprintf("%s%d", ip, port)
	conn, err := net.Dial(tcp, fullIp)
	if err != nil {
		panic(err)
	}

	// send the name to the server, expect
	_, err = conn.Write([]byte(name))
	if err != nil {
		panic(err)
	}

	// read the session id back
	sessionId := make([]byte, 64)
	_, err = conn.Read(sessionId)
	if err != nil {
		panic(err)
	}

	return &conn, string(sessionId)

}