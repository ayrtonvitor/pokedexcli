package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadConfig() (map[string]interface{}, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, fmt.Errorf("Could not load configs: %w\n", err.Error())
	}
	defer file.Close()

	var config map[string]interface{}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("Could not decode configs: %w\n", err.Error())
	}
	return config, nil
}
