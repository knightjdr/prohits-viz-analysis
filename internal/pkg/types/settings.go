// Package types contains type declarations.
package types

// Analysis settings.
type Analysis struct {
	Columns  map[string]string
	FileData []map[string]string
	Settings interface{}
	Type     string
}

// CircHeatmap contains command line argument for the circular heat map tool.
type CircHeatmap struct {
	File           `json:"file"`
	ConditionMap   string
	Known          bool
	KnownFile      string
	OtherAbundance []string
	Png            bool
	Species        string
	TissueFile     string
	Tissues        []string
}

// Dotplot contains command line argument for the dot plot tool.
type Dotplot struct {
	File                `json:"file"`
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
	Png                 bool
	ReadoutClustering   string
	ReadoutList         []string
	SecondaryFilter     float64
	XLabel              string
	YLabel              string
	WriteDistance       bool
	WriteDotplot        bool
	WriteHeatmap        bool
}

// File contains settings for reading and transforming the input file.
type File struct {
	Abundance            string
	Condition            string
	Control              string
	Files                []string
	LogBase              string
	Normalization        string
	NormalizationReadout string
	PrimaryFilter        float64
	Readout              string
	ReadoutLength        string
	Score                string
	ScoreType            string
}
