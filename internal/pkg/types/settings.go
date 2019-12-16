// Package types contains type declarations.
package types

// Analysis files, file settings, analysis settings and column mapping.
type Analysis struct {
	Columns  map[string]string
	Data     []map[string]string
	Settings Settings
}

// Settings contains tool-specific analysis settings.
type Settings struct {
	// Shared settings
	Abundance            string
	Condition            string
	Control              string
	Files                []string
	LogBase              string
	Normalization        string
	NormalizationReadout string
	Png                  bool
	PrimaryFilter        float64
	Readout              string
	ReadoutLength        string
	Score                string
	ScoreType            string
	Type                 string

	// dotplot settings
	AbundanceCap        float64
	BiclusteringApprox  bool
	Clustering          string
	ClusteringMethod    string
	ClusteringOptimize  bool
	ConditionClustering string
	ConditionList       []string
	EdgeColor           string
	FillColor           string
	Distance            string
	InvertColor         bool
	MinAbundance        float64
	ReadoutClustering   string
	ReadoutList         []string
	SecondaryFilter     float64
	XLabel              string
	YLabel              string
	WriteDistance       bool
	WriteDotplot        bool
	WriteHeatmap        bool

	// circheatmap settings
	ConditionMap   string
	Known          bool
	KnownFile      string
	OtherAbundance []string
	Species        string
	TissueFile     string
	Tissues        []string
}
