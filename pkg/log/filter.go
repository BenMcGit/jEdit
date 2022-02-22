package log

import (
	"fmt"
	"strings"
)

const (
	Equal              = "=="
	NotEqual           = "!="
	GreaterThen        = ">"
	GreaterThenOrEqual = ">="
	LessThen           = "<"
	LessThenOrEqual    = "<="
)

type Filter struct {
	Key       string
	Value     string
	Operation string
}

func GetFilters(filters []string) []Filter {
	if len(filters) == 0 {
		return []Filter{}
	}
	result := make([]Filter, len(filters))
	for i, f := range filters {
		op, _ := getOperator(f)
		filterArr := strings.Split(f, op)
		filter := Filter{
			Key:       strings.TrimSpace(filterArr[0]),
			Value:     strings.TrimSpace(filterArr[1]),
			Operation: op,
		}
		result[i] = filter
	}
	return result
}

func ValidateFilters(filters []string) error {
	for _, f := range filters {
		op, err := getOperator(f)
		if err != nil {
			return err
		}
		filterSplit := strings.Split(f, op)
		if len(filterSplit) != 2 || len(filterSplit[0]) == 0 || len(filterSplit[1]) == 0 {
			return fmt.Errorf("Filter not formatted correctly. Please use format 'key == value'. Incorrect provided filter: %s\n", f)
		}
	}
	return nil
}

func getOperator(filter string) (string, error) {
	// The order matters here. Need to GreaterThenOrEqual and LessThenOrEqual before GreaterThen and LessThen.
	// If not, a false positive could occur.
	operators := []string{Equal, NotEqual, GreaterThenOrEqual, LessThenOrEqual, GreaterThen, LessThen}
	for _, op := range operators {
		if strings.Contains(filter, op) {
			return op, nil
		}
	}
	return "", fmt.Errorf("Operator not found in filter. Incorrect provided filter: %s\n", filter)
}
