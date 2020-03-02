package nocluster_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNocluster(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Nocluster Suite")
}
