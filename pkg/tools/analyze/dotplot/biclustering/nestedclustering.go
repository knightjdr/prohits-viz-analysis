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
	moveFiles()
	return getClusteringOrder()
}

func runNestedCluster() {
	cmdStr := "docker run --rm -v $(pwd):/files/ nestedcluster -m biclustering/matrix.txt -p biclustering/parameters.txt"

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

func moveFiles() {
	fs.Instance.Rename("bait_lists", "biclustering/bait_lists")
	fs.Instance.Rename("bait2bait.pdf", "biclustering/bait2bait.pdf")
	fs.Instance.Rename("baitClusters", "biclustering/baitClusters")
	fs.Instance.Rename("condition_lists", "biclustering/condition_lists")
	fs.Instance.Rename("condition2condition.pdf", "biclustering/condition2condition.pdf")
	fs.Instance.Rename("conditionClusters", "biclustering/conditionClusters")
	fs.Instance.Rename("clustered-matrix.txt", "biclustering/clustered-matrix.txt")
	fs.Instance.Rename("clusteredData", "biclustering/clusteredData")
	fs.Instance.Rename("Clusters", "biclustering/Clusters")
	fs.Instance.Rename("estimated.pdf", "biclustering/estimated.pdf")
	fs.Instance.Rename("MCMCparameters", "biclustering/MCMCparameters")
	fs.Instance.Rename("NestedClusters", "biclustering/NestedClusters")
	fs.Instance.Rename("NestedMu", "biclustering/NestedMu")
	fs.Instance.Rename("NestedSigma2", "biclustering/NestedSigma2")
	fs.Instance.Rename("stats.pdf", "biclustering/stats.pdf")
	fs.Instance.RemoveAll("OPTclusters")
}

func getClusteringOrder() map[string][]string {
	order := make(map[string][]string, 0)

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
