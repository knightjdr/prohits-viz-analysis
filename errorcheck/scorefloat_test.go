package errorcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreFloat(t *testing.T) {
	// TEST1: valid score column returns no error.
	data := []map[string]string{
		{"score": "0.1"},
	}
	err := ScoreFloat(data)
	assert.Nil(t, err, "Valid score column should not return an error")

	// TEST2: invalid score column returns error.
	tests := [4]string{"a", "", "true"}
	for _, value := range tests {
		data = []map[string]string{
			{"score": value},
		}
		err = ScoreFloat(data)
		assert.NotNil(t, err, "Invalid score column should return an error")
	}
}
