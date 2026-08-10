package main

import (
	"encoding/json"
	"os"
	"sort"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	var sorted []string
	for key := range pages {
		sorted = append(sorted, key)
	}

	sort.Strings(sorted)
	sortedData := map[string]PageData{}
	for _, key := range sorted {
		sortedData[key] = pages[key]
	}

	data, err := json.MarshalIndent(sortedData, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
