package dotplot

import (
	"encoding/csv"
	"os/exec"

	"github.com/knightjdr/prohits-viz-analysis/fs"
	"github.com/knightjdr/prohits-viz-analysis/logmessage"
)

// BiclustOrder has the order for the biclustering baits and preys.
type BiclustOrder struct {
	Baits, Preys []string
}

// NestedClustering calls the script for the nested clustering algorithm and moves
// files to the appropriate subfolder. It returns the order of the baits and preys.
func NestedClustering() (order BiclustOrder) {
	// Run nested cluster.
	cmd := exec.Command("nestedcluster", "biclustering/matrix.txt", "biclustering/params.txt")
	cmdErr := cmd.Run()

	// Exit if run err.
	logmessage.CheckError(cmdErr, true)

	// Run R script for biclustering output.
	cmd = exec.Command("biclustering.R", "biclustering/matrix.txt")
	cmdErr = cmd.Run()

	// Exit if run err.
	logmessage.CheckError(cmdErr, true)

	// Move nested cluster files.
	fs.Instance.Rename("bait_lists", "biclustering/bait_lists")
	fs.Instance.Rename("bait2bait.pdf", "biclustering/bait2bait.pdf")
	fs.Instance.Rename("baitClusters", "biclustering/baitClusters")
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

	// Grab column and row names from clustered-matrix.txt (this is the clustering)
	// order.

	// Panic if file can't be opened.
	file, err := fs.Instance.Open("biclustering/clustered-matrix.txt")
	logmessage.CheckError(err, true)
	defer file.Close()

	// Read file.
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	lines, err := reader.ReadAll()
	// Panic if file can't be read.
	logmessage.CheckError(err, true)
	// Bait order.
	order.Baits = lines[0][1:]

	// Get prey order.
	for i := 1; i < len(lines); i++ {
		order.Preys = append(order.Preys, lines[i][0])
	}

	return
}
