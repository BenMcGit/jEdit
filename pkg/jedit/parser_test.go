package jedit

import (
	"testing"
)

var filtersInvalid []string

func init() {
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

func TestParseFileSmall(t *testing.T) {
	expected := 22
	logs, err := ParseFile("../../testdata/yesterday_reduced.json")
	if err != nil {
		t.Errorf("ParseFile returned an error unexpectedly: %q", err)
	}
	if len(logs.Data) != expected {
		t.Errorf("Expected %d logs found, recieved %d", expected, len(logs.Data))
	}
}

func TestParseFileLarge(t *testing.T) {
	expected := 14520
	logs, err := ParseFile("../../testdata/yesterday.json")
	if err != nil {
		t.Errorf("ParseFile returned an error unexpectedly: %q", err)
	}
	if len(logs.Data) != expected {
		t.Errorf("Expected %d logs found, recieved %d", expected, len(logs.Data))
	}
}
