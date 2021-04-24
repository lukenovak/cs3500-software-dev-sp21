package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/remote"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/remote/adversary"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
	"net"
	"os"
)

const (
	defaultAddress = "127.0.0.1"
	defaultPort    = 45678
	typeCommand    = "type"
	moveCommand    = "move"
)

func main() {
	// setup the gui application
	a := app.New()
	fyne.SetCurrentApp(a)
	gameWindow := a.NewWindow("snarl client")

	// parse command line arguments
	address, port := parseArguments()
	socket := fmt.Sprintf("%v:%v", address, port)

	// get adversary type from stdin
	input := bufio.NewReader(os.Stdin)
	var adversaryType string
	for {
		fmt.Print("Enter your adversary type (one of 'zombie' or 'ghost'): ")
		adversaryType, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read type.")
			panic(err)
		}
		if adversaryType == "zombie" || adversaryType == "ghost" {
			break
		} else {
			fmt.Println("Invalid type. Try again.")
		}
	}

	// connect to server
	fmt.Printf("Connecting to %v\n", socket)
	conn, err := net.Dial("tcp", socket)
	decoder := json.NewDecoder(conn)
	//encoder := json.NewEncoder(conn)
	if err != nil {
		fmt.Println("Failed to connect to the server.")
		panic(err)
	}

	// handle welcome
	var serverWelcome remote.ServerWelcome
	err = decoder.Decode(&serverWelcome)
	if err == nil {
		println(serverWelcome.Info)
	}

	// type handshake
	var serverTypeCommand string
	err = decoder.Decode(&serverTypeCommand)
	if err != nil {
		fmt.Println("error decoding json")
		panic(err)
	}
	if serverTypeCommand != typeCommand {
		panic(fmt.Errorf("did not recive type request as expected"))
	}
	_, err = conn.Write([]byte(fmt.Sprintf("%s\n", adversaryType)))
	if err != nil {
		fmt.Println("Failed to send type.")
		panic(err)
	}

	// set up connection for i/o
	connReader := bufio.NewReader(conn)

	// move to game Loop
	rawData := remote.BlockingRead(connReader)
	var parsedData remote.TypedJson
	json.Unmarshal(*rawData, &parsedData)
	switch parsedData.Type {
	case "start-level":
		go runGame(conn, connReader, adversaryType, gameWindow)
		gameWindow.Show()
		a.Run()
	default:
		println("cannot start game without start-level command")
	}
}

// runGame runs the main game loop
func runGame(conn net.Conn, connReader *bufio.Reader, adversaryType string, gameWindow fyne.Window) {

	// create a new client-side adversary representation for this adversary
	var client state.AdversaryClient
	switch adversaryType {
	case "zombie":
		client = &state.ZombieClient{}
	case "ghost":
		client = &state.GhostClient{}
	default:
		panic("invalid adversary type")
	}
	localAdversary := adversary.Adversary{
		Client:     client,
		GameWindow: gameWindow,
	}

	// run the main loop
	for {
		// see what the server sent us, and act depending on what kind of message it is
		rawData := remote.BlockingRead(connReader)
		var parsedData interface{}
		json.Unmarshal(*rawData, &parsedData)
		switch typedData := parsedData.(type) {
		case string:
			if typedData == moveCommand {
				localAdversary.HandleMove(conn, connReader)
			}
		case map[string]interface{}:
			var parsedData remote.TypedJson
			json.Unmarshal(*rawData, &parsedData)
			switch parsedData.Type {
			case "end-level":
				var endLevel remote.EndLevel
				json.Unmarshal(*rawData, &endLevel)
				fmt.Println(endLevel)
				return
			case "adversary-update":
				var updateMessage remote.AdversaryUpdateMessage
				json.Unmarshal(*rawData, &updateMessage)
				// TODO make GUI update
				//updateGui(updateMessage, gameWindow)
				localAdversary.Client.UpdatePosition(updateMessage.Position.ToPos2D())
			case "start-level":
				// in this case, we just want to go to the next message which should be a player update
				println("advancing to the next level!")
				continue
			case "end-game":
				var endGame remote.EndGame
				json.Unmarshal(*rawData, &endGame)
				fmt.Println(endGame)
				fyne.CurrentApp().Quit()
				conn.Close()
				return
			}
		}
	}
}

func parseArguments() (string, int) {
	address := flag.String("address", defaultAddress, "tells the client what IP address to connect over")
	port := flag.Int("port", defaultPort, "tells the client what port to connect over")
	flag.Parse()
	return *address, *port
}
