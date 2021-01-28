package numJson

import (
	"encoding/json"
	"fmt"
	"io"
)


type JInt int

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
				return nil, fmt.Errorf("JSON decode failed. Check that your input is valid")
			}
		}
		marshalledJson, err := marshalNumJson(r)
		if err != nil {
			return nil, err
		}
		njArray = append(njArray, marshalledJson)
	}
	return njArray, err
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
		fmt.Printf("%d\n", i)
		return NumJsonNum(i), nil
	} else if err = json.Unmarshal(r, &s); err == nil {
		fmt.Println(s)
		return NumJsonString(s), nil
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
		return NumJsonArray(njArray), nil
	} else if err := json.Unmarshal(r, &obj); err == nil {
		rawPayload := obj["payload"]
		if rawPayload == nil {
			return nil, badInputError
		}
		njObjMap := make(map[string]NumJson)
		marshalledPayloadJson, err := marshalNumJson(rawPayload)
		if err != nil {
			return nil, err
		}
		njObjMap["payload"] = marshalledPayloadJson
		return NumJsonObj(njObjMap), nil
	} else {
		return nil, badInputError
	}
}
