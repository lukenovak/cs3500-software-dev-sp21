package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/internal/render"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/remote"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/remote/client"
	"golang.org/x/image/colornames"
	"net"
	"os"
)

const (
	defaultAddress = "127.0.0.1"
	defaultPort    = 45678
	nameCommand    = "name"
	moveCommand    = "move"
)

func main() {
	// setup the gui application and keyboard
	a := app.New()
	fyne.SetCurrentApp(a)
	gameWindow := a.NewWindow("snarl client")


	// parse command line arguments
	address, port := parseArguments()
	socket := fmt.Sprintf("%v:%v", address, port)

	// get client name from stdin
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your client's name: ")
	name, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read name.")
		panic(err)
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

	// name handshake
	var serverNameCommand string
	err = decoder.Decode(&serverNameCommand)
	println(serverNameCommand)
	if err != nil {
		fmt.Println("error decoding json")
		panic(err)
	}
	if serverNameCommand != nameCommand {
		panic(fmt.Errorf("did not recive name request as expected"))
	}
	_, err = conn.Write([]byte(fmt.Sprintf("%s\n", name)))
	if err != nil {
		fmt.Println("Failed to send name.")
		panic(err)
	}

	// set up connection for i/o
	connReader := bufio.NewReader(conn)

	// move to game Loop
	rawData := remote.BlockingRead(connReader)
	var parsedData typedJson
	json.Unmarshal(*rawData, &parsedData)
	switch parsedData.Type {
	case "start-level":
		go runGame(conn, connReader, name, gameWindow)
		gameWindow.Show()
		a.Run()
	default:
		println("cannot start game without start-level command")
	}
}

// typedJson is used to unmarshal an unknown json to determine its type
type typedJson struct {
	Type string `json:"type"`
}

// runGame runs the main game loop
func runGame(conn net.Conn, connReader *bufio.Reader, playerName string, gameWindow fyne.Window) {

	// create a new client-side player representation for this player
	player := client.Player{
		Name:       playerName,
		Posn:       level.Position2D{},
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
				player.HandleMove(conn, connReader)
			}
		case map[string]interface{}:
			var parsedData typedJson
			json.Unmarshal(*rawData, &parsedData)
			switch parsedData.Type {
			case "end-level":
				var endLevel remote.EndLevel
				json.Unmarshal(*rawData, &endLevel)
				fmt.Println(endLevel)
				return
			case "player-update":
				var updateMessage remote.PlayerUpdateMessage
				json.Unmarshal(*rawData, &updateMessage)
				updateGui(updateMessage, gameWindow)
				player.Posn = updateMessage.Position.ToPos2D()
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

func updateGui(updateMessage remote.PlayerUpdateMessage, gameWindow fyne.Window) {

	// converting to tiles
	tiles := make([][]*level.Tile, 0)
	for i, row := range updateMessage.Layout {
		tiles = append(tiles, make([]*level.Tile, len(row)))
		for j, tileType := range row {
			tile := level.Tile{
				Type: tileType,
			}
			tiles[i][j] = &tile
		}
	}

	// add items to tiles
	for _, object := range updateMessage.Objects {
		var item level.Item
		switch object.Type {
		case "key":
			item = level.Item{Type: level.KeyID}
		case "exit":
			item = level.Item{Type: level.LockedExit}
		}
		tiles[object.Position[0]][object.Position[1]].Item = &item
	}

	// generate actor lists
	players := make([]actor.Actor, 0)
	adversaries := make([]actor.Actor, 0)
	convertToRelative := func(pos level.Position2D) level.Position2D {
		updatePosition := updateMessage.Position.ToPos2D()
		return level.NewPosition2D(pos.Row - updatePosition.Row + 2, pos.Col - updatePosition.Col + 2)
	}
	for _, actorData := range updateMessage.Actors {
		switch actorData.Type {
		case "player":
			players = append(players, actor.Actor{
				Type:        actor.PlayerType,
				Name:        actorData.Name,
				Position:    convertToRelative(actorData.Position.ToPos2D()),
				RenderedObj: canvas.NewCircle(colornames.Cornflowerblue),
			})
		case "zombie":
			adversaries = append(adversaries, actor.Actor{
				Type:        actor.ZombieType,
				Name:        actorData.Name,
				Position:    convertToRelative(actorData.Position.ToPos2D()),
				RenderedObj: canvas.NewCircle(colornames.Crimson),
			})
		case "ghost":
			adversaries = append(adversaries, actor.Actor{
				Type:        actor.GhostType,
				Name:        actorData.Name,
				Position:    convertToRelative(actorData.Position.ToPos2D()),
				RenderedObj: canvas.NewCircle(colornames.Hotpink),
			})
		}
	}

	render.GuiState(tiles, players, adversaries, gameWindow)
}

func parseArguments() (string, int) {
	address := flag.String("address", defaultAddress, "tells the client what IP address to connect over")
	port := flag.Int("port", defaultPort, "tells the client what port to connect over")
	flag.Parse()
	return *address, *port
}
