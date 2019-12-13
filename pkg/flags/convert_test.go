package flags

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Convert", func() {
	Describe("Float", func() {
		It("should convert from float64", func() {
			arg := float64(0.01)
			expected := float64(0.01)
			Expect(convertFloat(arg)).To(Equal(expected))
		})

		It("should convert from int", func() {
			arg := int(1)
			expected := float64(1)
			Expect(convertFloat(arg)).To(Equal(expected))
		})

		It("should convert from string", func() {
			arg := "1"
			expected := float64(1)
			Expect(convertFloat(arg)).To(Equal(expected))
		})

		It("should return default for unhandled type", func() {
			arg := int64(1)
			expected := float64(0)
			Expect(convertFloat(arg)).To(Equal(expected))
		})
	})

	Describe("Int", func() {
		It("should convert from float64", func() {
			arg := float64(1)
			expected := int(1)
			Expect(convertInt(arg)).To(Equal(expected))
		})

		It("should convert from int", func() {
			arg := int(1)
			expected := int(1)
			Expect(convertInt(arg)).To(Equal(expected))
		})

		It("should convert from string", func() {
			arg := "1"
			expected := int(1)
			Expect(convertInt(arg)).To(Equal(expected))
		})

		It("should return default for unhandled type", func() {
			arg := int64(1)
			expected := int(0)
			Expect(convertInt(arg)).To(Equal(expected))
		})
	})

	Describe("String", func() {
		It("should convert from string", func() {
			arg := "a"
			expected := "a"
			Expect(convertString(arg)).To(Equal(expected))
		})

		It("should return default for unhandled type", func() {
			arg := float64(1)
			expected := ""
			Expect(convertString(arg)).To(Equal(expected))
		})
	})
})
