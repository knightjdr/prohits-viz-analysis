package helper

import "bytes"

// StringConcat concatenates an array of strings.
func StringConcat(arr []string) string {
	var buffer bytes.Buffer
	for _, value := range arr {
		buffer.WriteString(value)
	}
	return buffer.String()
}
