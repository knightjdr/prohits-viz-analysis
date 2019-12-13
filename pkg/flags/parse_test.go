package flags_test

import (
	"os"

	"github.com/knightjdr/prohits-viz-analysis/pkg/flags"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Read arguments as specific type", func() {
	It("should return command line arguments as interface", func() {
		os.Args = []string{
			"cmd",
			"-optiona=a",
			"--optionb", "1",
		}

		expected := map[string]interface{}{
			"optiona": "a",
			"optionb": "1",
		}
		Expect(flags.Parse()).To(Equal(expected))
	})
})
