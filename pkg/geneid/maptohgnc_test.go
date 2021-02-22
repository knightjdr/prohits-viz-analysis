package geneid

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
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

var _ = Describe("Map to HGNC", func() {
	Describe("map IDs to HGNC", func() {
		It("should map ids", func() {
			oldFs := fs.Instance
			defer func() { fs.Instance = oldFs }()
			fs.Instance = afero.NewMemMapFs()

			filename := "test/gene-db.json"
			fs.Instance.MkdirAll("test", 0755)
			afero.WriteFile(fs.Instance, filename, []byte(geneFile), 0444)

			ids := map[string]string{
				"a": "111",
				"b": "222",
				"c": "333",
			}
			idType := "entrez"
			settings := &HGNCsettings{
				File: "test/gene-db.json",
			}

			expectedMapping := map[string]string{
				"a": "1",
				"b": "2",
				"c": "",
			}
			expectedSettings := &HGNCsettings{
				File: "test/gene-db.json",
				DB: HGNCmap{
					"1": {
						AliasSymbol: []string{"A"},
						Ensemblg:    "ENSG00000000001",
						Ensemblp:    []string{"ENSP00000000001"},
						Entrez:      "111",
						PrevSymbol:  []string{"aa", "aaa"},
						Refseqg:     []string{"NM_000001"},
						Refseqp:     []string{"NP_000001"},
						Symbol:      "a",
						Uniprotacc:  []string{"P11111"},
						Uniprotid:   []string{"A_HUMAN"},
					},
					"2": {
						AliasSymbol: []string{},
						Ensemblg:    "ENSG00000000002",
						Ensemblp:    []string{"ENSP00000000002"},
						Entrez:      "222",
						PrevSymbol:  []string{},
						Refseqg:     []string{"NM_000002"},
						Refseqp:     []string{"NP_000002"},
						Symbol:      "b",
						Uniprotacc:  []string{"P22222"},
						Uniprotid:   []string{"B_HUMAN"},
					},
					"3": {
						AliasSymbol: []string{},
						Ensemblg:    "ENSG00000000003",
						Ensemblp:    []string{"ENSP00000000003"},
						Entrez:      "",
						PrevSymbol:  []string{},
						Refseqg:     []string{"NM_000003"},
						Refseqp:     []string{},
						Symbol:      "c",
						Uniprotacc:  []string{"P33333"},
						Uniprotid:   []string{"C_HUMAN"},
					},
				},
				IntermediateIDtoHGNC: map[string]string{
					"111": "1",
					"222": "2",
				},
			}

			actualMapping, actualSettings := MapToHGNC(ids, idType, settings)
			Expect(actualMapping).To(Equal(expectedMapping))
			Expect(actualSettings).To(Equal(expectedSettings))
		})
	})

	Describe("read gene id file", func() {
		It("should read a json file", func() {
			oldFs := fs.Instance
			defer func() { fs.Instance = oldFs }()
			fs.Instance = afero.NewMemMapFs()

			filename := "test/gene-db.json"
			fs.Instance.MkdirAll("test", 0755)
			afero.WriteFile(fs.Instance, filename, []byte(geneFile), 0444)

			expected := HGNCmap{
				"1": {
					AliasSymbol: []string{"A"},
					Ensemblg:    "ENSG00000000001",
					Ensemblp:    []string{"ENSP00000000001"},
					Entrez:      "111",
					PrevSymbol:  []string{"aa", "aaa"},
					Refseqg:     []string{"NM_000001"},
					Refseqp:     []string{"NP_000001"},
					Symbol:      "a",
					Uniprotacc:  []string{"P11111"},
					Uniprotid:   []string{"A_HUMAN"},
				},
				"2": {
					AliasSymbol: []string{},
					Ensemblg:    "ENSG00000000002",
					Ensemblp:    []string{"ENSP00000000002"},
					Entrez:      "222",
					PrevSymbol:  []string{},
					Refseqg:     []string{"NM_000002"},
					Refseqp:     []string{"NP_000002"},
					Symbol:      "b",
					Uniprotacc:  []string{"P22222"},
					Uniprotid:   []string{"B_HUMAN"},
				},
				"3": {
					AliasSymbol: []string{},
					Ensemblg:    "ENSG00000000003",
					Ensemblp:    []string{"ENSP00000000003"},
					Entrez:      "",
					PrevSymbol:  []string{},
					Refseqg:     []string{"NM_000003"},
					Refseqp:     []string{},
					Symbol:      "c",
					Uniprotacc:  []string{"P33333"},
					Uniprotid:   []string{"C_HUMAN"},
				},
			}
			Expect(readGeneIDfile(filename)).To(Equal(expected))
		})
	})

	Describe("define ID parser for extracting IDs from file", func() {
		It("should parse ensemlbg", func() {
			idType := "ensemblg"
			parser := defineParserForIDFromDB(idType)
			ids := GeneIdentifiers{
				Ensemblg: "ENSG00000000001",
			}

			expected := []string{"ENSG00000000001"}
			Expect(parser(ids)).To(Equal((expected)))
		})

		It("should parse ensemlbp", func() {
			idType := "ensemblp"
			parser := defineParserForIDFromDB(idType)
			ids := GeneIdentifiers{
				Ensemblp: []string{"ENSP00000000001", "ENSP00000000011"},
			}

			expected := []string{"ENSP00000000001", "ENSP00000000011"}
			Expect(parser(ids)).To(Equal((expected)))
		})

		It("should parse entrez", func() {
			idType := "entrez"
			parser := defineParserForIDFromDB(idType)
			ids := GeneIdentifiers{
				Entrez: "111",
			}

			expected := []string{"111"}
			Expect(parser(ids)).To(Equal((expected)))
		})

		It("should parse refseqg", func() {
			idType := "refseqg"
			parser := defineParserForIDFromDB(idType)
			ids := GeneIdentifiers{
				Refseqg: []string{"NM_000011", "NM_000001.1", "NM_000001.2"},
			}

			expected := []string{"NM_000001", "NM_000011"}
			Expect(parser(ids)).To(Equal((expected)))
		})

		It("should parse refseqp", func() {
			idType := "refseqp"
			parser := defineParserForIDFromDB(idType)
			ids := GeneIdentifiers{
				Refseqp: []string{"NP_000011", "NP_000001.1", "NP_000001.2"},
			}

			expected := []string{"NP_000001", "NP_000011"}
			Expect(parser(ids)).To(Equal((expected)))
		})

		It("should parse symbol", func() {
			idType := "symbol"
			parser := defineParserForIDFromDB(idType)
			ids := GeneIdentifiers{
				AliasSymbol: []string{"A", "AA"},
				PrevSymbol:  []string{"aa"},
				Symbol:      "a",
			}

			expected := []string{"a", "*A", "*AA", "**aa"}
			Expect(parser(ids)).To(Equal((expected)))
		})

		It("should parse uniprotacc", func() {
			idType := "uniprotacc"
			parser := defineParserForIDFromDB(idType)
			ids := GeneIdentifiers{
				Uniprotacc: []string{"P11111"},
			}

			expected := []string{"P11111"}
			Expect(parser(ids)).To(Equal((expected)))
		})

		It("should parse uniprotid", func() {
			idType := "uniprotid"
			parser := defineParserForIDFromDB(idType)
			ids := GeneIdentifiers{
				Uniprotid: []string{"A_HUMAN"},
			}

			expected := []string{"A_HUMAN"}
			Expect(parser(ids)).To(Equal((expected)))
		})
	})

	Describe("create id to HGNC map", func() {
		It("should create a map for string ID type", func() {
			idType := "entrez"
			db := HGNCmap{
				"1": {Entrez: "111"},
				"2": {Entrez: "222"},
				"3": {Entrez: ""},
			}

			expected := map[string]string{
				"111": "1",
				"222": "2",
			}
			Expect(createIDtoHGNCmap(db, idType)).To(Equal(expected))
		})

		It("should create a map for slice ID type", func() {
			idType := "uniprotacc"
			db := HGNCmap{
				"1": {Uniprotacc: []string{"P11111", "P11112"}},
				"2": {Uniprotacc: []string{"P22222"}},
			}

			expected := map[string]string{
				"P11111": "1",
				"P11112": "1",
				"P22222": "2",
			}
			Expect(createIDtoHGNCmap(db, idType)).To(Equal(expected))
		})

		It("should create a map for symbol and prioritize conflicts by symbol then alias then prevsymbol", func() {
			idType := "symbol"
			db := HGNCmap{
				"1": {
					AliasSymbol: []string{"A", "AA"},
					PrevSymbol:  []string{"aa"},
					Symbol:      "a",
				},
				"2": {
					AliasSymbol: []string{"B", "a"},
					PrevSymbol:  []string{},
					Symbol:      "b",
				},
				"3": {
					AliasSymbol: []string{},
					PrevSymbol:  []string{"a"},
					Symbol:      "c",
				},
				"4": {
					AliasSymbol: []string{},
					PrevSymbol:  []string{"B"},
					Symbol:      "d",
				},
			}

			expected := map[string]string{
				"a":  "1",
				"A":  "1",
				"AA": "1",
				"aa": "1",
				"b":  "2",
				"B":  "2",
				"c":  "3",
				"d":  "4",
			}
			Expect(createIDtoHGNCmap(db, idType)).To(Equal(expected))
		})
	})

	Describe("define formatter for intermediate IDs", func() {
		It("should format entrez as is (i.e. unmodified)", func() {
			idType := "entrez"
			formatter := defineFormatterForIntermediateID(idType)
			id := "111"

			expected := "111"
			Expect(formatter(id)).To(Equal((expected)))
		})

		It("should format refseqg by stripping version numbers", func() {
			idType := "refseqg"
			formatter := defineFormatterForIntermediateID(idType)
			id := "NM_00001.1"

			expected := "NM_00001"
			Expect(formatter(id)).To(Equal((expected)))
		})

		It("should format refseqp by stripping version numbers", func() {
			idType := "refseqp"
			formatter := defineFormatterForIntermediateID(idType)
			id := "NP_00001.1"

			expected := "NP_00001"
			Expect(formatter(id)).To(Equal((expected)))
		})
	})
})
