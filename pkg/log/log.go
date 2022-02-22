package log

import (
	"encoding/json"
	"fmt"
)

type Log struct {
	Data  map[string]interface{}
}

func (l *Log) Print() error {
	str, err := l.ToString()
	if err != nil {
		return err
	}
	fmt.Println(str)
	return nil
}

func (l *Log) ToString() (string,error) {
	a, err := json.Marshal(l.Data)
	if err != nil {
		return "", err
	}
	return string(a), nil
}

func (l *Log) ToBytes() ([]byte,error) {
	a, err:= json.Marshal(l.Data)
	if err != nil {
		return []byte{}, err
	}
	return a, nil
}

func (l *Log) Add(key string, value string, retain bool) {
    data := l.Data
	if v,ok := data[key]; ok && retain {
		data[key] = v
	} else {
		data[key] = value
	}
}

func (l *Log) Remove(key string) {
    data := l.Data
	delete(data, key)
}

func (l *Log) Modify(key string, newKey string) {
    data := l.Data
	if v,ok := data[key]; ok {
		data[newKey] = v
		delete(data, key)
	}
}
