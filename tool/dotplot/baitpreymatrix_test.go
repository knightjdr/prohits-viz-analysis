package dotplot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaitPreyMatrix(t *testing.T) {
	dataset := []map[string]interface{}{
		{"bait": "abait", "prey": "xprey", "abundance": "5", "score": 0.01},
		{"bait": "abait", "prey": "zprey", "abundance": "10", "score": 0.02},
		{"bait": "abait", "prey": "yprey", "abundance": "23", "score": float64(0)},
		{"bait": "cbait", "prey": "zprey", "abundance": "7", "score": 0.01},
		{"bait": "cbait", "prey": "xprey", "abundance": "14.3", "score": 0.08},
		{"bait": "bbait", "prey": "yprey", "abundance": "17.8", "score": 0.01},
		{"bait": "bbait", "prey": "xprey", "abundance": "2", "score": 0.01},
	}

	// TEST1: dataset converted to matrix with smaller scores better.
	wantBaitList := []string{"abait", "bbait", "cbait"}
	wantAbundance := [][]float64{
		{5, 2, 14.3},
		{23, 17.8, 0},
		{10, 0, 7},
	}
	wantPreyList := []string{"xprey", "yprey", "zprey"}
	wantScore := [][]float64{
		{0.01, 0.01, 0.08},
		{0, 0.01, 0.08},
		{0.02, 0.08, 0.01},
	}
	data := BaitPreyMatrix(dataset, "lte")
	assert.Equal(t, wantAbundance, data.Abundance, "Data not converted to bait prey abundance matrix")
	assert.Equal(t, wantBaitList, data.Baits, "Bait list not correct")
	assert.Equal(t, wantPreyList, data.Preys, "Prey list not correct")
	assert.Equal(t, wantScore, data.Score, "Data not converted to bait prey score matrix")

	// TEST2: dataset converted to matrix with larger scores better.
	wantScore = [][]float64{
		{0.01, 0.01, 0.08},
		{0, 0.01, 0},
		{0.02, 0, 0.01},
	}
	data = BaitPreyMatrix(dataset, "gte")
	assert.Equal(t, wantScore, data.Score, "Data not converted to bait prey score matrix")
}
