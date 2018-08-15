// Package typedef has type definitions needed in several packages
package typedef

// Analysis Parameters.
type Parameters struct {
	Abundance          string
	AnalysisType       string
	Bait               string
	BaitClustering     string
	BaitList           []string
	BiclusteringApprox bool
	Clustering         string
	ClusteringMethod   string
	EdgeColor          string
	FillColor          string
	Control            string
	Distance           string
	Files              []string
	Invert             bool
	LogBase            string
	MaximumAbundance   float64
	MinimumAbundance   float64
	Normalization      string
	NormalizationPrey  string
	Pdf                bool
	Png                bool
	Prey               string
	PreyClustering     string
	PreyLength         string
	PreyList           []string
	PrimaryFilter      float64
	Score              string
	ScoreType          string
	SecondaryFilter    float64
	WriteDistance      bool
	WriteDotplot       bool
	WriteHeatmap       bool
}

// Dataset.
type Dataset struct {
	Data   []map[string]interface{}
	Params Parameters
}
