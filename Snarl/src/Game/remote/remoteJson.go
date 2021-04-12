package remote

type ServerWelcome struct {
	Type string `json:"type"`
	Info string `json:"info"`
}

type Name string

type StartLevel struct {
	Type    string `json:"type"`
	Level   int    `json:"level"`
	Players []Name `json:"players"`
}

type Point [2]int

type Object struct {
	Type     string `json:"type"`
	Position Point  `json:"position"`
}

type ActorPosition struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Position Point  `json:"position"`
}

type PlayerUpdateMessage struct {
	Type     string          `json:"type"`
	Layout   [][]int         `json:"layout"`
	Position Point           `json:"position"`
	Objects  []Object        `json:"objects"`
	Actors   []ActorPosition `json:"actors"`
	Message  string          `json:"message"`
}

type PlayerMove struct {
	Type string `json:"type"`
	To   *Point `json:"to"`
}

type Result string

type EndLevel struct {
	Type   string `json:"type"`
	Key    Name   `json:"key"`
	Exits  []Name `json:"exits"`
	Ejects []Name `json:"ejects"`
}

type PlayerScore struct {
	Type   string `json:"type"`
	Name   Name   `json:"name"`
	Exits  int    `json:"exits"`
	Ejects int    `json:"ejects"`
	Keys   int    `json:"keys"`
}

type EndGame struct {
	Type   string        `json:"type"`
	Scores []PlayerScore `json:"scores"`
}
