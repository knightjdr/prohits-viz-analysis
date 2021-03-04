package scv

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/types"
	"github.com/spf13/afero"
)

var geneFile = `{
	"1": {
		"aliasSymbol": ["A"],
		"ensemblg": "ENSG00000000001",
		"ensemblp": ["ENSP00000000001"],
		"entrez": "111",
		"prevSymbol": ["aa", "aaa"],
		"refseqg": ["NM_000001"],
		"refseqp": ["NP_000001"],
		"symbol": "a",
		"uniprotacc": ["P11111"],
		"uniprotid": ["A_HUMAN"]
	},
	"2": {
		"aliasSymbol": [],
		"ensemblg": "ENSG00000000002",
		"ensemblp": ["ENSP00000000002"],
		"entrez": "222",
		"prevSymbol": [],
		"refseqg": ["NM_000002"],
		"refseqp": ["NP_000002"],
		"symbol": "b",
		"uniprotacc": ["P22222"],
		"uniprotid": ["B_HUMAN"]
	},
	"3": {
		"aliasSymbol": [],
		"ensemblg": "ENSG00000000003",
		"ensemblp": ["ENSP00000000003"],
		"entrez": "",
		"prevSymbol": [],
		"refseqg": ["NM_000003"],
		"refseqp": [],
		"symbol": "c",
		"uniprotacc": ["P33333"],
		"uniprotid": ["C_HUMAN"]
	}
}`

var _ = Describe("Map IDs", func() {
	It("should map IDs by themselves", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		filename := "test/gene-db.json"
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, filename, []byte(geneFile), 0444)

		analysis := &types.Analysis{
			Data: []map[string]string{
				{"condition": "a", "readout": "b"},
				{"condition": "a", "readout": "c"},
				{"condition": "b", "readout": "a"},
				{"condition": "b", "readout": "c"},
			},
			Settings: types.Settings{
				ConditionIDType:    "symbol",
				ConditionMapColumn: "",
				ConditionMapFile:   "",
				GeneFile:           "test/gene-db.json",
				ReadoutIDType:      "symbol",
				ReadoutMapColumn:   "",
				ReadoutMapFile:     "",
			},
		}

		expected := map[string]map[string]string{
			"condition": {
				"a": "1",
				"b": "2",
			},
			"readout": {
				"a": "1",
				"b": "2",
				"c": "3",
			},
		}

		Expect(mapIDs(analysis)).To(Equal(expected))
	})

	It("should map IDs by column", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		filename := "test/gene-db.json"
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, filename, []byte(geneFile), 0444)

		analysis := &types.Analysis{
			Data: []map[string]string{
				{"condition": "a", "readout": "b", "readoutid": "P22222"},
				{"condition": "a", "readout": "c", "readoutid": "P33333"},
				{"condition": "b", "readout": "a", "readoutid": "P11111"},
				{"condition": "b", "readout": "c", "readoutid": "P33333"},
			},
			Settings: types.Settings{
				ConditionIDType:    "symbol",
				ConditionMapColumn: "",
				ConditionMapFile:   "",
				GeneFile:           "test/gene-db.json",
				ReadoutIDType:      "uniprotacc",
				ReadoutMapColumn:   "readoutid",
				ReadoutMapFile:     "",
			},
		}

		expected := map[string]map[string]string{
			"condition": {
				"a": "1",
				"b": "2",
			},
			"readout": {
				"a": "1",
				"b": "2",
				"c": "3",
			},
		}

		Expect(mapIDs(analysis)).To(Equal(expected))
	})

	It("should map IDs by file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		readoutMapFile := "a\tENSG00000000001\n" +
			"b\tENSG00000000002\n" +
			"c\tENSG00000000003\n"

		dbFileName := "test/gene-db.json"
		fs.Instance.MkdirAll("test", 0755)
		fs.Instance.MkdirAll("helper-files", 0755)
		afero.WriteFile(fs.Instance, dbFileName, []byte(geneFile), 0444)
		afero.WriteFile(fs.Instance, "helper-files/readout-map.txt", []byte(readoutMapFile), 0444)

		analysis := &types.Analysis{
			Data: []map[string]string{
				{"condition": "a", "readout": "b"},
				{"condition": "a", "readout": "c"},
				{"condition": "b", "readout": "a"},
				{"condition": "b", "readout": "c"},
			},
			Settings: types.Settings{
				ConditionIDType:    "symbol",
				ConditionMapColumn: "",
				ConditionMapFile:   "",
				GeneFile:           "test/gene-db.json",
				ReadoutIDType:      "ensemblg",
				ReadoutMapColumn:   "",
				ReadoutMapFile:     "helper-files/readout-map.txt",
			},
		}

		expected := map[string]map[string]string{
			"condition": {
				"a": "1",
				"b": "2",
			},
			"readout": {
				"a": "1",
				"b": "2",
				"c": "3",
			},
		}

		Expect(mapIDs(analysis)).To(Equal(expected))
	})
})
