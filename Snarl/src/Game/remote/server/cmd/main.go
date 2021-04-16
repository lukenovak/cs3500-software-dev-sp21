package main

import (
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
	nameMessage = "name"
)

func main() {
	// parse command line arguments
	timeout, levelPath, numClients, _, address, port := parseArguments()
	listener, _ := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
	var players []state.UserClient
	var gamePlayers []actor.Actor
	for connectedClients := 0; connectedClients < numClients; {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		conn.Write(remote.NewServerWelcomeMessage())
		conn.Write([]byte(nameMessage))
		var name []byte
		println("red")
		byteChan := make(chan []byte)
		go func() {
			for {
				b := make([]byte, 4096)
				n, _ := conn.Read(b)
				if n > 0 {
					byteChan <- b
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
		print("read")
		newPlayer := server.NewPlayerClient(playerName, conn, timeout)
		newGamePlayer, _ := newPlayer.RegisterClient()
		gamePlayers = append(gamePlayers, newGamePlayer)
		players = append(players, newPlayer.AsUserClient())
		connectedClients += 1
	}

	levels, _ := level.ParseLevelFile(levelPath, 1)
	state.GameManager(
		levels,
		players,
		gamePlayers,
		nil,
		nil,
		len(levels),
	)
}

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