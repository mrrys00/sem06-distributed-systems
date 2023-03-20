package utils

import (
	"encoding/json"
	"io"
)

func UnmarshalJson(r io.ReadCloser, target interface{}) error {
	return json.NewDecoder(r).Decode(target)
}

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}

	return string(val), nil
}
