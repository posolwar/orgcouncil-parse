package filters

import (
	"encoding/json"
	"os"
)

func GetFiltersFromFile(path string) (map[string]string, error) {
	if path == "" {
		return nil, nil
	}

	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	payload := make(map[string]string)

	err = json.Unmarshal(fileContent, &payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
