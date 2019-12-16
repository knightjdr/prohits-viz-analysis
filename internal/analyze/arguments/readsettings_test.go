package arguments

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var jsonText = `{
	"abundance": "avgspec",
	"abundanceCap": 50,
	"biclusteringApprox":  true,
	"clustering": "hierarchical",
	"condition": "bait",
	"control": "ctrl",
	"files": ["file1.txt", "file2.txt"],
	"readout":       "prey",
	"readoutLength": "preyLength",
	"score":         "fdr",
	"type": "dotplot"
}`

var _ = Describe("Read settings", func() {
	It("should read JSON file settings", func() {
		// Mock filesystem.
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test/settings.json", []byte(jsonText), 0644)

		expected := &types.Analysis{
			Settings: types.Settings{
				Abundance:          "avgspec",
				AbundanceCap:       50,
				BiclusteringApprox: true,
				Clustering:         "hierarchical",
				Condition:          "bait",
				Control:            "ctrl",
				Files:              []string{"file1.txt", "file2.txt"},
				Readout:            "prey",
				ReadoutLength:      "preyLength",
				Score:              "fdr",
				Type:               "dotplot",
			},
		}
		Expect(readSettings("test/settings.json")).To(Equal(expected))
	})
})
