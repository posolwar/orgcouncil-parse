package filters

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func GetFiltersFromFile(path string) (map[string]string, error) {
	if path == "" {
		return nil, nil
	}

	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	payload := make(map[string]string)

	err = json.Unmarshal(fileContent, &payload)
	if err != nil {
		return nil, err
	}

	if len(payload) == 0 {
		return nil, errors.New("параметры не могут быть пустыми")
	}

	return payload, nil
}
