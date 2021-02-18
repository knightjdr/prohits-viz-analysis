package data_test

import (
	"os"
	"regexp"

	"bou.ke/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	. "github.com/knightjdr/prohits-viz-analysis/pkg/tools/analyze/validate/data"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
)

var _ = Describe("Validate parsed data", func() {
	It("should return (not exit) when data is valid", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

		fakeExit := func(int) {
			panic("os.Exit called")
		}
		exitPatch := monkey.Patch(os.Exit, fakeExit)
		defer exitPatch.Unpatch()

		analysis := &types.Analysis{
			Data: []map[string]string{
				map[string]string{"condition": "conditionA", "readout": "readoutA", "readoutLength": "10", "score": "0.1"},
				map[string]string{"condition": "conditionB", "readout": "readoutB", "readoutLength": "15", "score": "0.01"},
				map[string]string{"condition": "conditionA", "readout": "readoutC", "readoutLength": "25", "score": "0"},
				map[string]string{"condition": "conditionB", "readout": "readoutC", "readoutLength": "7", "score": "0.05"},
				map[string]string{"condition": "conditionC", "readout": "readoutA", "readoutLength": "8", "score": "0"},
				map[string]string{"condition": "conditionC", "readout": "readoutB", "readoutLength": "12", "score": "0.01"},
			},
			Settings: types.Settings{
				ReadoutLength: "PreyLength",
				Type:          "dotplot",
			},
		}

		Expect(func() { Validate(analysis) }).To(Not(Panic()))
	})

	It("should exit when there is no data", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

		fakeExit := func(int) {
			panic("os.Exit called")
		}
		exitPatch := monkey.Patch(os.Exit, fakeExit)
		defer exitPatch.Unpatch()

		analysis := &types.Analysis{
			Data: []map[string]string{},
			Settings: types.Settings{
				ReadoutLength: "PreyLength",
				Type:          "dotplot",
			},
		}

		Expect(func() { Validate(analysis) }).To(Panic(), "should exit when there is no data to validate")
		logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
		expected := "no parsed results satisfying filter criteria"
		matched, _ := regexp.MatchString(expected, string(logfile))
		Expect(matched).To(BeTrue(), "should log error message")
	})

	It("should exit when the data is not valid", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "error.txt", []byte(""), 0644)

		fakeExit := func(int) {
			panic("os.Exit called")
		}
		exitPatch := monkey.Patch(os.Exit, fakeExit)
		defer exitPatch.Unpatch()

		analysis := &types.Analysis{
			Data: []map[string]string{
				map[string]string{"condition": "conditionA", "readout": "readoutA", "readoutLength": "10", "score": "NaN"},
				map[string]string{"condition": "conditionA", "readout": "", "readoutLength": "17.3", "score": "0"},
			},
			Settings: types.Settings{
				ReadoutLength: "PreyLength",
				Type:          "dotplot",
			},
		}

		Expect(func() { Validate(analysis) }).To(Panic(), "should exit when there is no data to validate")

		logfile, _ := afero.ReadFile(fs.Instance, "error.txt")
		expectedError := "there are not enough conditions for analysis, min: 2"
		matched, _ := regexp.MatchString(expectedError, string(logfile))
		Expect(matched).To(BeTrue(), "should log error message about condition requirement")
		expectedError = "all readouts should have a name"
		matched, _ = regexp.MatchString(expectedError, string(logfile))
		Expect(matched).To(BeTrue(), "should log error message about readout names")
		expectedError = "score column must contain numeric values, offending value: NaN"
		matched, _ = regexp.MatchString(expectedError, string(logfile))
		Expect(matched).To(BeTrue(), "should log error message about score type requirement")
		expectedError = "readout length column must contain integer values, offending value: 17.3"
		matched, _ = regexp.MatchString(expectedError, string(logfile))
		Expect(matched).To(BeTrue(), "should log error message about readout length type requirement")
	})
})
