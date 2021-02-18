package geneid_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGeneid(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Geneid Suite")
}
