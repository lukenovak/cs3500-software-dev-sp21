package travelerJson

import (
	"encoding/json"
	"github.ccs.neu.edu/CS4500-S21/Ormegland/A3/traveller-client/parse"
)

func ParsePlaceCommand(params json.RawMessage) CharacterData {
	charParam := parseCharParam(params)
	return CharacterData {
		Name: charParam.Character,
		Town: charParam.Town,
	}

}

func ParsePassageSafe(params json.RawMessage) QueryData {
	charParam := parseCharParam(params)
	return QueryData {
		Character: charParam.Character,
		Destination: charParam.Town,
	}
}

func parseCharParam(params json.RawMessage) parse.CharacterParam {
	var charParam parse.CharacterParam
	err := json.Unmarshal(params, &charParam)
	if err != nil {
		panic(err)
	}
	return charParam
}