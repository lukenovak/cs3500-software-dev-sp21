package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/remote"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/remote/server"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
	"net"
	"time"
)

const (
	defaultTimeout = 60
	defaultLevels = "snarl.levels"
	defaultClients = 1
	defaultObserve = false
	defaultAddress = "127.0.0.1"
	defaultPort = 45678
	nameMessage = "\"name\"\n"
)

// main runs the server
func main() {
	// parse command line arguments
	timeout, levelPath, numClients, _, address, port := parseArguments()
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		panic(err)
	}
	var players []state.UserClient
	var gamePlayers []actor.Actor
	var names []string
	for connectedClients := 0; connectedClients < numClients; {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		conn.Write(remote.NewServerWelcomeMessage())
		time.Sleep(500 * time.Millisecond)
		var name []byte
		byteChan := make(chan []byte)
		go func() {
			for {
				println("reading")
				b := make([]byte, 4096)
				conn.Write([]byte(nameMessage))
				n, _ := conn.Read(b)
				if n > 0 {
					println("got name")
					byteChan <- bytes.Trim(b[0:n], "\r\n")
					break
				}
			}
		}()
		var playerName string
		for {
			select {
			case name = <- byteChan:
				playerName = string(name)
				break
			default:
				continue
			}
			break
		}
		println("got player name")
		newPlayer := server.NewPlayerClient(playerName, conn, timeout)
		newGamePlayer, _ := newPlayer.RegisterClient()
		gamePlayers = append(gamePlayers, newGamePlayer)
		players = append(players, newPlayer.AsUserClient())
		names = append(names, playerName)
		connectedClients += 1
	}
	levels, _ := level.ParseLevelFile(levelPath, 1)
	for _, player := range players {
		// create start message
		startJson := remote.NewStartLevel(1, names)
		msg, _ := json.Marshal(startJson)
		player.(*server.PlayerClient).SendJsonMessage(msg)
	}
	state.GameManager(
		levels,
		players,
		gamePlayers,
		nil,
		nil,
		len(levels),
	)
}

// parseArguments initializes flags for the executable
func parseArguments() (time.Duration, string, int, bool, string, int) {
	timeout := flag.Int("wait", defaultTimeout, "used to determine the amount of time to wait for players to register from booting the server")
	levelPath := flag.String("levels", defaultLevels, "tells the server which levels file to use. Default is ./snarl.levels")
	clients := flag.Int("clients", defaultClients, "tells the server how many clients to wait for. Default is 4")
	shouldObserve := flag.Bool("observe", defaultObserve, "launches a local observer if toggled")
	address := flag.String("address", defaultAddress, "tells the server what ip address to listen on")
	port := flag.Int("port", defaultPort, "tells the server what por to listen on")
	flag.Parse()
	// dereferences should be safe because we have default values
	timeoutInt := *timeout
	timeoutSecond := time.Duration(timeoutInt) * time.Second
	return timeoutSecond, *levelPath, *clients, *shouldObserve, *address, *port
}