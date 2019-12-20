package dimensions_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDimensions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dimensions Suite")
}
