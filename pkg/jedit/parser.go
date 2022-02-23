package jedit

import (
	"encoding/json"
	"os"
	"log"
	"bufio"
)

func ParseFile(fileName string) (Logs, error) {
	logs := []Log{}

	// open file
	f, err := os.Open(fileName)
	if err != nil {
		return Logs{}, err
	}

	// close the file at the end of the function call
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(scanner.Text()), &jsonMap)
		logs = append(logs, Log{Data: jsonMap})
	}

	// assure there were no errors while scanning file
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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
