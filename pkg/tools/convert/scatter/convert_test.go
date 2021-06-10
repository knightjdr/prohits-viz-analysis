package scatter

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/spf13/afero"
)

var _ = Describe("Convert scatter plot", func() {
	It("should convert from format 1", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fileContents := "entry\tPIK3CA_E545K vs PIK3CA\tLog[e] AvgSpec\tLog[2](Fold change PIK3CA_E545K/PIK3CA)\tBinary\tBFDR\n" +
			"ALMS1\t34.57\t9.49\t6\tblack\n" +
			"AMOTL1\t59\t4.18\t6\tblack\n" +
			"AMOTL2\t7.5\t4.48\t6\tred\n"

		fs.Instance.MkdirAll("test", 0755)
		fs.Instance.MkdirAll("interactive", 0755)
		afero.WriteFile(fs.Instance, "test/file.txt", []byte(fileContents), 0444)

		Convert("test/file.txt")

		expected := "{\n" +
			"\t\"legend\": [{\"color\":\"#ff0000\",\"text\":\"Infinite fold change\"}],\n" +
			"\t\"parameters\": {\"abundanceColumn\":\"\",\"analysisType\":\"condition-condition\",\"conditionColumn\":\"\",\"controlColumn\":\"\",\"files\":[],\"imageType\":\"scatter\",\"mockConditionAbundance\":false,\"normalization\":\"\",\"readoutColumn\":\"\",\"scoreColumn\":\"BFDR\",\"scoreType\":\"\"},\n" +
			"\t\"settings\": {\"xFilter\":0,\"yFilter\":0},\n" +
			"\t\"plots\": [" +
			"{\"labels\":{\"x\":\"Log[e] AvgSpec\",\"y\":\"Log[2](Fold change PIK3CA_E545K/PIK3CA)\"},\"name\":\"PIK3CA_E545K vs PIK3CA\"," +
			"\"points\":[" +
			"{\"color\":\"black\",\"label\":\"ALMS1\",\"x\":34.57,\"y\":9.49}," +
			"{\"color\":\"black\",\"label\":\"AMOTL1\",\"x\":59.00,\"y\":4.18}," +
			"{\"color\":\"red\",\"label\":\"AMOTL2\",\"x\":7.50,\"y\":4.48}]}]" +
			"\n}\n"

		actual, _ := afero.ReadFile(fs.Instance, "interactive/file.json")
		Expect(string(actual)).To(Equal(expected))
	})

	It("should convert from format 2", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fileContents := "details:\t{\"tool\": \"Binary-biDist\", \"type\": \"Bait vs bait\", \"bait\": \"p120CtnWT v p120CtnS320A\", \"xAxis\": \"Total Spectral Count of p120CtnWT\", \"yAxis\": \"Total Spectral Count of p120CtnS320A\", \"score\": \"FDR\", \"filter\": 0, \"primary\": 0.01, \"secondary\": 0.05}\n" +
			"ALMS1\t34.57\t9.49\t6\t#509afb\n" +
			"AMOTL1\t59\t4.18\t6\t#509afb\n" +
			"AMOTL2\t7.5\t4.48\t6\t#0066cc\n"

		fs.Instance.MkdirAll("test", 0755)
		fs.Instance.MkdirAll("interactive", 0755)
		afero.WriteFile(fs.Instance, "test/file.txt", []byte(fileContents), 0444)

		Convert("test/file.txt")

		expected := "{\n" +
			"\t\"legend\": [{\"color\":\"#0066cc\",\"text\":\"FDR ≤ 0.01\"},{\"color\":\"#99ccff\",\"text\":\"0.01 \\u003c FDR ≤ 0.05\"}],\n" +
			"\t\"parameters\": {\"abundanceColumn\":\"\",\"analysisType\":\"condition-condition\",\"conditionColumn\":\"\",\"controlColumn\":\"\",\"files\":[],\"imageType\":\"scatter\",\"mockConditionAbundance\":false,\"normalization\":\"\",\"readoutColumn\":\"\",\"scoreColumn\":\"FDR\",\"scoreType\":\"lte\"},\n" +
			"\t\"settings\": {\"xFilter\":0,\"yFilter\":0},\n" +
			"\t\"plots\": [" +
			"{\"labels\":{\"x\":\"Total Spectral Count of p120CtnWT\",\"y\":\"Total Spectral Count of p120CtnS320A\"},\"name\":\"p120CtnWT v p120CtnS320A\"," +
			"\"points\":[" +
			"{\"color\":\"#509afb\",\"label\":\"ALMS1\",\"x\":34.57,\"y\":9.49}," +
			"{\"color\":\"#509afb\",\"label\":\"AMOTL1\",\"x\":59.00,\"y\":4.18}," +
			"{\"color\":\"#0066cc\",\"label\":\"AMOTL2\",\"x\":7.50,\"y\":4.48}]}]\n}\n"

		actual, _ := afero.ReadFile(fs.Instance, "interactive/file.json")
		Expect(string(actual)).To(Equal(expected))
	})
})

var _ = Describe("Determine input file format", func() {
	It("should identify format 1", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fileContents := "entry\t1title\txLabel\tyLabel\n" +
			"pointA\t1\t2\t5\n" +
			"pointB\t3\t4\t5\n"
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test/file.txt", []byte(fileContents), 0444)

		expected := 1
		Expect(determineFormat("test/file.txt")).To(Equal(expected))
	})

	It("should identify format 2", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fileContents := "details:\t{\"tool\": \"Binary-biDist\", \"type\": \"Bait vs bait\"}\n" +
			"pointA\t1\t2\t5\n" +
			"pointB\t3\t4\t5\n"
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test/file.txt", []byte(fileContents), 0444)

		expected := 2
		Expect(determineFormat("test/file.txt")).To(Equal(expected))
	})
})
