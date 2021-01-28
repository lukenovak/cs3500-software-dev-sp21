package numJson

import (
	"encoding/json"
	"fmt"
	"log"
)


type JInt int

// Decodes into string
func ParseNumJsonFromStream(d *json.Decoder) (string, error) {
	for {
		var r json.RawMessage
		if err := d.Decode(&r); err != nil {
			println("decode failed")
			break
		}
		var i int
		var s string
		var arr []interface{}
		var obj map[string]interface{}
		if err := json.Unmarshal(r, &i); err == nil {
			fmt.Printf("%d\n", i)
		} else if err = json.Unmarshal(r, &s); err == nil {
			fmt.Println(s)
		} else if err = json.Unmarshal(r, &arr); err == nil {
			fmt.Println("array of shit")
		} else if err := json.Unmarshal(r, &obj); err != nil {
			log.Println(err)
			return "", err
		} else {
			println(obj)
		}
	}
	return "", nil
}