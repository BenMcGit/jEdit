package parser

import (
	"encoding/json"
	"os"
	"io"

	"github.com/benmcgit/jedit/pkg/log"
)

func ParseStdin(stdin *os.File) log.Logs {
	logs := []log.Log{}
	decoder := json.NewDecoder(os.Stdin)
	decoder.UseNumber()
	for {
		data := make(map[string]interface{})
		err := decoder.Decode(&data)
		if err == io.EOF {
            break
        } else if err != nil {
			break
		} else {
			logs = append(logs, log.Log{Data: data})
		}
	}
	return log.Logs{Data:logs}
}