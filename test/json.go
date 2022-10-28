package test

import (
	"encoding/json"
	"reflect"
)

func CompareToJSON(data interface{}, str string) bool {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return false
	}
	var dataMp, jsonMp map[string]interface{}
	if err := json.Unmarshal(dataJSON, &dataMp); err != nil {
		return false
	}
	if err := json.Unmarshal([]byte(str), &jsonMp); err != nil {
		return false
	}
	return reflect.DeepEqual(dataMp, jsonMp)
}
