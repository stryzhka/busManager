package customErrors

import "encoding/json"

type JsonError struct {
	Error string
}

func NewJsonError(err error) string {
	jsonError := &JsonError{err.Error()}
	data, _ := json.MarshalIndent(jsonError, "", "    ")
	return string(data)
}
