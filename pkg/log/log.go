package log

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/benmcgit/jedit/pkg/filter"
)

const (
	Equal = "=="
	NotEqual = "!="
	GreaterThen = ">"
	GreaterThenOrEqual = ">="
	LessThen = "<"
	LessThenOrEqual = "<="
)

type Log struct {
	Data  map[string]interface{}
}

type Logs struct {
	Data []Log
}

func (l *Log) Print() {
	str, err := l.ToString()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Println(str)
}

func (l *Log) ToString() (string,error) {
	a, err := json.Marshal(l.Data)
	if err != nil {
		return "", fmt.Errorf("Error converting log to string\n")
	}
	return string(a), nil
}

func (l *Log) ToBytes() ([]byte,error) {
	a, err:= json.Marshal(l.Data)
	if err != nil {
		return []byte{}, fmt.Errorf("Error converting log to byte array\n")
	}
	return a, nil
}

func (l *Logs) Print() {
	for _,v := range l.Data {
		v.Print()
	}
}

func  (l *Logs) Filter(filters []filter.Filter) {
	result := []Log{}
	for _,log := range l.Data {
		retain := true
		for _,f := range filters {
			if !retainLog(log, f) {
				retain = false
				break
			}
		}
		if retain {
			result = append(result, log)
		}
	}
	l.Data = result
}

// OR filters together
// for _,f := range filters {
// 	val, ok := data[f.Key]
// 	str := fmt.Sprintf("%v", val)
// 	if ok && str == f.Value {
// 		//fmt.Println(data)
// 		a, _ := json.Marshal(data)
// 		fmt.Println(string(a))
// 		break
// 	}
// }

func (l *Logs) SortBy(key string, ascending bool) {
    data := l.Data
    sort.SliceStable(data, func(i, j int) bool {
		if ascending {
			return fmt.Sprintf("%v", data[i].Data[key]) > fmt.Sprintf("%v", data[j].Data[key])
		}
		return fmt.Sprintf("%v", data[i].Data[key]) < fmt.Sprintf("%v", data[j].Data[key])
    })
}

func (l *Logs) Add(key string, value string, retain bool) {
    data := l.Data
	for _,d := range data {
		if v,ok := d.Data[key]; ok && retain {
			d.Data[key] = v
		} else {
			d.Data[key] = value
		}
	}
}

func (l *Logs) Remove(key string) {
    data := l.Data
	for _,d := range data {
		delete(d.Data, key)
	}
}

func (l *Logs) Modify(key string, newKey string) {
    data := l.Data
	for _,d := range data {
		if v,ok := d.Data[key]; ok {
			d.Data[newKey] = v
			delete(d.Data, key)
		}
	}
}

func retainLog(log Log, filter filter.Filter) bool {
	val, ok := log.Data[filter.Key]
	
	// cast value to string for all instances.. assures numeric types can be compared
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