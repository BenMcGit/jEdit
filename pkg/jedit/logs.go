package jedit

import (
	"os"
	"io"
	"fmt"
	"sort"
)

const (
	Add = "addConst"
	Reduce = "reduceConst"
	Remove = "removeConst"
	Modify = "modifyConst"
)

type Logs struct {
	Data []Log
}

func (l *Logs) Print() error {
	for _,log := range l.Data {
		err := log.print()
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *Logs) WriteToFile(fileName string) error {
    file, err := os.Create(fileName)
    if err != nil {
        return err
    }
	// close the file reguardless if an error occurs below
    defer file.Close()

	for _,log := range l.Data {
		str, err := log.toString()
		if err != nil {
			return err
		}
		_, err = io.WriteString(file, str + "\n")
		if err != nil {
			return err
		}
	}
    return file.Sync()
}

func  (l *Logs) Filter(filters []Filter) {
	result := []Log{}
	for _,log := range l.Data {
		retain := checkLogMatchesFilters(log, filters)
		if retain {
			result = append(result, log)
		}
	}
	l.Data = result
}

func (l *Logs) SortBy(key string, ascending bool) {
    data := l.Data
    sort.SliceStable(data, func(i, j int) bool {
		if ascending {
			return fmt.Sprintf("%v", data[i].Data[key]) > fmt.Sprintf("%v", data[j].Data[key])
		}
		return fmt.Sprintf("%v", data[i].Data[key]) < fmt.Sprintf("%v", data[j].Data[key])
    })
}

func (l *Logs) Add(key string, value string, retain bool, filters []Filter) {
	empty := isFiltersEmpty(filters)
	for _,log := range l.Data {
		if empty || checkLogMatchesFilters(log, filters) {
			log.add(key, value, retain)
		}
	}
}

func (l *Logs) Remove(key string, filters []Filter) {
	empty := isFiltersEmpty(filters)
	for _,log := range l.Data {
		if empty || checkLogMatchesFilters(log, filters) {
			log.remove(key)
		}
	}
}

func (l *Logs) Modify(key string, newKey string, filters []Filter) {
	empty := isFiltersEmpty(filters)
	for _,log := range l.Data {
		if empty || checkLogMatchesFilters(log, filters) {
			log.modify(key, newKey)
		}
	}
}

func isMatch(log Log, filter Filter) bool {
	val, ok := log.Data[filter.Key]
	
	// cast value to string for all instances
	// assures numeric types can be compared
	str := fmt.Sprintf("%v", val)

	var result bool
	if filter.Operation == Equal {
		result = ok && str == filter.Value 
	} else if filter.Operation == NotEqual {
		result = ok && str != filter.Value
	} else if filter.Operation == GreaterThen {
		result = ok && str > filter.Value
	} else if filter.Operation == GreaterThenOrEqual {
		result = ok && (str > filter.Value || str == filter.Value)
	} else if filter.Operation == LessThen {
		result = ok && str < filter.Value
	} else if filter.Operation == LessThenOrEqual {
		result = ok && (str < filter.Value || str == filter.Value)
	}
	return result
}

func checkLogMatchesFilters(log Log, filters []Filter) bool {
	retain := true
	for _,f := range filters {
		if !isMatch(log, f) {
			retain = false
			break
		}
	}
	return retain
}

func isFiltersEmpty(filters []Filter) bool {
	return filters == nil || len(filters) == 0
}
