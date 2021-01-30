package numJson

import "encoding/json"

// public function that generates the output message from an array of NumJson
func GenerateOutput(numJsons []NumJson, mode int) json.RawMessage {
	var outputJsons []OutputJson
	for _, nj := range numJsons {
		outputJsons = append(outputJsons, OutputJson{
			Object: nj,
			Total:  nj.NumValue(mode),
		})
	}
	var rawOut json.RawMessage
	rawOut, err := json.Marshal(&outputJsons)
	if err != nil {
		panic(err)
	}
	return rawOut
}
