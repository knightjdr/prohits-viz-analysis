package scv_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestScv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scv Suite")
}
