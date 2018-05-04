package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterFunc(t *testing.T) {
	filter := 0.5

	// TEST1: gte returns func that considers higher scores better.
	filterFunc := FilterFunc("gte")
	assert.True(t, filterFunc(0.8, filter), "Scores higher than filter should pass FilterFunc")
	assert.True(t, filterFunc(0.5, filter), "Scores equal to filter should pass FilterFunc")
	assert.False(t, filterFunc(0.4, filter), "Scores less then filter should no pass FilterFunc")

	// TEST2: lte returns func that considers lower scores better.
	filterFunc = FilterFunc("lte")
	assert.True(t, filterFunc(0.4, filter), "Scores less then filter should pass FilterFunc")
	assert.True(t, filterFunc(0.5, filter), "Scores equal to filter should pass FilterFunc")
	assert.False(t, filterFunc(0.8, filter), "Scores higher than filter should not pass FilterFunc")

	// TEST3: default returns func that considers lower scores better.
	filterFunc = FilterFunc("")
	assert.True(t, filterFunc(0.4, filter), "Scores less then filter should pass FilterFunc")
	assert.True(t, filterFunc(0.5, filter), "Scores equal to filter should pass FilterFunc")
	assert.False(t, filterFunc(0.8, filter), "Scores higher than filter should not pass FilterFunc")
}
