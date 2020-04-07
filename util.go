package gotokens

import (
	"encoding/json"
	"fmt"
	"os"
)

const tokenFolder string = "/Tokens/"

func findJSONFile(fn string) (string, error) {
	if len(tokenSearchPath) == 0 {
		return "", &Error{
			Code:    ENOSEARCHPATH,
			Message: fmt.Sprint("Search Path not set"),
		}
	}

	for _, path := range tokenSearchPath {
		tf := path + tokenFolder + fn
		if fileExists(tf) {
			return tf, nil
		}
	}

	return "", &Error{
		Code:    EFILENOTFOUND,
		Message: fmt.Sprintf("File '%s' not found in Search Path: %v", fn, tokenSearchPath),
	}
}

func fileExists(fn string) bool {
	info, err := os.Stat(fn)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func convert(byteValue []byte) (interface{}, error) {
	const op string = "gotokens.convert"
	var jsonData map[string]interface{}
	err := json.Unmarshal([]byte(byteValue), &jsonData)
	if err != nil {
		return nil, &Error{Op: op, Context: "json.Unmarshal", Err: err}
	}

	var data interface{}
	var ok bool
	if data, ok = jsonData["tokens"]; !ok {
		return nil, &Error{
			Code:    EMALFORMEDJSON,
			Message: fmt.Sprintf("Malformed file: 'tokens' missing"),
		}
	}

	return data, nil
}
