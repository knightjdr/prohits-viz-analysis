package columnparser

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderMap(t *testing.T) {
	// TEST: both columns can be found
	columnMap := map[string]string{
		"key1": "column1",
		"key2": "column3",
	}
	header := []string{"column1", "column2", "column3"}
	headerMap, err := HeaderMap(columnMap, header, false)
	var want = map[string]int{
		"key1": 0,
		"key2": 2,
	}
	assert.Nil(t, err, "expected no error, got %s", err)
	assert.True(
		t,
		reflect.DeepEqual(headerMap, want),
		"HeaderMap(%v, %v) == %v, want %v", columnMap, header, headerMap, want,
	)

	// TEST: a column can not be found
	columnMap = map[string]string{
		"key1": "column1",
		"key2": "column4",
	}
	_, err = HeaderMap(columnMap, header, false)
	assert.NotNil(t, err, "expected error when a column cannot be found in the header")

	// TEST: a column can not be found but we are ignoring missing columns
	columnMap = map[string]string{
		"key1": "column1",
		"key2": "column4",
	}
	_, err = HeaderMap(columnMap, header, true)
	assert.Nil(t, err, "expected no error when a column is missing and we are ignoring those cases")

	// TEST: empty map values should be ignored
	columnMap = map[string]string{
		"key1": "column1",
		"key2": "column3",
		"key3": "",
	}
	headerMap, err = HeaderMap(columnMap, header, false)
	assert.Nil(t, err, "expected no error, got %s", err)
	assert.True(
		t,
		reflect.DeepEqual(headerMap, want),
		"HeaderMap(%v, %v) == %v, want %v", columnMap, header, headerMap, want,
	)
}
