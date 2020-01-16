package mapf_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMapf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Map Suite")
}
