package parser

import (
	"encoding/json"
	"io"
	"os"

	"github.com/benmcgit/jedit/pkg/log"
)

func ParseStdin(stdin *os.File) (log.Logs, error) {
	logs := []log.Log{}
	decoder := json.NewDecoder(os.Stdin)
	decoder.UseNumber()
	for {
		data := make(map[string]interface{})
		err := decoder.Decode(&data)
		if err == io.EOF {
			break
		} else if err != nil {
			return log.Logs{}, err
		} else {
			logs = append(logs, log.Log{Data: data})
		}
	}
	return log.Logs{Data: logs}, nil
}

func ParseFilters(filters []string) ([]log.Filter, error) {
	err := log.ValidateFilters(filters)
	if err != nil {
		return []log.Filter{}, err
	}
	return log.GetFilters(filters), nil
}
