package jedit

import (
	"bytes"
	"testing"
)

func TestToString(t *testing.T) {
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

func TestToBytes(t *testing.T) {
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

func TestAdd(t *testing.T) {
	resetTestData()
	expected := "{\"key1\":\"value1\",\"key2\":\"value2\",\"key3\":\"value3\",\"key4\":\"value4\"}"
	logSimple.Add("key4", "value4", false)
	str, _ := logSimple.ToString()
	if str != expected {
		t.Errorf("Output %q not equal to expected %q", str, expected)
	}
}

func TestAddRetain(t *testing.T) {
	resetTestData()
	expected := "{\"key1\":\"value1\",\"key2\":\"value2\",\"key3\":\"value3\"}"
	logSimple.Add("key3", "OriginalValueShouldBeRetained", true)
	str, _ := logSimple.ToString()
	if str != expected {
		t.Errorf("Output %q not equal to expected %q", str, expected)
	}
}

func TestRemove(t *testing.T) {
	resetTestData()
	expected := "{\"key1\":\"value1\",\"key2\":\"value2\"}"
	logSimple.Remove("key3")
	str, _ := logSimple.ToString()
	if str != expected {
		t.Errorf("Output %q not equal to expected %q", str, expected)
	}
}

func TestModify(t *testing.T) {
	resetTestData()
	expected := "{\"key1\":\"value1\",\"key2\":\"value2\",\"updated_key3\":\"value3\"}"
	logSimple.Modify("key3", "updated_key3")
	str, _ := logSimple.ToString()
	if str != expected {
		t.Errorf("Output %q not equal to expected %q", str, expected)
	}
}

func TestLogPrint(t *testing.T) {
	resetTestData()
	err := logSimple.Print()
	if err != nil {
		t.Errorf("Issue when printing log:  %q", err)
	}
}
