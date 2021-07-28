package biclustering

import (
	"bytes"
	"encoding/csv"
	"os/exec"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
)

func nestedClustering() map[string][]string {
	runNestedCluster()
	return getClusteringOrder()
}

func runNestedCluster() {
	cmdStr := "docker run --rm -v $(pwd)/biclustering/:/files/ --user $(id -u):$(id -g) nestedcluster -m matrix.txt -p parameters.txt"

	cmd := exec.Command(
		"/bin/sh",
		"-c",
		cmdStr,
	)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()

	log.CheckError(err, true)
}

func getClusteringOrder() map[string][]string {
	order := make(map[string][]string)

	file, err := fs.Instance.Open("biclustering/clustered-matrix.txt")
	log.CheckError(err, true)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	lines, err := reader.ReadAll()
	log.CheckError(err, true)

	order["condition"] = lines[0][1:]
	for i := 1; i < len(lines); i++ {
		order["readout"] = append(order["readout"], lines[i][0])
	}

	return order
}
