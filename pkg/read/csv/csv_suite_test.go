package csv_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCsv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Csv Suite")
}
