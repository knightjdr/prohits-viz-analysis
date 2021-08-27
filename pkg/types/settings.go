// Package types contains type declarations.
package types

import "strconv"

// Analysis files, file settings, analysis settings and column mapping.
type Analysis struct {
	Columns  map[string]string
	Data     []map[string]string
	Settings Settings
}

// CircHeatmap is a circular heatmap plot
type CircHeatmap struct {
	Name     string               `json:"name"`
	Readouts []CircHeatmapReadout `json:"readouts"`
}

// CircHeatmapReadout contains data on a readout for a circular heatmap
type CircHeatmapReadout struct {
	Known    bool                      `json:"known"`
	Label    string                    `json:"label"`
	Segments map[string]RoundedSegment `json:"segments"`
}

// CircHeatmapLegend is a slice of legend elements
type CircHeatmapLegend []CircHeatmapLegendElement

// CircHeatmapLegendElement contains settings for each metric in the circheatmap.
type CircHeatmapLegendElement struct {
	Attribute string  `json:"attribute"`
	Color     string  `json:"color"`
	Max       float64 `json:"max"`
	Min       float64 `json:"min"`
}

// Matrices holds input data formatted as matrices.
type Matrices struct {
	Abundance, Ratio, Score [][]float64
	Conditions, Readouts    []string
}

// RoundedSegment rounds circheatmap to a precision of 2
type RoundedSegment float64

// MarshalJSON rounds scatter points to a precision of 2 when exporting to JSON
func (r RoundedSegment) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(r), 'f', 2, 64)), nil
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
	AbundanceType                string
	AutomaticallySetFill         bool
	Condition                    string
	Control                      string
	Files                        []string
	FillColor                    string
	FillMax                      float64
	FillMin                      float64
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
	InvertColor         bool
	RatioDimension      string
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
	KnownFile             string
	OtherAbundance        []string
	ProteinExpressionFile string
	ProteinTissues        []string
	ReadoutIDType         string
	ReadoutMapColumn      string
	ReadoutMapFile        string
	RnaExpressionFile     string
	RnaTissues            []string
	Specificity           bool
	VerticalHeatmap       bool

	// specificity
	SpecificityMetric string
}
