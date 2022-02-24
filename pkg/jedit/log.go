package jedit

import (
	"encoding/json"
	"fmt"
)

type Log struct {
	Data map[string]interface{}
}

func (l *Log) print() error {
	str, err := l.toString()
	if err != nil {
		return err
	}
	fmt.Println(str)
	return nil
}

func (l *Log) toString() (string, error) {
	a, err := json.Marshal(l.Data)
	if err != nil {
		return "", err
	}
	return string(a), nil
}

func (l *Log) toBytes() ([]byte, error) {
	a, err := json.Marshal(l.Data)
	if err != nil {
		return []byte{}, err
	}
	return a, nil
}

func (l *Log) add(key string, value string, retain bool) {
	data := l.Data
	if v, ok := data[key]; ok && retain {
		data[key] = v
	} else {
		data[key] = value
	}
}

func (l *Log) remove(key string) {
	data := l.Data
	delete(data, key)
}

func (l *Log) modify(key string, newKey string) {
	data := l.Data
	if v, ok := data[key]; ok {
		data[newKey] = v
		delete(data, key)
	}
}
