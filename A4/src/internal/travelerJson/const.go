package travelerJson

// Top level JSON structures
type CreateRequest struct {
	Towns []string `json:"towns"`
	Roads []FromTo `json:"roads"`
}

type BatchRequest struct {
	Characters []CharacterData 	`json:"characters"`
	Query QueryData				`json:"query"`
}

type ResponseData struct {
	Invalid []CharacterData	`json:"invalid"`
	Response bool			`json:"response"`
}

// Data Object types DOTs
type FromTo struct {
	From string `json:"from"`
	To string	`json:"to"`
}

type CharacterData struct {
	Town string	`json:"town"`
	Name string	`json:"name"`
}

type QueryData struct {
	Character string 	`json:"character"`
	Destination string 	`json:"destination"`
}
