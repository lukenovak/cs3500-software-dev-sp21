package parse

import (
	"encoding/json"
	"fmt"
	"io"
)

type Command struct {
	Command string
	Params  json.RawMessage
}

type RoadArray []RoadParam

type RoadParam struct {
	From string
	To   string
}

type CharacterParam struct {
	Character string
	Town      string
}

func ParseCommands(reader io.Reader) ([]Command, error) {
	decoder := json.NewDecoder(reader)
	var commandList []Command

	for {
		var command Command
		if err := decoder.Decode(&command); err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, fmt.Errorf("JSON decode failed. Check that your input is valid")
			}
		}
		commandList = append(commandList, command)
	}

	return commandList, nil
}
