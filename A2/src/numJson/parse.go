package numJson

import (
	"encoding/json"
	"fmt"
	"io"
)

// Reads from the stream and decodes each entry into a NumJson
func ParseNumJsonFromStream(d *json.Decoder) ([]NumJson, error) {
	var njArray []NumJson
	var err error
	for {
		var r json.RawMessage
		if err = d.Decode(&r); err != nil {
			if err == io.EOF {
				break
			} else {
				return njArray, fmt.Errorf("JSON decode failed. Check that your input is valid")
			}
		}
		marshalledJson, err := marshalNumJson(r)
		if err != nil {
			return nil, err
		}
		njArray = append(njArray, marshalledJson)
	}
	return njArray, nil
}

// takes a raw json message and marshals it into a NumJson
func marshalNumJson(r json.RawMessage) (NumJson, error) {
	var i int
	var s string
	var arr []json.RawMessage
	var obj map[string]json.RawMessage
	badInputError := fmt.Errorf("non NumJson input")

	// NumJson is a union type, so we try to unmarshal it to each subtype
	if err := json.Unmarshal(r, &i); err == nil {
		return Num(i), nil
	} else if err = json.Unmarshal(r, &s); err == nil {
		return String(s), nil
	} else if err = json.Unmarshal(r, &arr); err == nil {
		var njArray []NumJson
		for _, rawNumJson := range arr {
			marshalledJson, err := marshalNumJson(rawNumJson)
			if err != nil {
				return nil, err
			} else {
				njArray = append(njArray, marshalledJson)
			}
		}
		return Array(njArray), nil
	} else if err := json.Unmarshal(r, &obj); err == nil {
		rawPayload := obj["payload"]
		delete(obj, "payload")
		if rawPayload == nil {
			return nil, badInputError
		}
		marshalledPayloadJson, err := marshalNumJson(rawPayload)
		if err != nil {
			return nil, err
		}
		return Obj{
			Payload: marshalledPayloadJson,
			Other:   obj,
		}, nil
	} else {
		// if we cannot unmarshal the json, it's not a valid numJson
		return nil, badInputError
	}
}
