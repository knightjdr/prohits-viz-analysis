package biclustering_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBiclustering(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Biclustering Suite")
}
