package pagerduty

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCustomFieldDataTypeIsKnown(t *testing.T) {
	for k, v := range customFieldDataTypeToString {
		if v == "unknown" {
			if k.IsKnown() {
				t.Errorf("'unknown' data type should not be known")
			}
		} else if !k.IsKnown() {
			t.Errorf("'%s' data type should be known", v)
		}
	}
}

func TestCustomFieldDataTypeIsAllowedOnField(t *testing.T) {
	for k, v := range customFieldDataTypeToString {
		if v == "unknown" || v == "field_option" {
			if k.IsAllowedOnField() {
				t.Errorf("'%s' data type should not be allowed on fields", v)
			}
		} else if !k.IsAllowedOnField() {
			t.Errorf("'%s' data type should be allowed on fields", v)
		}
	}
}

type customFieldDataTypeWrapper struct {
	DataType CustomFieldDataType `json:"datatype"`
}

func TestCustomFieldDataTypeMarshalJSON(t *testing.T) {
	for k, v := range customFieldDataTypeToString {
		o := customFieldDataTypeWrapper{DataType: k}
		b, _ := json.Marshal(o)
		s := string(b)
		exp := fmt.Sprintf(`{"datatype":"%s"}`, v)
		if s != exp {
			t.Errorf(`%s was not marshalled correctly. want:\n%s\ngot:\n%s`, v, exp, s)
		}
	}
}

func TestCustomFieldDataTypeUnmarshalJSON(t *testing.T) {
	for k, v := range customFieldDataTypeToString {
		js := fmt.Sprintf(`{"datatype":"%s"}`, v)
		var o customFieldDataTypeWrapper
		err := json.Unmarshal([]byte(js), &o)
		if err != nil {
			t.Errorf("Error when unmarhsalling %s", js)
		}
		if o.DataType != k {
			t.Errorf(`%s was not unmarshalled correctly. want:\n%s\ngot:\n%s`, js, k, o.DataType)
		}
	}
}

func TestCustomFieldDataTypeUnmarshalJSON_Error(t *testing.T) {
	js := `{"datatype":1234}`
	var o customFieldDataTypeWrapper
	err := json.Unmarshal([]byte(js), &o)
	if err == nil {
		t.Errorf("Unmarshalling %s should have produced an error, but didn't.", js)
	}
}
