package main

import (
	"bufio"
	"flag"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/actor"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/internal/render"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/level"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/remote"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/Snarl/src/Game/state"
	"os"
)

const (
	// argument flag names
	levelFlagName        = "levels"
	playerFlagName       = "players"
	startLevelFlagName   = "start"
	showObserverFlagName = "observe"

	// argument defaults
	defaultNumPlayers   = 1
	defaultFilename     = "snarl.levels"
	defaultStartLevel   = 1
	defaultShowObserver = false
)

// main executable function
func main() {
	// initialize flags
	levelFlag := flag.String(levelFlagName, defaultFilename, "Points the game to the desired level file")
	playerFlag := flag.Int(playerFlagName, defaultNumPlayers, "The number of players in this game")
	startLevelFlag := flag.Int(startLevelFlagName, defaultStartLevel, "The level number to start at")
	showObserverFlag := flag.Bool(showObserverFlagName, defaultShowObserver, "Opens an observer window if present")

	flag.Parse()

	// error checking
	if *playerFlag < 1 || *playerFlag > 4 {
		fmt.Println("invalid number of players")
		os.Exit(1)
	}
	levels, err := level.ParseLevelFile(*levelFlag, *startLevelFlag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if *startLevelFlag > len(levels) || *startLevelFlag < 1 {
		fmt.Println("invalid start level")
		os.Exit(1)
	}

	// setup the gui application
	a := app.New()
	fyne.SetCurrentApp(a)

	// generate player clients based on user input
	var players []state.UserClient
	for i := 1; i <= *playerFlag; i++ {
		newPlayer := getLocalPlayer(i)
		players = append(players, newPlayer)
	}

	// set up the local observer if requested
	var observers []state.GameObserver
	if *showObserverFlag {
		observerWindow := a.NewWindow("snarl observer")
		observer := state.NewGameObserver(func(gs state.GameState) {
			render.GuiState(gs.Level.Tiles, gs.Players, gs.Adversaries, observerWindow)
		})
		observers = append(observers, observer)
	}

	// register players and prepare game player representations
	var gamePlayers []actor.Actor
	for _, player := range players {
		newPlayer, _ := player.RegisterClient()
		gamePlayers = append(gamePlayers, newPlayer)
	}

	var playerScores []remote.PlayerScore
	// local func run game and get return value
	runGame := func() {
		playerScores = state.ManageGame(levels, players, gamePlayers, observers, 1)
		fmt.Println("Player Leaderboard:")
		fmt.Println("Name, Exits, Keys, Ejections")
		for _, score := range playerScores {
			fmt.Println(score)
		}
		a.Quit()
	}

	// launch the main game loop
	go runGame()

	// display the window (this is blocking!)
	a.Run()

	os.Exit(0)
}

// DEPRECATED: generates an example level
func generateGameStateLevel() level.Level {
	newLevel, err := level.NewEmptyLevel(32, 32)
	if err != nil {
		panic(err)
	}
	// first room from 0,0 to 5, 4
	err = newLevel.GenerateRectangularRoom(level.NewPosition2D(0, 0),
		5,
		6,
		[]level.Position2D{level.NewPosition2D(4, 1), level.NewPosition2D(3, 5)})
	if err != nil {
		panic(err)
	}

	// second room from 9, 9 to 14, 16
	err = newLevel.GenerateRectangularRoom(level.NewPosition2D(9, 9),
		8,
		6,
		[]level.Position2D{level.NewPosition2D(13, 9)})
	if err != nil {
		panic(err)
	}

	// Third room from 20, 21 to 28, 29
	err = newLevel.GenerateRectangularRoom(level.NewPosition2D(21, 20),
		9,
		9,
		[]level.Position2D{level.NewPosition2D(25, 20)})
	if err != nil {
		panic(err)
	}

	// connecting hallways
	hallwayWaypoints := []level.Position2D{{3, 7}, {13, 7}}
	err = newLevel.GenerateHallway(level.NewPosition2D(3, 5), level.NewPosition2D(13, 9), hallwayWaypoints)
	if err != nil {
		panic(err)
	}

	hallwayWaypoints = []level.Position2D{{25, 1}}
	err = newLevel.GenerateHallway(level.NewPosition2D(4, 1), level.NewPosition2D(25, 20), hallwayWaypoints)
	if err != nil {
		panic(err)
	}

	exit := level.Item{Type: level.LockedExit}
	newLevel.PlaceItem(level.NewPosition2D(14, 12), &exit)

	key := level.Item{Type: level.KeyID}
	newLevel.PlaceItem(level.NewPosition2D(25, 25), &key)

	return newLevel
}

func getLocalPlayer(playerNumber int) *state.LocalKeyClient {
	fmt.Printf("Enter player %d's name: ", playerNumber)
	name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return &state.LocalKeyClient{
		Name:       name,
		GameWindow: nil,
	}

}
