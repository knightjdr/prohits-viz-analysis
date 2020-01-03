package heatmap

import "fmt"

func parseRows(data *Data, writeString func(string)) {
	writeRow := defineRowWriter(data, writeString)
	writeRowEnd := defineRowEndWriter(writeString, len(data.Matrices.Abundance))
	writeRowStart := defineRowStartWriter(writeString)

	writeString("\t\"rowDB\": [\n")
	for rowIndex := range data.Matrices.Abundance {
		writeRowStart(data.Matrices.Readouts[rowIndex])
		writeRow(rowIndex)
		writeRowEnd(rowIndex)
	}
	writeString("\t],\n")
}

func defineRowEndWriter(writeString func(string), length int) func(int) {
	penultimateRowIndex := length - 1
	return func(index int) {
		if index < penultimateRowIndex {
			writeString("\t\t\t]\n\t\t},\n")
		} else {
			writeString("\t\t\t]\n\t\t}\n")
		}
	}
}

func defineRowStartWriter(writeString func(string)) func(string) {
	return func(name string) {
		writeString("\t\t{\n")
		writeString(fmt.Sprintf("\t\t\t\"name\": \"%s\",\n", name))
		writeString("\t\t\t\"data\": [\n")
	}
}

func defineRowWriter(data *Data, writeString func(string)) func(int) {
	writeColumnTerminator := defineColumnTerminatorWriter(writeString, len(data.Matrices.Abundance[0]))

	if data.ImageType == "dotplot" {
		return func(rowIndex int) {
			for columnIndex := range data.Matrices.Abundance[rowIndex] {
				abundance := data.Matrices.Abundance[rowIndex][columnIndex]
				ratio := data.Matrices.Ratio[rowIndex][columnIndex]
				score := data.Matrices.Score[rowIndex][columnIndex]
				writeString(fmt.Sprintf("\t\t\t\t{\"ratio\": %0.2f, \"score\": %0.2f, \"value\": %0.2f}", ratio, score, abundance))
				writeColumnTerminator(columnIndex)
			}
		}
	}

	return func(rowIndex int) {
		for columnIndex, abundance := range data.Matrices.Abundance[rowIndex] {
			writeString(fmt.Sprintf("\t\t\t\t{\"value\": %0.2f}", abundance))
			writeColumnTerminator(columnIndex)
		}
	}
}

func defineColumnTerminatorWriter(writeString func(string), length int) func(int) {
	penultimateColumnIndex := length - 1
	return func(index int) {
		if index < penultimateColumnIndex {
			writeString(",\n")
		} else {
			writeString("\n")
		}
	}
}
