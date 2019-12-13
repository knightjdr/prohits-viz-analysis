package arguments

import (
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var settingsText = `{
	"file": {
		"abundance": "avgspec",
		"condition": "bait",
		"control": "ctrl",
		"files": ["file1.txt", "file2.txt"],
		"readout":       "prey",
		"readoutLength": "preyLength",
		"score":         "fdr"
	},
	"abundanceCap": 50,
	"biclusteringApprox":  true,
	"clustering": "hierarchical"
}`

var _ = Describe("Read settings", func() {
	It("should read JSON file settings", func() {
		// Mock filesystem.
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test/settings.json", []byte(settingsText), 0644)

		expected := &types.Dotplot{
			File: types.File{
				Abundance:     "avgspec",
				Condition:     "bait",
				Control:       "ctrl",
				Files:         []string{"file1.txt", "file2.txt"},
				Readout:       "prey",
				ReadoutLength: "preyLength",
				Score:         "fdr",
			},
			AbundanceCap:       50,
			BiclusteringApprox: true,
			Clustering:         "hierarchical",
		}
		Expect(readSettings("dotplot", "test/settings.json")).To(Equal(expected))
	})
})
