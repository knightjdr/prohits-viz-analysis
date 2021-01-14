package scatter_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestScatter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scatter Suite")
}
