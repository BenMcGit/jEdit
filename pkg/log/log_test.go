package log

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/benmcgit/jedit/pkg/filter"
)


var logSimple Log
var logsSimple Logs
var logAdvanced Log
var logsAdvanced Logs

func init() {
	resetTestData()
}

func resetTestData() {
	// simple test data
    logSimple = Log{Data: make(map[string]interface{})}
	logSimple.Data["key1"] = "value1"
	logSimple.Data["key2"] = "value2"
	logSimple.Data["key3"] = "value3"

	logsSimple = Logs{Data: []Log{}}
	logsSimple.Data = append(logsSimple.Data, logSimple)

	// advanaced test data
	l1 := Log{Data: make(map[string]interface{})}
	l1.Data["team"] = "team-a"
	l1.Data["severity"] = "1"
	l2 := Log{Data: make(map[string]interface{})}
	l2.Data["team"] = "team-a"
	l2.Data["severity"] = "2"
	l3 := Log{Data: make(map[string]interface{})}
	l3.Data["team"] = "team-b"
	l3.Data["severity"] = "2"
	l4 := Log{Data: make(map[string]interface{})}
	l4.Data["team"] = "team-b"
	l4.Data["severity"] = "3"

	logsAdvanced = Logs{Data: []Log{}}
	// append in a different order to validate the sort function
	logsAdvanced.Data = append(logsAdvanced.Data, l4)
	logsAdvanced.Data = append(logsAdvanced.Data, l2)
	logsAdvanced.Data = append(logsAdvanced.Data, l1)
	logsAdvanced.Data = append(logsAdvanced.Data, l3)
}

func TestToString(t *testing.T){
	resetTestData()
	expected := "{\"key1\":\"value1\",\"key2\":\"value2\",\"key3\":\"value3\"}"
	str, err := logSimple.ToString()
	if err != nil {
		t.Errorf("Error occurred in ToString: %q", err)
	}
	if str != expected {
		t.Errorf("Output %q not equal to expected %q", str, expected)
	}
}

func TestToBytes(t *testing.T){
	resetTestData()
	expected := []byte("{\"key1\":\"value1\",\"key2\":\"value2\",\"key3\":\"value3\"}")
	b, err := logSimple.ToBytes()
	if err != nil {
		t.Errorf("Error occurred in ToBytes: %q", err)
	}
	if res := bytes.Compare(b, expected); res != 0 {
		t.Errorf("Output %q not equal to expected %q", b, expected)
	}
}

func TestRetainLogEquals(t *testing.T){
	resetTestData()
	operator := "=="

	// Positive use case
	expected := true
	fil := filter.Filter{Key: "key2", Value: "value2", Operation: operator}
	result := retainLog(logSimple, fil)
	if result != expected {
		t.Errorf("Output %t not equal to expected %t", result, expected)
	}

	// Negative use case
	expected = false
	fil = filter.Filter{Key: "key2", Value: "value1", Operation: operator}
	result = retainLog(logSimple, fil)
	if result != expected {
		t.Errorf("Output %t not equal to expected %t", result, expected)
	}
}

func TestRetainLogNotEquals(t *testing.T){
	resetTestData()
	operator := "!="

	// Positive use case
	expected := true
	fil := filter.Filter{Key: "key2", Value: "value1", Operation: operator}
	result := retainLog(logSimple, fil)
	if result != expected {
		t.Errorf("Output %t not equal to expected %t", result, expected)
	}

	// Negative use case
	expected = false
	fil = filter.Filter{Key: "key2", Value: "value2", Operation: operator}
	result = retainLog(logSimple, fil)
	if result != expected {
		t.Errorf("Output %t not equal to expected %t", result, expected)
	}
}

func TestRetainLogGreaterThan(t *testing.T){
	resetTestData()
	operators := []string{">=", ">"}

	for _,operator := range operators {
		// Positive use case
		expected := true
		fil := filter.Filter{Key: "key2", Value: "value1", Operation: operator}
		result := retainLog(logSimple, fil)
		if result != expected {
			t.Errorf("Output %t not equal to expected %t", result, expected)
		}

		// Negative use case
		expected = false
		fil = filter.Filter{Key: "key2", Value: "value3", Operation: operator}
		result = retainLog(logSimple, fil)
		if result != expected {
			t.Errorf("Output %t not equal to expected %t", result, expected)
		}
	}
}

func TestRetainLogLessThan(t *testing.T){
	resetTestData()
	operators := []string{"<=", "<"}

	for _,operator := range operators {
		// Positive use case
		expected := true
		fil := filter.Filter{Key: "key2", Value: "value3", Operation: operator}
		result := retainLog(logSimple, fil)
		if result != expected {
			t.Errorf("Output %t not equal to expected %t", result, expected)
		}

		// Negative use case
		expected = false
		fil = filter.Filter{Key: "key2", Value: "value1", Operation: operator}
		result = retainLog(logSimple, fil)
		if result != expected {
			t.Errorf("Output %t not equal to expected %t", result, expected)
		}
	}
}

func TestLogsAdd(t *testing.T){
	resetTestData()
	k,v := "newKey", "newValue"

	logsSimple.Add(k, v, false)
	if logsSimple.Data[0].Data[k] != v {
		t.Errorf("Output %s not equal to expected %s", logsSimple.Data[0].Data[k], v)
	}
}

func TestLogsAddRetain(t *testing.T){
	resetTestData()
	k,v := "key1", "thisshouldnotbepresent"

	expected := "value1"
	logsSimple.Add(k, v, true)
	if logsSimple.Data[0].Data[k] != expected {
		t.Errorf("Output %s not equal to expected %s", logsSimple.Data[0].Data[k], expected)
	}
}

func TestLogsDelete(t *testing.T){
	resetTestData()
	k := "key1"

	logsSimple.Remove(k)
	if _,ok := logsSimple.Data[0].Data[k]; ok {
		t.Errorf("The key %s was not removed as expected.", k)
	}
}

func TestLogsModify(t *testing.T){
	resetTestData()
	k, kNew := "key2", "key_new_2"

	logsSimple.Modify(k, kNew)
	if _,ok := logsSimple.Data[0].Data[k]; ok {
		t.Errorf("The key %s was not removed as expected. Data: %v", k, logsSimple.Data[0].Data)
	}
	if _,ok := logsSimple.Data[0].Data[kNew]; !ok {
		t.Errorf("The key %s was not added as expected. Data: %v", kNew, logsSimple.Data[0].Data)
	}
}


func TestLogsSortBy(t *testing.T){
	resetTestData()
	pre, post, postAsc := []string{"3", "2", "1", "2"}, []string{"1", "2", "2", "3"}, []string{"3", "2", "2", "1"}
	key := "severity"

	for i,log := range logsAdvanced.Data {
		if log.Data[key] != pre[i] {
			fmt.Println(log.Data)
			t.Errorf("Output %s not equal to expected %s", log.Data[key], pre[i])
		}
	}
	// validate sort asc=false
	logsAdvanced.SortBy(key, false)
	for i,log := range logsAdvanced.Data {
		if log.Data[key] != post[i] {
			t.Errorf("Output %s not equal to expected %s", log.Data[key], post[i])
		}
	}
	logsAdvanced.SortBy(key, true)
	// validate sort asc=true
	for i,log := range logsAdvanced.Data {
		if log.Data[key] != postAsc[i] {
			t.Errorf("Output %s not equal to expected %s", log.Data[key], postAsc[i])
		}
	}
}

func TestLogsFilter(t *testing.T){
	resetTestData()
	preSev, postSev := []string{"3", "2", "1", "2"}, []string{"3", "2"}
	preTeam, postTeam := []string{"team-b", "team-a", "team-a", "team-b"}, []string{"team-b", "team-b"}
	keySev, keyTeam := "severity", "team"
	fils := []filter.Filter{
		{Key: "severity", Value: "2", Operation: ">="},
		{Key: "team", Value: "team-b", Operation: "=="},
	}
	if len(logsAdvanced.Data) != 4 {
		t.Errorf("Expected %d records. Found %d.", 4, len(logsAdvanced.Data))
	}
	for i,log := range logsAdvanced.Data {
		if log.Data[keySev] != preSev[i] {
			t.Errorf("Output %s not equal to expected %s", log.Data[keySev], preSev[i])
		}
		if log.Data[keyTeam] != preTeam[i] {
			t.Errorf("Output %s not equal to expected %s", log.Data[keyTeam], preTeam[i])
		}
	}
	logsAdvanced.Filter(fils)
	if len(logsAdvanced.Data) != 2 {
		t.Errorf("Expected %d records. Found %d.", 2, len(logsAdvanced.Data))
	}
	for i,log := range logsAdvanced.Data {
		if log.Data[keySev] != postSev[i] {
			t.Errorf("Output %s not equal to expected %s", log.Data[keySev], postSev[i])
		}
		if log.Data[keyTeam] != postTeam[i] {
			t.Errorf("Output %s not equal to expected %s", log.Data[keyTeam], postTeam[i])
		}
	}
}