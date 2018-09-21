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
	Data   []map[string]interface{}
	Params Parameters
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

// Parameters for command line arguments.
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
