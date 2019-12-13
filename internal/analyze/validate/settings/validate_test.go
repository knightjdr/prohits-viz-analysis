package settings

import (
	"os"

	"github.com/bouk/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/internal/pkg/types"
)

var _ = Describe("Setting validation", func() {
	It("should validate circheatmap", func() {
		settings := &types.CircHeatmap{
			File: types.File{
				Abundance:     "avgSpec",
				Condition:     "bait",
				Control:       "ctrl",
				Files:         []string{"file.txt"},
				Readout:       "prey",
				ReadoutLength: "preyLength",
				Score:         "fdr",
			},
			OtherAbundance: []string{"column1", "column2"},
		}
		analysis := types.Analysis{
			Settings: settings,
			Type:     "circheatmap",
		}

		expectedColumnMap := map[string]string{
			"abundance":     "avgSpec",
			"column1":       "column1",
			"column2":       "column2",
			"condition":     "bait",
			"control":       "ctrl",
			"readout":       "prey",
			"readoutLength": "preyLength",
			"score":         "fdr",
		}
		expectedSettings := settings
		actualColumnMap, acutalSettings := Validate(analysis)
		Expect(actualColumnMap).To(Equal(expectedColumnMap), "should return column map")
		Expect(acutalSettings).To(Equal(expectedSettings), "should return validated settings")
	})

	It("should validate dotplot", func() {
		settings := &types.Dotplot{
			File: types.File{
				Abundance:     "avgSpec",
				Condition:     "bait",
				Control:       "ctrl",
				Files:         []string{"file.txt"},
				Readout:       "prey",
				ReadoutLength: "preyLength",
				Score:         "fdr",
			},
		}
		analysis := types.Analysis{
			Settings: settings,
			Type:     "dotplot",
		}

		expectedColumnMap := map[string]string{
			"abundance":     "avgSpec",
			"condition":     "bait",
			"control":       "ctrl",
			"readout":       "prey",
			"readoutLength": "preyLength",
			"score":         "fdr",
		}
		expectedSettings := settings
		actualColumnMap, acutalSettings := Validate(analysis)
		Expect(actualColumnMap).To(Equal(expectedColumnMap), "should return column map")
		Expect(acutalSettings).To(Equal(expectedSettings), "should return validated settings")
	})

	It("should exit when missing required file settings", func() {
		// Mock filesystem for logging errors.
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

		analysis := types.Analysis{
			Settings: types.Dotplot{},
			Type:     "unknown",
		}

		// Mock exit.
		fakeExit := func(int) {
			panic("os.Exit called")
		}
		exitPatch := monkey.Patch(os.Exit, fakeExit)
		defer exitPatch.Unpatch()

		Expect(func() { Validate(analysis) }).To((Panic()), "should exit when missing required settings")
	})
})
