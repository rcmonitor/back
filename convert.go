package main

import (
	"encoding/json"
)

func getJson(properties []string, data []string) (strResponse string) {
	fields := map[string]string{}
	for i, field := range properties {
		fields[field] = data[i]
	}
	jsonBytes, err := json.MarshalIndent(fields, "", " ")
	if err == nil {
		strResponse = string(jsonBytes)
	}

	return
}

