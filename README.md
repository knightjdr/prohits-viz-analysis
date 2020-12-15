# ProHits-viz analysis scripts

Golang scripts for analyzing data at ProHits-viz.

## Prerequisites

* [librsvg](https://gitlab.gnome.org/GNOME/librsvg/)
* [nestedcluster](https://sourceforge.net/projects/nestedcluster/), only required when using biclustering option for clustering

## Install

```
go get github.com/knightjdr/prohits-viz-analysis
cd $HOME/go/src/github.com/knightjdr/prohits-viz-analysis
go get -d ./...
go install ./...
```

## Nomenclature

We consider screen data to consist of a series of "conditions" that could be, for example, experimental timepoints, treatments or proteomic baits. For each condition there are "readouts" with associated data for the condition. Readouts will typically by genes or proteins. Readouts have an "abundance" that could be a gene expression value, spectral count or peptide intensity. The abundance is always assumed to be a non-negative number. Finally readouts have a "score" which is an indication of the confidence in the abundance.

All of the tools use this nomenclature. Input data files are expected to be in tabular format with a row for each data point, with each data point having a value for the condition, readout, abundance and score. Files can have additional columns and they will be ignored.

| condition   | readout   | abundance | score |
|-------------|-----------|-----------|-------|
| condition a | readout x | 5         | 0.23  |
| condition a | readout y | 15        | 0.04  |
| condition a | readout z | 47        | 0.01  |
| condition b | readout x | 8         | 0.21  |
| condition b | readout y | 13        | 0.06  |
| condition b | readout z | 35        | 0.02  |
| condition c | readout x | 15        | 0.04  |
| condition c | readout y | 5         | 0.23  |
| condition c | readout z | 93        | 0.00  |

A sample file for testing can be found in `sample-files/analysis-file.txt`

## Command line arguments

Use a json file to specify all analysis settings. The `type` field is used to specify
the type of analysis being performed.

```
pvanalyze --settings="settings.json"
```

Settings file format
```
{
  "fileList": ["file1.txt", "file2.txt"],
  "primaryFilter": 0.01,
  "type": "dotplot"
}
```

## Dot plot analysis

### Required arguments

| Argument        | Description                                                                    |
|-----------------|--------------------------------------------------------------------------------|
| fileList        | list of files in csv or tsv format                                             |
| abundance       | name of column containing abundance values                                     |
| condition       | name of column containing condition names                                      |
| readout         | name of column containing readout names                                        |
| score           | name of column containing readout scores                                       |
| primaryFilter   | score filter for readouts, i.e. a readout must pass this filter to be included |
| secondaryFilter | secondary filter for visually marking readouts below the primary filter, but above another threshold |

### Output options

| Argument      | Description                                                      |
|---------------|------------------------------------------------------------------|
| png           | out pngs in addition to svg                                      |
| writeDotplot  | output a dot plot image                                          |
| writeHeatmap  | output a heat map image                                          |
| writeDistance | output condition-condition and readout-readout distance matrices |

### Image options

| Argument              | Description                                   |
|-----------------------|-----------------------------------------------|
| abundanceCap          | threshold for capping abundances on the image |
| edgeColor<sup>1</sup> | edge color on dot plots                       |
| fillColor<sup>1</sup> | fill color on dot plots and heat maps         |

<sup>1</sup> Options: blue, green, grey, red, yellow

### Data filtering
| Argument         | Description                                                |
|------------------|------------------------------------------------------------|
| minimumAbundance | minimum threshold a readout must satisfy                   |
| scoreType        | specify if smaller (lte) or larger (gte) scores are better |

### Data transformation

Data can be transformed prior to analysis:

1. Control values can be subtracted from readout abundances
2. Abundances can be adjusted to the "length" of a readout. For example if a readout is a protein, the abundance can be adjusted to the protein length, with abundances from smaller proteins increased and adundances from larger proteins reduced. Specifically the median readout length is divided by each readout's length to calculate the multiplier to use for adjustment.
3. Abundances can be normalized by condition, either by 1) total condition abundance or 2) a specific readout. For 1) the total readout abundance is summed for each condition and then each condition's readouts are normalized relative to the median of these sums. For 2) the median abundance of a specific readout is used to normalize all other readouts between conditions.
4. Abundances can be log transformed

| Argument             | Description                                                                            |
|----------------------|----------------------------------------------------------------------------------------|
| control              | name of column with control values to subtract from abundances (must be a pipe-separated list) |
| logBase              | log transform data; options: 2, e, 10                                                  |
| normalization        | normalize data; options: none, readout, total                                          |
| normalizationReadout | readout to use for normalization                                                       |
| readoutLength        | name of column with readout lengths for abundance adjustment                           |

### Clustering options

| Argument   | Description                               |
|------------|-------------------------------------------|
| clustering | options: biclustering, hierarchical, none |

#### Hierarchical

| Argument           | Description                                                                           |
|--------------------|---------------------------------------------------------------------------------------|
| clusteringMethod   | linkage method; options: average, centroid, complete, mcquitty, median, single and ward                    |
| clusteringOptimize | optimize leaf order using the method of [Bar-Jospeh, et al.](https://www.ncbi.nlm.nih.gov/pubmed/11472989) |
| distance           | distance metric; options: binary, canberra, euclidean, jaccard, manhattan and maximum |

#### Biclustering

| Argument           | Description                              |
|--------------------|------------------------------------------|
| biclusteringApprox | perform approximate biclustering(faster) |

#### No clustering

| Argument            | Description                                       |
|---------------------|---------------------------------------------------|
| conditionClustering | cluster by condition; options: none or conditions |
| conditionList       | ordered and comma separated list of conditions    |
| readoutClustering   | cluster by readouts; options: none or readouts    |
| readoutList         | ordered and comma separated list of readouts      |

To create images with conditions in a specific order, set conditionClustering to "none" and specify a list of conditions
in the order you wish them to appear. If conditionClustering is set to "conditions", conditions will be hierarchically clustered.

You can control what conditions and readouts are shown on the image by setting both conditionClustering and readoutClustering
to "none" and supplying lists for each. Alternatively, if you only what to specify a list of conditions, set readoutClustering
to "readouts" and all readouts will be included in the analysis and they will be hierarchically clustered.

## Tests

`go test`
