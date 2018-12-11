// Package typedef has type definitions needed in several packages
package typedef

// Annotations is the struct for heatmap annotations.
type Annotations struct {
	FontSize int          `json:"fontSize"`
	List     []Annotation `json:"list"`
}

// Annotation is the struct for heatmap annotation text.
type Annotation struct {
	Text string  `json:"text"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
}

// Dataset contains data to analysis plus parameters.
type Dataset struct {
	FileData   []map[string]string
	Parameters Parameters
}

// Markers is the struct for heatmap markers.
type Markers struct {
	Color string   `json:"color"`
	List  []Marker `json:"list"`
}

// Marker is the struct for heatmap marker boxes.
type Marker struct {
	Height int `json:"height"`
	Width  int `json:"width"`
	X      int `json:"x"`
	Y      int `json:"y"`
}

// Matrices holds input data formatted as matrices
type Matrices struct {
	Abundance, Score     [][]float64
	Conditions, Readouts []string
}

// Parameters for command line arguments.
type Parameters struct {
	Abundance            string
	AbundanceCap         float64
	AnalysisType         string
	BiclusteringApprox   bool
	Clustering           string
	ClusteringMethod     string
	ClusteringOptimize   bool
	Condition            string
	ConditionClustering  string
	ConditionList        []string
	EdgeColor            string
	FillColor            string
	Control              string
	Distance             string
	Files                []string
	InvertColor          bool
	LogBase              string
	MinAbundance         float64
	Normalization        string
	NormalizationReadout string
	Pdf                  bool
	Png                  bool
	PrimaryFilter        float64
	Readout              string
	ReadoutClustering    string
	ReadoutLength        string
	ReadoutList          []string
	Score                string
	ScoreType            string
	SecondaryFilter      float64
	WriteDistance        bool
	WriteDotplot         bool
	WriteHeatmap         bool
}
