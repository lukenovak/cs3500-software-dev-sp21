package numJson

import (
	"encoding/json"
	"fmt"
)


type JInt int

// Decodes into string
func ParseNumJsonFromStream(d *json.Decoder) ([]NumJson, error) {
	var njArray []NumJson
	var err error
	for {
		var r json.RawMessage
		if err = d.Decode(&r); err != nil {
			println("decode failed")
			break
		}
		njArray = append(njArray, marshalNumJson(r))
	}
	return njArray, err
}

// takes a raw json message and marshals it into a NumJson
func marshalNumJson(r json.RawMessage) NumJson {
	var i int
	var s string
	var arr []json.RawMessage
	var obj map[string]json.RawMessage

	// NumJson is a union type, so we try to unmarshal it to each subtype
	if err := json.Unmarshal(r, &i); err == nil {
		fmt.Printf("%d\n", i)
		return NumJsonNum(i)
	} else if err = json.Unmarshal(r, &s); err == nil {
		fmt.Println(s)
		return NumJsonString(s)
	} else if err = json.Unmarshal(r, &arr); err == nil {
		var njArray []NumJson
		for _, rawNumJson := range arr {
			njArray = append(njArray, marshalNumJson(rawNumJson))
		}
		return NumJsonArray(njArray)
	} else if err := json.Unmarshal(r, &obj); err == nil {
		rawPayload := obj["payload"]
		njObjMap := make(map[string]NumJson)
		njObjMap["payload"] = marshalNumJson(rawPayload)
		return NumJsonObj(njObjMap)
	} else {
		panic(obj)
	}
}