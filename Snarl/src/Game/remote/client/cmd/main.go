package main

import (
	"bufio"
	"encoding/asn1"
	"encoding/json"
	"flag"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/remote"
	"net"
	"os"
)

const (
	defaultAddress = "127.0.0.1"
	defaultPort    = 45678
)

func main() {
	// setup the gui application
	a := app.New()
	fyne.SetCurrentApp(a)

	// parse command line arguments
	address, port := parseArguments()
	socket := fmt.Sprintf("%v:%v", address, port)

	// get client name from stdin
	fmt.Print("Enter your client's name: ")
	input := bufio.NewReader(os.Stdin)
	name, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read name.")
		panic(err)
	}

	// connect to server
	fmt.Printf("Connecting to %v\n", socket)
	conn, err := net.Dial("tcp", socket)
	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)
	if err != nil {
		fmt.Println("Failed to connect to the server.")
		panic(err)
	}

	// name handshake
	var nameCommand string
	err = decoder.Decode(&nameCommand)
	if err != nil {
		fmt.Println("error decoding json")
		panic(err)
	}
	if nameCommand != "name" {
		panic(fmt.Errorf("did not recive name request as expected"))
	}
	err = encoder.Encode(name)
	if err != nil {
		fmt.Println("Failed to send name.")
		panic(err)
	}

	// game loop
	for {
		rawData := remote.BlockingRead(conn)
		var parsedData typedJson
		json.Unmarshal(*rawData, &parsedData)
		switch parsedData.Type {
		case "end-game":
			var endGame remote.EndGame
			asn1.Unmarshal(*rawData, &endGame)
			fmt.Println(endGame.Scores)
			return
		case "start-level":
			runLevel(conn)
		}
	}
}

type typedJson struct {
	Type string `json:"type"`
}

func runLevel(conn net.Conn) {
	for {
		rawData := remote.BlockingRead(conn)
		var parsedData interface{}
		json.Unmarshal(*rawData, parsedData)
		switch typedData := parsedData.(type) {
		case string:
			if typedData == "move" {
				handleMove()
			}
		case map[string]interface{}:
			var parsedData typedJson
			json.Unmarshal(*rawData, &parsedData)
			switch parsedData.Type {
			case "end-level":
				return
			case "player-update":
				updateGui()
			}
		}
	}
}

func handleMove() {

}

func parseArguments() (string, int) {
	address := flag.String("address", defaultAddress, "tells the client what IP address to connect over")
	port := flag.Int("port", defaultPort, "tells the client what port to connect over")
	flag.Parse()
	return *address, *port
}
