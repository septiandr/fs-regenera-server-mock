package utils

import (
	"encoding/json"
	"os"
)

func ReadJSONFile(path string, target interface{}) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(file, target); err != nil {
		return err
	}

	return nil
}
