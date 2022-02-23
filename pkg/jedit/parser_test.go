package jedit

import (
	"log"
	"os"
	"testing"
	"errors"
)

var filtersInvalid []string
var fileName string

func init() {
	fileName = "example.json"
	filtersValid = []string{
		"team == team-x",
		"severity != 4",
		"key > 88",
		"example < ppp",
		"floatTest >= 99.99",
		"runeTest <= r",
		"nospacekey==nospacevalue",
	}
	filtersInvalid = []string{
		"team == team-x",
		"severity != 4",
		"key > 88",
		"nooperator",
		"floatTest >= 99.99",
		"runeTest <= r",
		"nospacekey==nospacevalue",
	}
}

func cleanUp() {
	// Remove the testable json file
	err := os.Remove(fileName)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		log.Fatal(err)
	}
}

func createFile() (*os.File) {
    // check if file exists
    var _, err = os.Stat(fileName)

    // create file if not exists
    if os.IsNotExist(err) {
        var file, err = os.Create(fileName)
        if isError(err) {
            log.Fatal(err)
        }
        defer file.Close()
		// Add text to Json file
		json := []string{
			"{\"team\":\"team-a\",\"severity\":\"1\"}\n",
			"{\"team\":\"team-a\",\"severity\":\"2\"}\n",
			"{\"team\":\"team-b\",\"severity\":\"2\"}\n",
			"{\"team\":\"team-b\",\"severity\":\"2\"}\n",
		}
		for _,v := range json {
			_, err := file.WriteString(v)
			if err != nil {
				log.Fatal(err)
			}
		}
		return file
    }

	return nil
}

func deleteFile() {
    // delete file
    var err = os.Remove(fileName)
    if isError(err) {
        return
    }
}

func isError(err error) bool {
    if err != nil {
        log.Fatal(err)
    }
    return (err != nil)
}

func TestParseFiltersValid(t *testing.T) {
	filters, err := ParseFilters(filtersValid)
	if err != nil {
		t.Errorf("Expected ParseFilters to succeed. Failed with error: %q", err)
	}
	if len(filters) != 7 {
		t.Errorf("Expected %d filters, found %d", 7, len(filters))
	}
}

func TestParseFiltersInvalid(t *testing.T) {
	filters, err := ParseFilters(filtersInvalid)
	expectedError := "Operator not found in filter. Incorrect provided filter: nooperator\n"
	if err.Error() != expectedError {
		t.Errorf("Expected ParseFilters to fail with error message '%s', found '%s'", expectedError, err.Error())
	}
	if len(filters) != 0 {
		t.Errorf("Expected %d filters, found %d", 0, len(filters))
	}
}

func TestParseJson(t *testing.T) {
	exampleFile := createFile()
	_, err := ParseJson(exampleFile)
	if err != nil {
		t.Errorf("ParseJson returned an error unexpectedly: %q", err)
	}
	deleteFile()
}