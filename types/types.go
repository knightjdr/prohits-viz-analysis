// Package types has type definitions needed in several packages
package types

// Analysis Parameters.
type Parameters struct {
	Abundance         string
	AnalysisType      string
	Bait              string
	BaitList          []string
	Clustering        string
	ClusteringMethod  string
	Control           string
	Distance          string
	Files             []string
	LogBase           string
	LogFile           string
	Normalization     string
	NormalizationPrey string
	Prey              string
	PreyLength        string
	PreyList          []string
	PrimaryFilter     float64
	Score             string
	ScoreType         string
	SecondaryFilter   float64
}

// Dataset.
type Dataset struct {
	Data   []map[string]interface{}
	Params Parameters
}
