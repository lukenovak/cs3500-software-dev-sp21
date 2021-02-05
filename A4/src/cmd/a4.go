package main

import (
	"bytes"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/A3/traveller-client/parse"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/A4/src/internal/travelerJson"
	"encoding/json"
	"fmt"
	"io"
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

	// user messages
	sessionMsg = "the server will call me"
	invalidMsg = "\"invalid placement\""
	responseMsg = "\"the response for\""

	// command parsing
	place = "place"
	passageSafe = "passage-safe?"
)

func main() {
	os.Exit(batchCommandLoop(handleFirstCommand(connectToServer(parseArgs(os.Args[1:])))))
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
	fullIp := fmt.Sprintf("%s:%d", ip, port)
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
	sessionIdString := string(bytes.Trim(sessionId, "\u0000\n"))
	_, err = os.Stdout.Write(generateSessionMessage(sessionIdString))
	return &conn, string(sessionId)

}

// the first command needs to be a create command, so it gets its own function
func handleFirstCommand(conn *net.Conn, sessionId string) *net.Conn {
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

	if err != nil {
		panic(err)
	}
	return conn
}

func generateSessionMessage(sessionId string) json.RawMessage {
	var stringArray []string
	stringArray = append(append(stringArray, sessionMsg), sessionId)
	msg, err := json.Marshal(stringArray)
	if err != nil {
		panic(err)
	}
	return msg
}

// the main loop for reading commands from stdin, batching them, and sending them to the server
func batchCommandLoop(conn *net.Conn) int  {
	defer(*conn).Close()

	decoder := json.NewDecoder(os.Stdin)
	var err error
	var charData []travelerJson.CharacterData
	for err != io.EOF {
		err = nil
		var command parse.Command
		err = decoder.Decode(&command)
		if err == nil {
			switch command.Command {
			case place:
				charData = append(charData, parsePlaceCommand(command.Params))
			case passageSafe:
				queryData := parsePassageSafe(command.Params)
				writeOutput(sendBatchRequest(conn, charData, queryData), queryData)
			default:
				panic(fmt.Errorf("invalid command type! Killing sesison"))
			}
		} else {
			println(err)
		}
	}

	return 0
}

func parsePlaceCommand(params json.RawMessage) travelerJson.CharacterData {
	charParam := parseCharParam(params)
	return travelerJson.CharacterData{
		Name: charParam.Character,
		Town: charParam.Town,
	}

}

func parsePassageSafe(params json.RawMessage) travelerJson.QueryData {
	charParam := parseCharParam(params)
	return travelerJson.QueryData{
		Character: charParam.Character,
		Destination: charParam.Town,
	}
}

func parseCharParam(params json.RawMessage) parse.CharacterParam {
	var charParam parse.CharacterParam
	err := json.Unmarshal(params, &charParam)
	if err != nil {
		panic(err)
	}
	return charParam
}

func sendBatchRequest(conn *net.Conn,
	charData []travelerJson.CharacterData,
	queryData travelerJson.QueryData) travelerJson.ResponseData {

	batchRequest := travelerJson.BatchRequest{Characters: charData, Query: queryData}
	writeMsg, err := json.Marshal(batchRequest)


	if err != nil {
		panic(err)
	}
	_, err = (*conn).Write(writeMsg)
	if err != nil {
		panic(err)
	}

	// read back and decode the response
	decoder := json.NewDecoder(*conn)
	var responseData travelerJson.ResponseData
	if err = decoder.Decode(&responseData); err != nil {
		panic(err)
	}
	return responseData
}

func writeOutput(responseData travelerJson.ResponseData, queryData travelerJson.QueryData) {
	for _, placement := range responseData.Invalid {
		placementString := fmt.Sprintf("[%s, {\"name\" : \"%s\", \"town\" : \"%s\"}]", invalidMsg,
			placement.Name, placement.Town)
		_, _ = os.Stdout.WriteString(placementString)
	}
	responseString := fmt.Sprintf("[%s, {\"character\" : \"%s\", \"destination\" : \"%s\"}, \"is\", %b",
		responseMsg, queryData.Character, queryData.Destination, responseData.Response)
	_, _ = os.Stdout.WriteString(responseString)
}
