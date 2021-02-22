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

// ScatterAxesLabels are labels for the x and y axis.
type ScatterAxesLabels struct {
	X string `json:"x"`
	Y string `json:"y"`
}

// ScatterPlot contains data for a scatter plot
type ScatterPlot struct {
	Labels ScatterAxesLabels `json:"labels"`
	Name   string            `json:"name"`
	Points []ScatterPoint    `json:"points"`
}

// ScatterPoint contains data for a point on a scatter plot
type ScatterPoint struct {
	Color string  `json:"color"`
	Label string  `json:"label"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
}

// Settings contains tool-specific analysis settings.
type Settings struct {
	// Shared settings
	Abundance                    string
	AbundanceCap                 float64
	Condition                    string
	Control                      string
	Files                        []string
	LogBase                      string
	MinAbundance                 float64
	MinConditions                int
	MockConditionAbundance       bool
	Normalization                string
	NormalizationReadout         string
	ParsimoniousReadoutFiltering bool
	Png                          bool
	PrimaryFilter                float64
	Readout                      string
	ReadoutLength                string
	Score                        string
	ScoreType                    string
	SecondaryFilter              float64
	Type                         string

	// condition-condition
	ConditionX string
	ConditionY string

	// correlation
	ConditionAbundanceFilter  float64
	ConditionScoreFilter      float64
	Correlation               string
	CytoscapeCutoff           float64
	IgnoreSourceTargetMatches bool
	ReadoutAbundanceFilter    float64
	ReadoutScoreFilter        float64
	UseReplicates             bool

	// dotplot
	BiclusteringApprox  bool
	Clustering          string
	ClusteringMethod    string
	ClusteringOptimize  bool
	ConditionClustering string
	ConditionList       []string
	Distance            string
	EdgeColor           string
	FillColor           string
	InvertColor         bool
	ReadoutClustering   string
	ReadoutList         []string
	XLabel              string
	YLabel              string
	WriteDistance       bool
	WriteHeatmap        bool

	// scv
	AbundanceFilterColumn string
	ConditionIDType       string
	ConditionMapColumn    string
	ConditionMapFile      string
	GeneFile              string
	Known                 string
	OtherAbundance        []string
	ProteinTissues        []string
	ReadoutIDType         string
	ReadoutMapColumn      string
	ReadoutMapFile        string
	RnaTissues            []string
	Specificity           bool
	VerticalHeatmap       bool

	// specificity
	SpecificityMetric string
}
