package main

import (
	"encoding/json"
	"fmt"
	"os"
	"../parse"
)

func main() {
	commandList, err := parse.ParseCommands(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	var n Network
	outputs, err := executeCommands(commandList, n)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	for _, output := range outputs {
		_, err := os.Stdout.WriteString(fmt.Sprintf("%s\n", output))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	os.Exit(0)
}

// runs through all commands and compiles the responses into an array of strings
func executeCommands(commandList []parse.Command, n Network) ([]string, error) {
	var responseList []string
	roadsNotFirstError := fmt.Errorf("first command must be a road")
	for idx, command := range commandList {
		switch command.Command {
		case "roads":
			var roads parse.RoadArray
			err := json.Unmarshal(command.Params, &roads)
			if err != nil {
				return responseList, err
			}
			for _, roadCom := range roads {
				startTown := getTown(roadCom.From, n)
				if startTown == nil {
					startTown = traveller.CreateTown(roadCom.From, n)
				}
				endTown := getTown(roadCom.To, n)
				if endTown == nil {
					endTown = traveller.CreateTown(roadCom.To, n)
				}
				traveller.LinkTowns(startTown, endTown)
			}
		case "place":
			if idx == 0 {
				return nil, roadsNotFirstError
			}
			charParams, err := parseCharacterParam(command.Params)
			if err == nil {
				c := Character(charParams.Character)
				town := getTown(charParams.Town, n)
				if town != nil {
					traveller.PlaceExistingCharacter(town, c)
				}
			}
		case "passage-safe?":
			if idx == 0 {
				return nil, roadsNotFirstError
			}
			charParams, err := parseCharacterParam(command.Params)
			if err == nil {
				c := getCharacter(charParams.Character, n)
				endTown := getTown(charParams.Town, n)
				canTravel := traveler.CanTravel(c, endTown, n)
				responseList = append(responseList, string(canTravel))
			}
		}
	}
	return responseList, nil
}

func parseCharacterParam(params json.RawMessage) (parse.CharacterParam, error) {
	var charParam parse.CharacterParam
	err := json.Unmarshal(params, &charParams)
	if err != nil {
		return parse.CharacterParam{}, err
	}
	return charParam, nil
}

func getTown(name string, n Network) Town {
	for _, town := range n {
		if town.Name == name {
			return town
		}
	}
	return nil
}

func getCharacter(charName string, n Network) Character {
	for _, town := range n {
		if town.Character.Name == charName {
			return town.Character
		}
	}
	return nil
}

type Network []interface{}

type Town interface {}

type Character interface {}
