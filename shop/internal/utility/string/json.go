package string

import "encoding/json"

func ParseObjectToJsonString(input any) (output string) {
	b, _ := json.Marshal(input)
	return string(b)
}
