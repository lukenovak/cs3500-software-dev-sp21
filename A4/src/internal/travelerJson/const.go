package travelerJson

import (
	travellerParse "github.ccs.neu.edu/CS4500-S21/Ormegland/A3/traveller-client/parse"
)

// Top level JSON structures
type CreateRequest struct {
	Towns []string 						`json:"towns"`
	Roads travellerParse.RoadArray  	`json:"roads"`
}

type BatchRequest struct {
	Characters []CharacterData 	`json:"characters"`
	Query QueryData				`json:"query"`
}

type ResponseData struct {
	Invalid []CharacterData	`json:"invalid"`
	Response bool			`json:"response"`
}

// Json data objects
type CharacterData struct {
	Town string	`json:"town"`
	Name string	`json:"name"`
}

type QueryData struct {
	Character string 	`json:"character"`
	Destination string 	`json:"destination"`
}
