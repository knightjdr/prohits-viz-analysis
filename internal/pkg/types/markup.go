package types

// Annotations is the struct for heatmap annotations.
type Annotations struct {
	FontSize int
	List     map[string]Annotation `json:"list"`
}

// Annotation is the struct for heatmap annotation text.
type Annotation struct {
	Position AnnotationPosition
	Text     string
}

// AnnotationPosition is the x,y position for an annotation.
type AnnotationPosition struct {
	X float64
	Y float64
}

// Markers is the struct for heatmap markers.
type Markers struct {
	Color string
	List  map[string]Marker `json:"list"`
}

// Marker is the struct for heatmap marker boxes.
type Marker struct {
	Height int
	Width  int
	X      float64
	Y      float64
}
