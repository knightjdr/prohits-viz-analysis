package flags_test

import (
	"github.com/knightjdr/prohits-viz-analysis/pkg/flags"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Read arguments as specific type", func() {
	Describe("Float", func() {
		It("should set flag from argument", func() {
			args := map[string]interface{}{
				"arg": "1.5",
			}
			Expect(flags.SetFloat("arg", args, 0.5)).To(Equal(1.5))
		})

		It("should set flag from default", func() {
			args := map[string]interface{}{}
			Expect(flags.SetFloat("arg", args, 0.5)).To(Equal(0.5))
		})
	})

	Describe("Int", func() {
		It("should set flag from argument", func() {
			args := map[string]interface{}{
				"arg": "2",
			}
			Expect(flags.SetInt("arg", args, 1)).To(Equal(2))
		})

		It("should set flag from default", func() {
			args := map[string]interface{}{}
			Expect(flags.SetInt("arg", args, 1)).To(Equal(1))
		})
	})

	Describe("String", func() {
		It("should set flag from argument", func() {
			args := map[string]interface{}{
				"arg": "a",
			}
			Expect(flags.SetString("arg", args, "b")).To(Equal("a"))
		})

		It("should set flag from default", func() {
			args := map[string]interface{}{}
			Expect(flags.SetString("arg", args, "b")).To(Equal("b"))
		})
	})
})
