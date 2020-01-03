package heatmap_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHeatmap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Heatmap Suite")
}
