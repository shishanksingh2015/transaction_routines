package utils

import (
	"encoding/json"
	"io"
)

func ReadInto(res io.Reader, receiver any) error {
	data, err := io.ReadAll(res)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, receiver)
}

func StructToJson(data any) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
