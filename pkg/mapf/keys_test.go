package mapf_test

import (
	"sort"

	. "github.com/knightjdr/prohits-viz-analysis/pkg/mapf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keys", func() {
	Describe("Map of type map[string]interface", func() {
		It("should return keys from map[string]interface", func() {
			m := map[string]interface{}{
				"a": 0.01,
				"d": 0.05,
				"c": 2,
			}

			actual := KeysStringInterface(m)
			sort.Strings(actual)
			expected := []string{"a", "c", "d"}
			Expect(actual).To(Equal(expected), "should return integer keys")
		})
	})

	Describe("Map of type map[int]bool", func() {
		It("should return integer keys", func() {
			m := map[int]bool{
				3: true,
				1: true,
				5: true,
			}

			actual := KeysIntBool(m)
			sort.Ints(actual)
			expected := []int{1, 3, 5}
			Expect(actual).To(Equal(expected), "should return integer keys")
		})
	})

	Describe("Map of type map[int]float64", func() {
		It("should return integer keys", func() {
			m := map[int]float64{
				3: 0.1,
				1: 0.1,
				5: 0.2,
			}

			actual := KeysIntFloat(m)
			sort.Ints(actual)
			expected := []int{1, 3, 5}
			Expect(actual).To(Equal(expected), "should return integer keys")
		})
	})

	Describe("Map of type map[string]bool", func() {
		It("should return string keys", func() {
			m := map[string]bool{
				"a": true,
				"d": true,
				"c": true,
			}

			actual := KeysStringBool(m)
			sort.Strings(actual)
			expected := []string{"a", "c", "d"}
			Expect(actual).To(Equal(expected), "should return string keys")
		})
	})

	Describe("Map of type map[string]float64", func() {
		It("should return string keys", func() {
			m := map[string]float64{
				"a": 0.01,
				"d": 0.05,
				"c": 2,
			}

			actual := KeysStringFloat(m)
			sort.Strings(actual)
			expected := []string{"a", "c", "d"}
			Expect(actual).To(Equal(expected), "should return string keys")
		})
	})

	Describe("Map of type map[string]string", func() {
		It("should return string keys", func() {
			m := map[string]string{
				"a": "test",
				"d": "test",
				"c": "test",
			}

			actual := KeysStringString(m)
			sort.Strings(actual)
			expected := []string{"a", "c", "d"}
			Expect(actual).To(Equal(expected), "should return string keys")
		})
	})
})
