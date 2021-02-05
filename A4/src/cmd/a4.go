package main

import (
	"../../../A3/traveller-client/parse"
	"../internal/travelerJson"
	"encoding/json"
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
	handleFirstCommand(connectToServer(parseArgs(os.Args[1:])))
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

// the first command needs to be a create command, so it gets its own function
func handleFirstCommand(conn *net.Conn, sessionId string) {
	decoder := json.NewDecoder(os.Stdin)
	var roadsCommand parse.Command
	var roadsArray parse.RoadArray
	err := decoder.Decode(&roadsCommand)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(roadsCommand.Params, &roadsArray)
	if err != nil {
		panic(err)
	}
	createRequest := travelerJson.CreateRequest {
		Roads: roadsArray,
		Towns: travelerJson.GetUniqueTowns(roadsArray),
	}
	message, err := json.Marshal(createRequest)
	if err != nil {
		panic(err)
	}
	_, err = (*conn).Write(message)
	if err != nil {
		panic(err)
	}
}
