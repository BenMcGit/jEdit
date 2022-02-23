package jedit

import (
	"testing"
)

var filtersValid []string
var filtersInvalidMap map[string]string

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

	// needs to be a map because Validate will return an error
	// on the first invalid filter it sees
	filtersInvalidMap = make(map[string]string)
	filtersInvalidMap["nooperator"] = "Operator not found in filter. Incorrect provided filter: nooperator\n"
	filtersInvalidMap["keynovalue =="] = "Filter not formatted correctly. Please use format 'key == value'. Incorrect provided filter: keynovalue ==\n"
	filtersInvalidMap["== valuenokey"] = "Filter not formatted correctly. Please use format 'key == value'. Incorrect provided filter: == valuenokey\n"
	filtersInvalidMap["=="] = "Filter not formatted correctly. Please use format 'key == value'. Incorrect provided filter: ==\n"
	filtersInvalidMap[">= < adf"] = "Filter not formatted correctly. Please use format 'key == value'. Incorrect provided filter: >= < adf\n"
	filtersInvalidMap["key =a= value"] = "Operator not found in filter. Incorrect provided filter: key =a= value\n"

}

func TestValidateFiltersPositive(t *testing.T) {
	err := validateFilters(filtersValid)
	if err != nil {
		t.Errorf("Expected data to be valid but found this error: %q", err)
	}
}

func TestValidateFiltersNegative(t *testing.T) {
	for k, v := range filtersInvalidMap {
		err := validateFilters([]string{k})
		if err == nil {
			t.Errorf("Expected data to be invalid but found the error was nil. Provided key: %s", k)
		}
		if err.Error() != v {
			t.Errorf("Expected error message '%s', found '%s'", v, err.Error())
		}
	}
}

func TestGetOperator(t *testing.T) {
	expected := []string{"==", "!=", ">", "<", ">=", "<=", "=="}
	for i := 0; i < len(expected); i++ {
		op, err := getOperator(filtersValid[i])
		if err != nil {
			t.Errorf("Expected data to be valid but found this error: %q", err)
		}
		if op != expected[i] {
			t.Errorf("Expected operator to be %s, found %s", expected[i], op)
		}
	}
}

func TestGetFilters(t *testing.T) {
	filters := getFilters(filtersValid)
	expectedKey := []string{"team", "severity", "key", "example", "floatTest", "runeTest", "nospacekey"}
	expectedValue := []string{"team-x", "4", "88", "ppp", "99.99", "r", "nospacevalue"}
	expectedOp := []string{"==", "!=", ">", "<", ">=", "<=", "=="}
	for i, f := range filters {
		if f.Key != expectedKey[i] || f.Value != expectedValue[i] || f.Operation != expectedOp[i] {
			t.Errorf("Filter contains unexpected data. Expected key='%s', value='%s', operator='%s'. Recieved %v",
				expectedKey[i], expectedValue[i], expectedOp[i], f)
		}
	}
}

func TestGetFiltersEmpty(t *testing.T) {
	filters := getFilters([]string{})
	if len(filters) != 0 {
		t.Errorf("Expected %d filters, found %d", 0, len(filters))
	}
}
