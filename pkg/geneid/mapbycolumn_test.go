package geneid

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Map by column", func() {
	It("should map IDs using another column/key from data and not trim target name", func() {
		data := []map[string]string{
			{"condition": "A_X", "conditionid": "1"},
			{"condition": "B", "conditionid": "2_2"},
			{"condition": "C", "conditionid": "3"},
			{"condition": "D_X", "conditionid": ""},
		}
		mapColumn := "conditionid"
		sourceColumn := "condition"
		trimTarget := false

		expected := map[string]string{
			"A_X": "1",
			"B":   "2_2",
			"C":   "3",
			"D_X": "D",
		}
		Expect(MapByColumn(data, sourceColumn, mapColumn, trimTarget)).To(Equal(expected))
	})

	It("should map IDs using another column/key from data and trim target name", func() {
		data := []map[string]string{
			{"condition": "A_X", "conditionid": "1"},
			{"condition": "B", "conditionid": "2_2"},
			{"condition": "C", "conditionid": "3"},
			{"condition": "D_X", "conditionid": ""},
		}
		mapColumn := "conditionid"
		sourceColumn := "condition"
		trimTarget := true

		expected := map[string]string{
			"A_X": "1",
			"B":   "2",
			"C":   "3",
			"D_X": "D",
		}
		Expect(MapByColumn(data, sourceColumn, mapColumn, trimTarget)).To(Equal(expected))
	})
})
