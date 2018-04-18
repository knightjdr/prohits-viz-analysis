package columnparser

import "reflect"
import "testing"

func TestHeadermap(t *testing.T) {
	// both columns can be found
	columns := []string{"column1", "column3"}
	header := []string{"column1", "column2", "column3"}
	var want = map[string]int{
		"column1": 0,
		"column3": 2,
	}
	headerMap, err := Headermap(columns, header)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if !reflect.DeepEqual(headerMap, want) {
		t.Errorf("HeaderMap(%v, %v) == %v, want %v", columns, header, headerMap, want)
	}

	// a column can not be found
	columns = []string{"column1", "column4"}
	headerMap, err = Headermap(columns, header)
	if err == nil {
		t.Errorf("Expected an error when a column cannot be found in the header")
	}
}
