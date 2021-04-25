package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"fyne.io/fyne/v2/app"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/internal/render"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/remote"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/remote/server"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
	"log"
	"net"
	"time"
)

const (
	defaultTimeout     = 60
	defaultLevels      = "snarl.levels"
	defaultClients     = 1
	defaultObserve     = false
	defaultAddress     = "127.0.0.1"
	defaultPort        = 45678
	defaultAdversaries = 0
	nameMessage        = "\"name\"\n"
)

// main runs the server
func main() {
	// parse command line arguments
	timeout, levelPath, numClients, numAdversaries, shouldObserve, address, port := parseArguments()
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		panic(err)
	}
	var players []state.UserClient
	var gamePlayers []actor.Actor
	var names []string

	// first phase- register players
	log.Println("registering players...")
	for connectedClients := 0; connectedClients < numClients; {
		// initial name handshake
		conn, playerName := initialHandshake(listener)

		// start client with name
		newPlayer := server.NewPlayerClient(playerName, conn, timeout)
		newGamePlayer, _ := newPlayer.RegisterClient()
		gamePlayers = append(gamePlayers, newGamePlayer)
		players = append(players, newPlayer.AsUserClient())
		names = append(names, playerName)
		connectedClients += 1
	}

	// second phase- register adversaries
	var adversaries []state.AdversaryClient
	log.Println("registering adversaries...")
	for connectedAdversasries := 0; connectedAdversasries < numAdversaries; {
		// get name
		conn, name := initialHandshake(listener)

		// get type
		_, err = conn.Write([]byte("\"type (1 for zombie, 2 for ghost)\"\n"))
		if err != nil {
			// if we can't write to the connection for the type, close it and move onto the next
			conn.Close()
			continue
		}
		connReader := bufio.NewReader(conn)
		var adversaryType int
		json.Unmarshal(*remote.BlockingRead(connReader), &adversaryType)

		// create a client with that name and type
		adversaries = append(adversaries, server.NewServerAdversary(name, adversaryType, conn))
	}


	levels, _ := level.ParseLevelFile(levelPath, 1)
	for _, player := range players {
		// create start message
		startJson := remote.NewStartLevel(1, names)
		msg, _ := json.Marshal(startJson)
		player.(*server.PlayerClient).SendJsonMessage(msg)
	}

	// handle observers
	observerApp := app.New()
	observerWindow := observerApp.NewWindow("snarl server observer")
	observer := state.NewGameObserver(func(gs state.GameState) {
		render.GuiState(gs.Level.Tiles, gs.Players, gs.Adversaries, observerWindow)
	})
	var observers []state.GameObserver = nil
	if shouldObserve {
		observers = append(observers, observer)
	}

	// start the game manager (blocking)
	scores := state.ManageGame(
		levels,
		players,
		gamePlayers,
		observers,
		adversaries,
		len(levels),
	)

	// if we get here, the game has ended
	endGame(players, scores)
	err = listener.Close()
}

// parseArguments initializes flags for the executable
func parseArguments() (time.Duration, string, int, int, bool, string, int) {
	timeout := flag.Int("wait", defaultTimeout, "used to determine the amount of time to wait for players to register from booting the server")
	levelPath := flag.String("levels", defaultLevels, "tells the server which levels file to use. Default is ./snarl.levels")
	clients := flag.Int("clients", defaultClients, "tells the server how many clients to wait for. Default is 4")
	numAdversaries := flag.Int("adversaries", defaultAdversaries, "tells the server how many adversaries to wait for. Default is 0 (no remote adversaries)")
	shouldObserve := flag.Bool("observe", defaultObserve, "launches a local observer if toggled")
	address := flag.String("address", defaultAddress, "tells the server what ip address to listen on")
	port := flag.Int("port", defaultPort, "tells the server what por to listen on")
	flag.Parse()
	// dereferences should be safe because we have default values
	timeoutInt := *timeout
	timeoutSecond := time.Duration(timeoutInt) * time.Second
	return timeoutSecond, *levelPath, *clients, *numAdversaries, *shouldObserve, *address, *port
}

// initialHandshake gets a connection remotely and gets the name of the incoming connection
func initialHandshake(listener net.Listener) (net.Conn, string) {
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	conn.Write(remote.NewServerWelcomeMessage())
	time.Sleep(500 * time.Millisecond)
	var name []byte
	byteChan := make(chan []byte)

	// special blocking read for initial handshake
	go func() {
		for {
			b := make([]byte, 4096)
			conn.Write([]byte(nameMessage))
			n, _ := conn.Read(b)
			if n > 0 {
				byteChan <- bytes.Trim(b[0:n], "\r\n")
				break
			}
		}
	}()
	var playerName string
	for {
		select {
		case name = <-byteChan:
			playerName = string(name)
			break
		default:
			continue
		}
		break
	}

	return conn, playerName
}

func endGame(players []state.UserClient, scores []remote.PlayerScore) {
	endGameMessage := remote.EndGame{
		Type:   "end-game",
		Scores: scores,
	}

	jsonMessage, _ := json.Marshal(endGameMessage)
	for _, player := range players {
		playerClient := player.(*server.PlayerClient)
		playerClient.SendJsonMessage(jsonMessage)
		playerClient.CloseConnection()
	}
}
