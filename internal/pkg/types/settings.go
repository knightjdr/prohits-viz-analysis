// Package types contains type declarations.
package types

// Analysis files, file settings, analysis settings and column mapping.
type Analysis struct {
	Columns  map[string]string
	Data     []map[string]string
	Settings Settings
}

// Matrices holds input data formatted as matrices.
type Matrices struct {
	Abundance, Ratio, Score [][]float64
	Conditions, Readouts    []string
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

	// circheatmap
	ConditionMap   string
	Known          bool
	KnownFile      string
	OtherAbundance []string
	Species        string
	TissueFile     string
	Tissues        []string

	// correlation
	AlwaysIncludePreysPassingFilter bool
	BaitAbundanceFilter             float64
	BaitScoreFilter                 float64
	Correlation                     string
	CytoscapeCutoff                 float64
	IgnoreSourceGenes               bool
	MinBait                         int
	MockCountsForBait               bool
	PreyAbundanceFilter             float64
	PreyScoreFilter                 float64
	UseReplicates                   bool

	// dotplot
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
	WriteHeatmap        bool
}
