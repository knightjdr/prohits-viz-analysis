package dotplot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaitPreyMatrix(t *testing.T) {
	data := []map[string]interface{}{
		{"bait": "abait", "prey": "xprey", "abundance": "5"},
		{"bait": "abait", "prey": "zprey", "abundance": "10"},
		{"bait": "abait", "prey": "yprey", "abundance": "23"},
		{"bait": "cbait", "prey": "zprey", "abundance": "7"},
		{"bait": "cbait", "prey": "xprey", "abundance": "14.3"},
		{"bait": "bbait", "prey": "yprey", "abundance": "17.8"},
		{"bait": "bbait", "prey": "xprey", "abundance": "2"},
	}

	// TEST1: data converted to matrix.
	wantBaitList := []string{"abait", "bbait", "cbait"}
	wantMatrix := [][]float64{
		{5, 2, 14.3},
		{23, 17.8, 0},
		{10, 0, 7},
	}
	wantPreyList := []string{"xprey", "yprey", "zprey"}
	matrix, baitList, preyList := BaitPreyMatrix(data)
	assert.Equal(t, wantMatrix, matrix, "Data not converted to bait prey matrix")
	assert.Equal(t, wantBaitList, baitList, "Bait list not correct")
	assert.Equal(t, wantPreyList, preyList, "Prey list not correct")
}
