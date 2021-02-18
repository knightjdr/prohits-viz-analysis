// Package geneid contains functions for mapping gene identifiers.
package geneid

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/knightjdr/prohits-viz-analysis/pkg/fs"
	"github.com/knightjdr/prohits-viz-analysis/pkg/log"
	"github.com/knightjdr/prohits-viz-analysis/pkg/slice"
	"github.com/spf13/afero"
)

// GeneIdentifiers are identifiers an HGNC ID can be mapped to.
type GeneIdentifiers struct {
	AliasSymbol []string
	Ensemblg    string
	Ensemblp    []string
	Entrez      string
	PrevSymbol  []string
	Refseqg     []string
	Refseqp     []string
	Symbol      string
	Uniprotacc  []string
	Uniprotid   []string
}

// HGNCmap are identifiers an HGNC ID can be mapped to.
type HGNCmap map[string]GeneIdentifiers

// HGNCsettings are configuration options for reading and mapping HGNC identifiers.
type HGNCsettings struct {
	File     string
	DB       HGNCmap
	IDtoHGNC map[string]string
}

// MapToHGNC maps identifiers to HGNC
func MapToHGNC(ids map[string]string, idType string, settings *HGNCsettings) {
	if settings.IDtoHGNC == nil && settings.DB == nil {
		settings.DB = readGeneIDfile(settings.File)
		settings.IDtoHGNC = createIDtoHGNCmap(settings.DB, idType)
	}
	if settings.IDtoHGNC == nil {
		settings.IDtoHGNC = createIDtoHGNCmap(settings.DB, idType)
	}
}

func readGeneIDfile(filename string) HGNCmap {
	file, err := fs.Instance.Open(filename)
	log.CheckError(err, true)

	bytes, err := afero.ReadAll(file)
	log.CheckError(err, true)

	genemap := &HGNCmap{}
	json.Unmarshal(bytes, genemap)
	return *genemap
}

func createIDtoHGNCmap(db HGNCmap, idType string) map[string]string {
	preliminaryMap := make(map[string]string, 0)

	parser := defineIDParser(idType)
	for hgncID, geneIDs := range db {
		parsed := parser(geneIDs)
		for _, id := range parsed {
			if _, ok := preliminaryMap[id]; !ok {
				preliminaryMap[id] = hgncID
			}
		}
	}

	consolidatedMap := make(map[string]string, 0)
	if idType == "symbol" {
		for id, hgncID := range preliminaryMap {
			symbol := strings.TrimLeft(id, "*")
			alias := fmt.Sprintf("*%s", symbol)
			prev := fmt.Sprintf("**%s", symbol)

			if id == prev {
				_, isAlias := preliminaryMap[alias]
				_, isSymbol := preliminaryMap[symbol]
				_, isDupe := consolidatedMap[symbol]
				if !isAlias && !isSymbol && !isDupe {
					consolidatedMap[symbol] = hgncID
				}
			}

			if id == alias {
				_, isSymbol := preliminaryMap[symbol]
				_, isDupe := consolidatedMap[symbol]
				if !isSymbol && !isDupe {
					consolidatedMap[symbol] = hgncID
				}
			}

			if id == symbol {
				_, isDupe := consolidatedMap[symbol]
				if !isDupe {
					consolidatedMap[symbol] = hgncID
				}
			}
		}
	} else {
		consolidatedMap = preliminaryMap
	}

	return consolidatedMap
}

func defineIDParser(idType string) func(GeneIdentifiers) []string {
	if idType == "ensemblg" {
		return func(geneIDs GeneIdentifiers) []string {
			return []string{geneIDs.Ensemblg}
		}
	}

	if idType == "ensemblp" {
		return func(geneIDs GeneIdentifiers) []string {
			return geneIDs.Ensemblp
		}
	}

	if idType == "entrez" {
		return func(geneIDs GeneIdentifiers) []string {
			return []string{geneIDs.Entrez}
		}
	}

	if idType == "refseqg" {
		return func(geneIDs GeneIdentifiers) []string {
			parsed := make([]string, 0)
			for _, id := range geneIDs.Refseqg {
				parsed = append(parsed, strings.Split(id, ".")[0])
			}
			return slice.UniqueStrings(parsed)
		}
	}

	if idType == "refseqp" {
		return func(geneIDs GeneIdentifiers) []string {
			parsed := make([]string, 0)
			for _, id := range geneIDs.Refseqp {
				parsed = append(parsed, strings.Split(id, ".")[0])
			}
			return slice.UniqueStrings(parsed)
		}
	}

	if idType == "symbol" {
		return func(geneIDs GeneIdentifiers) []string {
			parsed := make([]string, 0)
			parsed = append(parsed, geneIDs.Symbol)
			for _, id := range geneIDs.AliasSymbol {
				parsed = append(parsed, fmt.Sprintf("*%s", id))
			}
			for _, id := range geneIDs.PrevSymbol {
				parsed = append(parsed, fmt.Sprintf("**%s", id))
			}
			return parsed
		}
	}

	if idType == "uniprotacc" {
		return func(geneIDs GeneIdentifiers) []string {
			return geneIDs.Uniprotacc
		}
	}

	if idType == "uniprotid" {
		return func(geneIDs GeneIdentifiers) []string {
			return geneIDs.Uniprotid
		}
	}

	return func(geneIDs GeneIdentifiers) []string {
		return []string{}
	}
}
