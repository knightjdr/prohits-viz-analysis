package columnparser

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeadermap(t *testing.T) {
	// TEST1: both columns can be found
	columnMap := map[string]string{
		"key1": "column1",
		"key2": "column3",
	}
	header := []string{"column1", "column2", "column3"}
	var want = map[string]int{
		"key1": 0,
		"key2": 2,
	}
	headerMap, err := Headermap(columnMap, header)
	assert.Nil(t, err, "expected no error, got %s", err)
	assert.True(
		t,
		reflect.DeepEqual(headerMap, want),
		"HeaderMap(%v, %v) == %v, want %v", columnMap, header, headerMap, want,
	)

	// TEST2: a column can not be found
	columnMap = map[string]string{
		"key1": "column1",
		"key2": "column4",
	}
	headerMap, err = Headermap(columnMap, header)
	assert.NotNil(t, err, "expected error when a column cannot be found in the header")

	// TEST3: empty map values should be ignored
	columnMap = map[string]string{
		"key1": "column1",
		"key2": "column3",
		"key3": "",
	}
	headerMap, err = Headermap(columnMap, header)
	assert.Nil(t, err, "expected no error, got %s", err)
	assert.True(
		t,
		reflect.DeepEqual(headerMap, want),
		"HeaderMap(%v, %v) == %v, want %v", columnMap, header, headerMap, want,
	)
}
