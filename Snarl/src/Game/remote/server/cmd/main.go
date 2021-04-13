package main

import (
	"flag"
)

const (
	defaultTimeout = 60
	defaultLevels = "snarl.levels"
	defaultClients = 4
	defaultObserve = false
	defaultAddress = "127.0.0.1"
	defaultPort = 45678
)

func main() {
	// parse command line arguments
	_, _, _, _, _, _ = parseArguments()
}

func parseArguments() (int, string, int, bool, string, int) {
	timeout := flag.Int("wait", defaultTimeout, "used to determine the amount of time to wait for players to register from booting the server")
	levelPath := flag.String("levels", defaultLevels, "tells the server which levels file to use. Default is ./snarl.levels")
	clients := flag.Int("clients", defaultClients, "tells the server how many clients to wait for. Default is 4")
	shouldObserve := flag.Bool("observe", defaultObserve, "launches a local observer if toggled")
	address := flag.String("address", defaultAddress, "tells the server what ip address to listen on")
	port := flag.Int("port", defaultPort, "tells the server what por to listen on")
	flag.Parse()
	// dereferences should be safe because we have default values
	return *timeout, *levelPath, *clients, *shouldObserve, *address, *port
}