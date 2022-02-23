package jedit

import (
	"encoding/json"
	"io"
	"os"
)

func ParseStdin(stdin *os.File) (Logs, error) {
	logs := []Log{}
	decoder := json.NewDecoder(os.Stdin)
	decoder.UseNumber()
	for {
		data := make(map[string]interface{})
		err := decoder.Decode(&data)
		if err == io.EOF {
			break
		} else if err != nil {
			return Logs{}, err
		} else {
			logs = append(logs, Log{Data: data})
		}
	}
	return Logs{Data: logs}, nil
}

func ParseFilters(filters []string) ([]Filter, error) {
	err := validateFilters(filters)
	if err != nil {
		return []Filter{}, err
	}
	return getFilters(filters), nil
}
