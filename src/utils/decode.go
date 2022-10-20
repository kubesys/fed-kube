package utils

import (
	"encoding/json"
	"reflect"
)

/*
decode json to struct
*/
func ConstructStructOfPtr(structPtr interface{}, content []byte) error {
	//CreateRequest function return a pointer to the request
	err := json.Unmarshal(content, structPtr)
	return err
}

/*
decode json to struct value of reflect.Value
*/
func ConstructStructOfValue(structValue reflect.Value, content []byte) error {
	structInterface := structValue.Interface()
	err := json.Unmarshal(content, structInterface)
	return err
}
